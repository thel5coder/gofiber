package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gofiber/pkg/str"
	"gofiber/server/requests"
	"gofiber/usecase"
	"net/http"
)

type UserHandler struct {
	Handler
}

//browse
func (handler UserHandler) Browse(ctx *fiber.Ctx) error {
	search := ctx.Query("search")
	orderBy := ctx.Query("order_by")
	sort := ctx.Query("sort")
	page := str.StringToInt(ctx.Query("page"))
	limit := str.StringToInt(ctx.Query("limit"))

	uc := usecase.UserUseCase{UcContract: handler.UcContract}
	res, meta, err := uc.Browse(search, orderBy, sort, page, limit)

	return handler.SendResponse(ctx, res, meta, err, 0)
}

//read
func (handler UserHandler) Read(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")

	uc := usecase.UserUseCase{UcContract: handler.UcContract}
	res, err := uc.Read(ID)

	return handler.SendResponse(ctx, res, nil, err, 0)
}

//edit
func (handler UserHandler) Edit(ctx *fiber.Ctx) error {
	ID := ctx.Params("id")
	input := new(requests.UserRequest)

	if err := ctx.BodyParser(input); err != nil {
		return handler.SendErrorResponse(ctx,err.Error(),http.StatusBadRequest)
	}
	if err := handler.Validator.Struct(input); err != nil {
		errMessage := handler.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return handler.SendErrorResponse(ctx,errMessage,http.StatusBadRequest)
	}

	uc := usecase.UserUseCase{UcContract: handler.UcContract}
	res, err := uc.Edit(input, ID)
	if err != nil {
		return handler.SendErrorResponse(ctx,err.Error(),http.StatusUnprocessableEntity)
	}

	return handler.SendSuccessResponse(ctx,res,nil)
}

//add
func(handler UserHandler) Add(ctx *fiber.Ctx) error{
	input := new(requests.UserRequest)

	if err := ctx.BodyParser(input); err != nil {
		return handler.SendErrorResponse(ctx,err.Error(),http.StatusBadRequest)
	}
	if err := handler.Validator.Struct(input); err != nil {
		errMessage := handler.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return handler.SendErrorResponse(ctx,errMessage,http.StatusBadRequest)
	}

	uc := usecase.UserUseCase{UcContract: handler.UcContract}
	res, err := uc.Add(input)
	if err != nil {
		return handler.SendErrorResponse(ctx,err.Error(),http.StatusUnprocessableEntity)
	}

	return handler.SendSuccessResponse(ctx,res,nil)
}

//delete
func(handler UserHandler) Delete(ctx *fiber.Ctx) error{
	ID := ctx.Params("id")

	uc := usecase.UserUseCase{UcContract: handler.UcContract}
	res, err := uc.Delete(ID)

	return handler.SendResponse(ctx, res, nil, err, 0)
}
