package repositories

import (
	"fmt"
	"log"

	"github.com/priscila-albertini-ciandt/dojo-nov/internal/entities"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type IOrderRepository interface {
	CreateOrder(order entities.Order) error
	UpdateOrder(order entities.Order) (entities.Order, error)
	DeleteOrder(order entities.Order) (entities.Order, error)
	DeleteOrderById(id int64) error
	FindAllOrders() ([]entities.Order, error)
	FindOrdersNotCanceled() ([]entities.Order, error)
	FindNewerOrders() ([]entities.Order, error)
	FindFirtsTenOrders() ([]entities.Order, error)
	FindOrdersWithTotalGreaterThan(total decimal.Decimal) ([]entities.Order, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return OrderRepository{
		db: db,
	}
}

func (r OrderRepository) CreateOrder(order entities.Order) error {

	// Exemplo transaction
	tx := r.db.Begin()

	result := tx.Create(&order)

	if result.Error != nil {

		log.Panic("Error in create order:", result.Error)

		tx.Rollback()

		return result.Error
	}

	tx.Commit()

	return nil
}

func (r OrderRepository) UpdateOrder(order entities.Order) (entities.Order, error) {

	result := r.db.Save(&order)

	if result.Error != nil {
		log.Panic("Error in update order:", result.Error)
		return entities.Order{}, result.Error
	}

	return order, nil
}

func (r OrderRepository) DeleteOrder(order entities.Order) (entities.Order, error) {

	result := r.db.Delete(&order)

	if result.Error != nil {
		log.Panic("Error in delete order:", result.Error)
		return entities.Order{}, result.Error
	}

	return order, nil
}

func (r OrderRepository) DeleteOrderById(id int64) error {

	result := r.db.Delete(&entities.Order{}, id)

	if result.Error != nil {
		log.Panic("Error in delete order by id:", result.Error)
		return result.Error
	}

	return nil
}

func (r OrderRepository) FindAllOrders() ([]entities.Order, error) {
	var orders []entities.Order

	result := r.db.Find(&orders)

	if result.Error != nil {
		log.Panic("Error in find all orders:", result.Error)
		return orders, result.Error
	}

	// verificar quantos registros foram afetados
	fmt.Println("Retornou ", result.RowsAffected, " ordens")

	return orders, nil
}

func (r OrderRepository) FindOrdersNotCanceled() ([]entities.Order, error) {
	var orders []entities.Order

	result := r.db.Not("status = ?", "canceled").Find(&orders)

	if result.Error != nil {
		log.Panic("Error in find orders not canceled:", result.Error)
		return orders, result.Error
	}

	return orders, nil
}

func (r OrderRepository) FindNewerOrders() ([]entities.Order, error) {
	var orders []entities.Order

	result := r.db.Order("created_at desc").Find(&orders)

	if result.Error != nil {
		log.Panic("Error in find newer orders:", result.Error)
		return orders, result.Error
	}

	return orders, nil
}

func (r OrderRepository) FindFirstTenOrders() ([]entities.Order, error) {
	var orders []entities.Order

	result := r.db.Order("created_at asc").Limit(10).Find(&orders)

	if result.Error != nil {
		log.Panic("Error in find first ten orders:", result.Error)
		return orders, result.Error
	}

	return orders, nil
}

func (r OrderRepository) FindOrdersWithTotalGreaterThan(total decimal.Decimal) ([]entities.Order, error) {
	var orders []entities.Order

	result := r.db.Raw("SELECT id, total FROM orders WHERE total > ?", total).Find(&orders)

	if result.Error != nil {
		log.Panic("Error in find orders with total greater than:", result.Error)
		return orders, result.Error
	}

	return orders, nil
}
