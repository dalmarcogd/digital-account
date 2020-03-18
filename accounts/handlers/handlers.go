package handlers

import (
	v1 "github.com/dalmarcogd/digital-account/accounts/handlers/v1"
	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET("/health-check", HealthCheckHandler)
	// Pedidos
	pedidosGroup := e.Group("/accounts-api")
	pedidosGroupV1 := pedidosGroup.Group("/v1")
	pedidosGroupV1.POST("/accounts", v1.AccountsCreateV1Handler)
	pedidosGroupV1.GET("/accounts/:accountId", v1.AccountsGetV1Handler)
}
