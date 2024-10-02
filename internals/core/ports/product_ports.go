package ports

import (
	model "crud-hex/internals/core/domain"
	"crud-hex/pkg/utils"
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

type IProfilingService interface {
	Log(profiling model.Profiling) error
}

type IProfilingRepository interface {
	Create(profiling model.Profiling) error
}