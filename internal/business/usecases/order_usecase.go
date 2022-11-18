package usecases

import (
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/entities"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/infra/repositories"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/ui/models"
	"gorm.io/gorm"
)

type IOrderUseCase interface {
	CreateOrder(o models.Order)
	ListAllOrders() ([]entities.Order, error)
}

type OrderUseCase struct {
	db              *gorm.DB
	orderRepository repositories.IOrderRepository
}

func NewCreateOrderUseCase(db *gorm.DB, orderRepository repositories.IOrderRepository) IOrderUseCase {
	return OrderUseCase{
		db:              db,
		orderRepository: orderRepository,
	}
}

func (u OrderUseCase) CreateOrder(o models.Order) {

}

func (u OrderUseCase) ListAllOrders() ([]entities.Order, error) {
	return u.orderRepository.FindAllOrders()
}
