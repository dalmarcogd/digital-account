package handlers

import (
	v1 "github.com/dalmarcogd/digital-account/transactions/handlers/v1"
	"github.com/labstack/echo"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET("/health-check", HealthCheckHandler)
	// Pedidos
	pedidosGroup := e.Group("/transactions-api")
	pedidosGroupV1 := pedidosGroup.Group("/v1")
	pedidosGroupV1.POST("/transactions", v1.TransactionsCreateV1Handler)
}
