package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gofiber/pkg/functioncaller"
	jwtPkg "gofiber/pkg/jwt"
	"gofiber/pkg/logruslogger"
	"gofiber/pkg/messages"
	"gofiber/server/handlers"
	"gofiber/usecase"
	"net/http"
	"strings"
	"time"
)

type JwtMiddleware struct {
	*usecase.UcContract
}

//jwt middleware
func (jwtMiddleware JwtMiddleware) New(ctx *fiber.Ctx) (err error) {
	claims := &jwtPkg.CustomClaims{}
	handler := handlers.Handler{UcContract: jwtMiddleware.UcContract}

	//check header is present or not
	header := ctx.Get("Authorization")
	if !strings.Contains(header, "Bearer") {
		logruslogger.Log(logruslogger.WarnLevel, messages.HeaderNotPresent, functioncaller.PrintFuncName(), "middleware-jwt-checkHeader")
		return handler.SendResponse(ctx,nil,nil,messages.HeaderNotPresent,http.StatusUnauthorized)
	}

	//check claims and signing method
	token := strings.Replace(header, "Bearer", "", -1)
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			logruslogger.Log(logruslogger.WarnLevel, messages.UnexpectedSigningMethod, functioncaller.PrintFuncName(), "middleware-jwt-checkSigningMethod")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := jwtMiddleware.JwtConfig.SigningKey
		return secret, nil
	})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.UnexpectedClaims, functioncaller.PrintFuncName(), "middleware-jwt-checkClaims")
		return handler.SendResponse(ctx,nil,nil, messages.UnexpectedClaims, http.StatusUnauthorized)
	}

	//check token live time
	if claims.ExpiresAt < time.Now().Unix() {
		logruslogger.Log(logruslogger.WarnLevel, messages.ExpiredToken, functioncaller.PrintFuncName(), "middleware-jwt-checkTokenLiveTime")
		return handler.SendResponse(ctx,nil,nil, messages.ExpiredToken, http.StatusUnauthorized)
	}

	//jwe roll back encrypted id
	jweRes, err := jwtMiddleware.JweCred.Rollback(claims.Id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-rollback")
		return handler.SendResponse(ctx,nil,nil, messages.Unauthorized, http.StatusUnauthorized)
	}
	if jweRes == nil {
		logruslogger.Log(logruslogger.WarnLevel, messages.Unauthorized, functioncaller.PrintFuncName(), "pkg-jwe-resultNil")
		return handler.SendResponse(ctx,nil,nil, messages.Unauthorized, http.StatusUnauthorized)
	}

	//set id to uce case contract
	claims.Id = fmt.Sprintf("%v", jweRes["id"])
	jwtMiddleware.UcContract.UserID = claims.Id

	return ctx.Next()
}
