package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/business/usecases"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/ui/models"
)

// type IOrderController interface {
// 	CreateOrder(ctx *fiber.Ctx) error
// 	ListAllOrders(ctx *fiber.Ctx) error
// }

type OrderController struct {
	usecase usecases.IOrderUseCase
}

func NewOrderController(usecase usecases.IOrderUseCase) OrderController {
	return OrderController{
		usecase: usecase,
	}
}

// CreateOrder func for creates a new order.
// @Description Create a new order.
// @Summary create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Param products body slice of int "Products"
// @Param user_id body string true "User ID"
// @Success 200 {object} models.Order
// @Router /v1/order [post]
func (c OrderController) CreateOrder(ctx *fiber.Ctx) error {

	order := &models.Order{}

	if err := ctx.BodyParser(order); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return nil
}

// ListAllOrders func for list all orders.
// @Description Create a new order.
// @Summary create a new order
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} []models.Order
// @Router /v1/orders [get]
func (c OrderController) ListAllOrders(ctx *fiber.Ctx) error {

	orders, err := c.usecase.ListAllOrders()

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":  true,
			"msg":    "orders were not found",
			"count":  0,
			"orders": nil,
		})
	}

	return ctx.JSON(fiber.Map{
		"error":  false,
		"msg":    nil,
		"count":  len(orders),
		"orders": orders,
	})
}
