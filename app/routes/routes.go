package routes

import (
	"api-loyalty-point-agent/app/middlewares"
	users "api-loyalty-point-agent/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	LoggerMiddleware echo.MiddlewareFunc
	JWTMiddleware    echojwt.Config
	AuthController   users.AuthController
}

func (cl *ControllerList) RegisterRoutes(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)
	user := e.Group("auth")
	user.POST("/register", cl.AuthController.Register)
	user.POST("/login", cl.AuthController.Login)

	users := e.Group("/users", echojwt.WithConfig(cl.JWTMiddleware))
	users.Use(middlewares.VerifyToken)
	users.POST("/logout", cl.AuthController.Logout)
	users.GET("/customers", cl.AuthController.GetAllCustomers)
}
