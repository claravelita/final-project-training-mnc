package handler

import (
	"fmt"

	"github.com/claravelita/final-project-training-mnc/common/command"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func UseCustomValidatorHandler(c *echo.Echo) {
	c.Validator = &CustomValidator{validator: validator.New()}

	c.HTTPErrorHandler = func(err error, c echo.Context) {
		var MessageValidation []string
		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is required",
						err.Field()))
				case "email":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s is not valid email",
						err.Field()))
				case "gte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param()))
				case "lte":
					MessageValidation = append(MessageValidation, fmt.Sprintf("%s value must be lower than %s",
						err.Field(), err.Param()))
				}
			}
			c.JSON(400, command.BadRequestResponses(MessageValidation))
		} else if castedObject, ok := err.(*echo.HTTPError); ok {
			c.JSON(castedObject.Code, castedObject.Message)
		}

	}

}
