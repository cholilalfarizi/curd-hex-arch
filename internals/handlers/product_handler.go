package handlers

import (
	"crud-hex/internals/core/ports"
	"net/http"

	// "time"

	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct{
	productService ports.IProductService
}

func NewProductController(productService ports.IProductService) *ProductHandler {
	return &ProductHandler{
		productService:   productService,
	}
}

func (c *ProductHandler) FindAll(ctx *fiber.Ctx) error {
	// startTime := time.Now()
	response := c.productService.FindAll()
	// c.logProfiling("FindAll", startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) Create(ctx *fiber.Ctx) error {
	// startTime := time.Now()
	productData := map[string]string{
		"name":  ctx.FormValue("name"),
		"stock": ctx.FormValue("stock"),
		"price": ctx.FormValue("price"),
	}

	response := c.productService.Create(productData)
	// c.logProfiling("Create", startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) FindByID(ctx *fiber.Ctx) error {
	// Retrieve id as string from the URL
	idStr := ctx.Params("id")

	// Convert idStr to an integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// If conversion fails, return a bad request response
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	// Call the service with the integer id
	response := c.productService.FindByID(id)

	// Return the service response
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) Update(ctx *fiber.Ctx) error {
	// startTime := time.Now()
	idStr := ctx.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// If conversion fails, return a bad request response
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	productData := map[string]string{
		"name":  ctx.FormValue("name"),
		"stock": ctx.FormValue("stock"),
		"price": ctx.FormValue("price"),
	}

	response := c.productService.Update(id, productData)
	// c.logProfiling("Update :"+idStr, startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) Delete(ctx *fiber.Ctx) error {
	// startTime := time.Now()
	idStr := ctx.Params("id")
id, err := strconv.Atoi(idStr)
	if err != nil {
		// If conversion fails, return a bad request response
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}
	response := c.productService.Delete(id)
	// c.logProfiling("Delete: "+idStr, startTime)
	return ctx.Status(response.Code).JSON(response)
}
