package routes

import (
	owner "test/fitur/owner"
	handlerowner "test/fitur/owner/handler"
	customer "test/fitur/user"
	handlercostumer "test/fitur/user/handler"
	"test/middlewares"

	"github.com/labstack/echo/v4"
)

func NewHandlerCostumer(Service customer.CostumerService, e *echo.Echo) {
	handlers := &handlercostumer.CustomerHandler{
		CostumerServices: Service,
	}

	e.POST("/costumer/form", handlers.FormData)
	e.POST("/admin/form", handlers.FormAdmin)
	e.POST("/costumer/login", handlers.Login)
	e.GET("/costumer/profile", handlers.Profile, middlewares.JWTMiddleware())
	e.PUT("/costumer", handlers.Update, middlewares.JWTMiddleware())
	e.DELETE("/costumer", handlers.Delete, middlewares.JWTMiddleware())
}
func NewHandlerOwner(Service owner.OwnerService, e *echo.Echo) {
	handlers := &handlerowner.OwnerHandler{
		OwnerServices: Service,
	}

	e.POST("/costumer/owner", handlers.AddOwner, middlewares.JWTMiddleware())

}
