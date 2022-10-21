package dtos

type JSONResponses struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type JSONSwaggerCreatedResponses struct {
	Code   int         `json:"code" example:"201"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type JSONSwaggerOKResponses struct {
	Code   int         `json:"code" example:"200"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

type MessageResponses struct {
	Message string `json:"message"`
}
