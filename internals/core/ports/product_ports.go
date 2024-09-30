package ports

import (
	model "crud-hex/internals/core/domain"
	utils "crud-hex/internals/utils"
	// fiber "github.com/gofiber/fiber/v2"
)

type IProductService interface{
	FindAll() utils.ServiceResponse
	Create(productData map[string]string) utils.ServiceResponse
	FindByID(id int) utils.ServiceResponse
	Update(id int, productData map[string]string) utils.ServiceResponse
	Delete(id int) utils.ServiceResponse
}

type IProductRepository interface{
	FindAll() ([]model.Product, error)
	Create(product *model.Product) error
	FindByID(id int) (model.Product, error)
	Update(product model.Product) error
	Delete(id int) error
}

// type IProductHandlers interface{
// 	FindAll(c *fiber.Ctx) utils.ServiceResponse
// 	Create(c *fiber.Ctx) utils.ServiceResponse
// }

// type IServer interface{
// 	Initialize()
// }