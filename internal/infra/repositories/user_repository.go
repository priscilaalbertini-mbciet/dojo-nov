package repositories

import (
	"errors"
	"log"

	"github.com/priscila-albertini-ciandt/dojo-nov/internal/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindUserById(userId int64) entities.User
	FindUserByName(name string) entities.User
	FindUsersTotalOrders() []entities.Order
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) FindUserById(userId int64) entities.User {

	var user entities.User

	err := r.db.First(&user, userId).Error

	if err != nil {
		// Verificar tipo do erro
		if errors.Is(err, gorm.ErrRecordNotFound) {

			log.Panic("User not found")
		} else {

			log.Panic("Error finding user by id:", err)
		}
	}

	return user
}

func (r UserRepository) FindUserByName(name string) entities.User {
	var user entities.User

	result := r.db.Where("name = ?", name).First(&user)

	if result.Error != nil {
		log.Panic("Error finding user by name:", result.Error)
	}

	return user
}

func (r UserRepository) FindUsersTotalOrders() []entities.Order {
	var orders []entities.Order

	result := r.db.Model(&entities.User{}).Joins("Order").Select("name, sum(total) as total").Group("name").Find(&orders)

	if result.Error != nil {
		log.Panic("Error in find users total orders:", result.Error)
	}

	return orders
}
