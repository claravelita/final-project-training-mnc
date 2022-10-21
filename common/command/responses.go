package command

import (
	"net/http"

	"github.com/claravelita/final-project-training-mnc/dtos"
)

func SuccessCreatedResponses(data interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Data: data,
		Code: http.StatusCreated,
	}
}

func SuccessResponses(data interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Data: data,
		Code: http.StatusOK,
	}
}

func NotFoundResponses(status string) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Data: nil,
		Code: http.StatusNotFound,
	}
}

func BadRequestResponses(err interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Errors: err,
		Code:   http.StatusBadRequest,
	}
}

func InternalServerResponses(data string, err error) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Data:   nil,
		Code:   http.StatusInternalServerError,
		Errors: err,
	}
}
