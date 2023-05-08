package api

import (
	"github.com/devdowns/ingredients-hexagonal/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"net/http"
)

type handler struct {
	ingredientService domain.IngredientService
}

func (h *handler) Get(ctx *fiber.Ctx) error {
	param := ctx.Params("id")
	id, err := uuid.Parse(param)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(nil)
		return nil
	}

	p, err := h.ingredientService.Find(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Post(ctx *fiber.Ctx) error {
	p := &domain.Ingredient{}
	if err := ctx.BodyParser(&p); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(nil)
	}
	err := h.ingredientService.Store(p)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Put(ctx *fiber.Ctx) error {
	p := &domain.Ingredient{}
	if err := ctx.BodyParser(&p); err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(nil)
	}
	err := h.ingredientService.Update(p)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(nil)
	}
	return ctx.JSON(&p)
}

func (h *handler) Delete(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(nil)
	}

	err = h.ingredientService.Delete(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{"message": "deleted successfully"})
}

func (h *handler) GetAll(ctx *fiber.Ctx) error {
	p, err := h.ingredientService.FindAll()
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(nil)
	}
	return ctx.JSON(&p)
}

// NewHandler  New handler instantiates a http handler for our ingredient service
func NewHandler(ingredientService domain.IngredientService) *handler {
	return &handler{ingredientService: ingredientService}
}
