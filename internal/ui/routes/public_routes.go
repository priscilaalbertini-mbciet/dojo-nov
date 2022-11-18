package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/business/usecases"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/ui/controllers"
)

func PublicRoutes(a *fiber.App) {

	handler := controllers.NewOrderController(usecases.OrderUseCase{})

	route := a.Group("/app/v1")

	route.Get("/orders", handler.ListAllOrders)
	route.Post("/order", handler.CreateOrder)
	// route.Post("/order", controllers.ListAllOrders)
}
