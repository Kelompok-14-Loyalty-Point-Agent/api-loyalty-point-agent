package routes

import (
	"api-loyalty-point-agent/app/middlewares"
	customers "api-loyalty-point-agent/controllers/customers"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	AuthController   customers.AuthController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)
	customer := e.Group("/customer")
	customer.POST("/register", cl.AuthController.Register)
	customer.POST("/login", cl.AuthController.Login)

	customers := e.Group("/customers", echojwt.WithConfig(cl.JWTMiddleware))
	customers.Use(middlewares.VerifyToken)
	customers.POST("/logout", cl.AuthController.Logout)
	customers.GET("/customersAll", cl.AuthController.GetAllCustomers)
}
