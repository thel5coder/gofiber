package responsedto

import (
	"gofiber/usecase/viewmodel"
)

func ErrorResponse(message interface{}) viewmodel.ResponseErrorVm {
	return viewmodel.ResponseErrorVm{Messages: message}
}

func SuccessResponse(data interface{}, meta interface{}) viewmodel.ResponseSuccessVm {
	return viewmodel.ResponseSuccessVm{
		Data: data,
		Meta: meta,
	}
}
