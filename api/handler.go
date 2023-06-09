package api

import "github.com/gofiber/fiber/v2"

// IngredientHandler  an interface with operations to be implemented by a specific handler, ie http, gRCP
type IngredientHandler interface {
	Get(ctx *fiber.Ctx)
	Post(ctx *fiber.Ctx)
	Put(ctx *fiber.Ctx)
	Delete(ctx *fiber.Ctx)
	GetAll(ctx *fiber.Ctx)
}
