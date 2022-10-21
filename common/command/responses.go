package command

import (
	"net/http"

	"github.com/claravelita/final-project-training-mnc/dtos"
)

func SuccessCreatedResponses(data interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status: "Success",
		Data:   data,
		Code:   http.StatusCreated,
	}
}

func SuccessResponses(data interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status: "Success",
		Data:   data,
		Code:   http.StatusOK,
	}
}

func NotFoundResponses(status string) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status: status,
		Data:   nil,
		Code:   http.StatusNotFound,
	}
}

func BadRequestResponses(err interface{}) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status: "Bad Request",
		Errors: err,
		Code:   http.StatusBadRequest,
	}
}

func InternalServerResponses(data string, err error) (result dtos.JSONResponses) {
	return dtos.JSONResponses{
		Status: data,
		Data:   nil,
		Code:   http.StatusInternalServerError,
		Errors: err,
	}
}
