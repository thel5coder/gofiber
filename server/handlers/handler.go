package handlers

import (
	"database/sql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gofiber/pkg/jwe"
	"gofiber/pkg/responsedto"
	"gofiber/pkg/str"
	"gofiber/usecase"
	"net/http"
	"strings"
)

type Handler struct {
	FiberApp   *fiber.App
	UcContract *usecase.UcContract
	Jwe        jwe.Credential
	Db         *sql.DB
	Validator  *validator.Validate
	Translator ut.Translator
}

//base send response
func (h Handler) SendResponse(ctx *fiber.Ctx, data interface{}, meta interface{}, err interface{}, code int) error {
	if code == 0 && err != nil {
		code = http.StatusUnprocessableEntity
		err = err.(error).Error()
	}

	if code != http.StatusOK && err != nil {
		return h.SendErrorResponse(ctx, err, code)
	}

	return h.SendSuccessResponse(ctx, data, meta)
}

//send response if status code 200
func (h Handler) SendSuccessResponse(ctx *fiber.Ctx, data interface{}, meta interface{}) error {
	response := responsedto.SuccessResponse(data, meta)

	return ctx.Status(http.StatusOK).JSON(response)
}

//send response if status code != 200
func (h Handler) SendErrorResponse(ctx *fiber.Ctx, err interface{}, code int) error {
	response := responsedto.ErrorResponse(err)

	return ctx.Status(code).JSON(response)
}

//extract error message from validator
func (h Handler) ExtractErrorValidationMessages(error validator.ValidationErrors) map[string][]string {
	errorMessage := map[string][]string{}
	errorTranslation := error.Translate(h.Translator)

	for _, err := range error {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	return errorMessage
}

//handling error
func(h Handler) ErrorHandling(err error) string{
	if err != nil {
		return err.Error()
	}

	return ""
}
