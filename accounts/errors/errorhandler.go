package errors

import (
	"github.com/dalmarcogd/digital-account/accounts/utils"
	"github.com/labstack/echo"
)

func HttpErrorHandler() echo.HTTPErrorHandler {
	return func(err error, context echo.Context) {
		var status int
		if utils.IsInstanceOf(err, &echo.HTTPError{}) {
			status = err.(*echo.HTTPError).Code
		} else if utils.IsInstanceOf(err, &Error{}) {
			status = err.(*Error).StatusCode
		}

		_ = context.JSON(status, err)
	}
}
