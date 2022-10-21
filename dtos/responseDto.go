package dtos

type JSONResponses struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type JSONSwaggerCreatedResponses struct {
	Code   int         `json:"code" example:"201"`
	Status string      `json:"status" example:"Success"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type JSONSwaggerOKResponses struct {
	Code   int         `json:"code" example:"200"`
	Status string      `json:"status" example:"Success"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type MessageResponses struct {
	Message string `json:"message"`
}
