package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/business/usecases"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/entities"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/infra/repositories"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/ui/controllers"
	"github.com/priscila-albertini-ciandt/dojo-nov/internal/ui/routes"
	"github.com/priscila-albertini-ciandt/dojo-nov/pkg/configs"
	"github.com/priscila-albertini-ciandt/dojo-nov/pkg/database"
	"github.com/priscila-albertini-ciandt/dojo-nov/pkg/middleware"
	"github.com/priscila-albertini-ciandt/dojo-nov/pkg/utils"
	"go.uber.org/fx"
)

func main() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Product{}, &entities.Order{})

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)

	utils.StartServer(app)

	fx.New(
		fx.Provide(
			controllers.NewOrderController,
			usecases.NewCreateOrderUseCase,
			repositories.NewOrderRepository,
			repositories.NewUserRepository,
		),
	).Run()

}
