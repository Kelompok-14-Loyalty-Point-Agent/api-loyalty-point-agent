package routes

import (
	"api-loyalty-point-agent/app/middlewares"
	"api-loyalty-point-agent/controllers/customers"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      echojwt.Config
	AuthController     customers.AuthController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {

	e.Use(cl.LoggerMiddleware)

	customers := e.Group("/customer")
	customers.POST("/register", cl.AuthController.Register)
	customers.POST("/login", cl.AuthController.Login)
	// customers.POST("/logout", cl.AuthController.Logout)

	customer := e.Group("/customerData", echojwt.WithConfig(cl.JWTMiddleware))
	customer.Use(middlewares.VerifyToken)
	// customer.GET("/customerAll", cl.AuthController.GetAll)
}
