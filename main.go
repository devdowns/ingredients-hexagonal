package main

import (
	"github.com/devdowns/ingredients-hexagonal/api"
	"github.com/devdowns/ingredients-hexagonal/config"
	"github.com/devdowns/ingredients-hexagonal/repository"
	"github.com/devdowns/ingredients-hexagonal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {

	conf, _ := config.NewConfig("config/.env")
	conn, _ := repository.GetConnection(&conf.Database)
	repo := repository.NewIngredientRepo(conn)
	ingredientService := service.NewIngredientService(repo)
	ingredientHandler := api.NewHandler(ingredientService)

	r := fiber.New()
	r.Use(logger.New())

	r.Get("/api/ingredient/:id", ingredientHandler.Get)
	r.Post("/api/ingredient", ingredientHandler.Post)
	r.Delete("/api/ingredient/:id", ingredientHandler.Delete)
	r.Get("/api/ingredient", ingredientHandler.GetAll)
	r.Put("/api/ingredient", ingredientHandler.Put)

	if err := r.Listen(":" + conf.Server.Port); err != nil {
		log.Fatal("Failed to start")
	}

}
