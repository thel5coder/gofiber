package bootstrap

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gofiber/pkg/jwe"
	"gofiber/pkg/jwt"
	"gofiber/usecase"
)

type Bootstrap struct {
	App        *fiber.App
	DB         *sql.DB
	UcContract usecase.UcContract
	JweCred    jwe.Credential
	Validator  *validator.Validate
	Translator ut.Translator
	JwtCred    jwt.JwtCredential
}
