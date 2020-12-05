package responsedto

import (
	"gofiber/usecase/viewmodel"
)

func ErrorResponse(message interface{}) viewmodel.ResponseErrorVm {
	err := []interface{}{message}
	res :=  viewmodel.ResponseErrorVm{Messages: err}

	return res
}

func SuccessResponse(data interface{}, meta interface{}) viewmodel.ResponseSuccessVm {
	return viewmodel.ResponseSuccessVm{
		Data: data,
		Meta: meta,
	}
}
