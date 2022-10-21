package command

import "github.com/labstack/echo/v4"

func ValidateRequest(context echo.Context, source interface{}) error {
	if err := context.Bind(source); err != nil {
		return err
	}

	if errs := context.Validate(source); errs != nil {
		return errs
	}

	return nil
}
