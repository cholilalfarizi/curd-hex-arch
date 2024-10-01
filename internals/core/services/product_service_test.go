package services_test

import (
	"crud-hex/internals/core/domain"
	"crud-hex/internals/core/services"
	"database/sql"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct{
	mock.Mock
}

func (m *MockProductRepository) FindAll() ([]domain.Product, error){
	args := m.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *MockProductRepository) Create(product *domain.Product)error{
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) FindByID(id int) (domain.Product, error){
	args := m.Called(id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) Update(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}


func TestProductService_FindAll_Success(t *testing.T){
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockProducts := []domain.Product{
		{ID: 1 ,Name: "Product1", Stock: 10, Price: 1000, IsAvailable: true},
		{ID: 2 ,Name: "Product2", Stock: 20, Price: 2000, IsAvailable: true},
	}

	mockRepo.On("FindAll").Return(mockProducts, nil)

	response := service.FindAll()

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Get Products Successfully", response.Message)
	assert.Equal(t, mockProducts, response.Data)
	mockRepo.AssertExpectations(t)
}

func TestProductService_FindAll_NotFound(t *testing.T){
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockRepo.On("FindAll").Return([]domain.Product{}, sql.ErrNoRows)

	response := service.FindAll()

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, "Product not found", response.Message)
	assert.Nil(t, response.Data)
	mockRepo.AssertExpectations(t)
	
}

func TestProductService_Create_Success(t *testing.T){
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	productData := map[string]string{
		"name":"New Product",
		"stock":"10",
		"price":"100",
	}

	product := &domain.Product{
		Name:        "New Product",
		Stock:       10,
		Price:       100,
		IsAvailable: true,
	}

	mockRepo.On("Create", product).Return(nil)

	response := service.Create(productData)

	assert.Equal(t, http.StatusCreated, response.Code)
	assert.Equal(t, "Product added successfully", response.Message)
	assert.Equal(t, product, response.Data)
	mockRepo.AssertExpectations(t)
}

func TestProductService_Create_ValidationError(t *testing.T){
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	productData := map[string]string{
		"name":  "",
		"stock": "-10",
		"price": "-100",
	}

	response := service.Create(productData)

	assert.Equal(t, http.StatusBadRequest, response.Code)
	assert.Equal(t, "Validation Error", response.Message)
	assert.Contains(t, response.Data, "Name cannot be empty")
	assert.Contains(t, response.Data, "Invalid Stock Value, must be a number and greater than 0")
	assert.Contains(t, response.Data, "Invalid Price Value, must be a number and greater than 0")
}

func TestProductService_FindByID_Success(t *testing.T){
	mockRepo := new(MockProductRepository)
	services := services.NewProductService(mockRepo)

	mockProduct := domain.Product{Name: "product 1", Stock: 10, Price: 1000}
	mockRepo.On("FindByID", 1).Return(mockProduct, nil)

	response := services.FindByID(1)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Get Product Successfully", response.Message)
	assert.Equal(t, mockProduct, response.Data)
	mockRepo.AssertExpectations(t)
}

func TestProductService_FindByID_NotFound(t *testing.T){
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockRepo.On("FindByID", 1).Return(domain.Product{}, sql.ErrNoRows)

	response := service.FindByID(1)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, "Product with ID "+strconv.Itoa(1)+" not found", response.Message)
}

func TestProductService_Update_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	// Existing product
	existingProduct := domain.Product{Name: "Old Product", Stock: 5, Price: 50}
	mockRepo.On("FindByID", 1).Return(existingProduct, nil)

	// Updated product data
	productData := map[string]string{
		"name":  "Updated Product",
		"stock": "15",
		"price": "150",
	}

	updatedProduct := domain.Product{Name: "Updated Product", Stock: 15, Price: 150}
	mockRepo.On("Update", updatedProduct).Return(nil)

	response := service.Update(1, productData)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Product updated successfully", response.Message)
	assert.Equal(t, updatedProduct, response.Data)
	mockRepo.AssertExpectations(t)
}

func TestProductService_Update_NotFound(t *testing.T){
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockRepo.On("FindByID", 1).Return(domain.Product{}, sql.ErrNoRows)

	response := service.FindByID(1)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, "Product with ID "+strconv.Itoa(1)+" not found", response.Message)
}

func TestProductService_Delete_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockProduct := domain.Product{ ID: 1 ,Name: "Product1", Stock: 10, Price: 100, IsAvailable: true}
	mockRepo.On("FindByID", 1).Return(mockProduct, nil)
	mockRepo.On("Delete", 1).Return(nil)

	response := service.Delete(1)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Product deleted successfully", response.Message)
}

func TestProductService_Delete_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := services.NewProductService(mockRepo)

	mockRepo.On("FindByID", 1).Return(domain.Product{}, sql.ErrNoRows)

	response := service.Delete(1)

	assert.Equal(t, http.StatusNotFound, response.Code)
	assert.Equal(t, "Product with ID "+strconv.Itoa(1)+" not found", response.Message)
}