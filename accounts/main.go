package main

import (
	"fmt"
	"github.com/dalmarcogd/digital-account/accounts/cache"
	"github.com/dalmarcogd/digital-account/accounts/database"
	"github.com/dalmarcogd/digital-account/accounts/environments"
	"github.com/dalmarcogd/digital-account/accounts/errors"
	"github.com/dalmarcogd/digital-account/accounts/handlers"
	"github.com/dalmarcogd/digital-account/accounts/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"

)

func main() {
	database.Migrate()
	defer database.CloseConnection()
	defer cache.CloseConnection()
	env := environments.GetEnvironment()

	e := echo.New()
	e.Validator = utils.NewCustomValidator(validator.New())
	e.HTTPErrorHandler = errors.HttpErrorHandler()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handlers.RegisterHandlers(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", env.Port)))
}
