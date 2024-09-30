package handlers

import (
	model "crud-hex/internals/core/domain"
	"crud-hex/internals/core/ports"
	"net/http"
	"time"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductHandler struct{
	productService ports.IProductService
	profilingService ports.IProfilingService
}

func NewProductController(productService ports.IProductService, profilingService ports.IProfilingService) *ProductHandler {
	return &ProductHandler{
	 productService:   productService,
	 profilingService: profilingService,
	}
   }

func (c *ProductHandler) logProfiling(method string, startTime time.Time) error {
	duration := time.Since(startTime).Milliseconds()
   
	profiling := model.Profiling{
	 ID:        uuid.New(),
	 Method:   method,
	 Duration:  duration,
	 Timestamp: time.Now(),
	}
   
	c.profilingService.Log(profiling)
   
	return nil
   }

func (c *ProductHandler) FindAll(ctx *fiber.Ctx) error {
	startTime := time.Now()
	response := c.productService.FindAll()
	c.logProfiling("FindAll", startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) Create(ctx *fiber.Ctx) error {
	startTime := time.Now()
	productData := map[string]string{
		"name":  ctx.FormValue("name"),
		"stock": ctx.FormValue("stock"),
		"price": ctx.FormValue("price"),
	}

	response := c.productService.Create(productData)
	c.logProfiling("Create", startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) FindByID(ctx *fiber.Ctx) error {
	startTime := time.Now()
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

	c.logProfiling("FindByID: "+idStr, startTime)
	// Return the service response
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) Update(ctx *fiber.Ctx) error {
	startTime := time.Now()
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
	c.logProfiling("Update :"+idStr, startTime)
	return ctx.Status(response.Code).JSON(response)
}

func (c *ProductHandler) Delete(ctx *fiber.Ctx) error {
	startTime := time.Now()
	idStr := ctx.Params("id")
id, err := strconv.Atoi(idStr)
	if err != nil {
		// If conversion fails, return a bad request response
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}
	response := c.productService.Delete(id)
	c.logProfiling("Delete: "+idStr, startTime)
	return ctx.Status(response.Code).JSON(response)
}
