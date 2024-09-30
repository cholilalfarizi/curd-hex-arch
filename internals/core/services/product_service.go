package services

import (
	model "crud-hex/internals/core/domain"
	"crud-hex/internals/core/ports"
	"crud-hex/internals/utils"
	"database/sql"
	"net/http"
	"strconv"
)

type ProductService struct {
	productRepository ports.IProductRepository
}


func NewProductService(repository ports.IProductRepository) *ProductService{
	return &ProductService{
		productRepository: repository,
	}
}

func (s *ProductService) FindAll() utils.ServiceResponse {
	products, err := s.productRepository.FindAll()
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to fetch products",
			Data:    err.Error(),
		}
	}
	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Get Products successfully",
		Data:    products,
	}
}

func (s *ProductService) Create(productData map[string]string) utils.ServiceResponse {
	name := productData["name"]
	stockStr := productData["stock"]
	priceStr := productData["price"]

	// Validate inputs
	var arrErrors []string

	if name == "" {
		arrErrors = append(arrErrors, "Name cannot be empty")
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil || stock < 0 {
		arrErrors = append(arrErrors, "Invalid Stock Value, must be a number and greater than 0")
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil || price < 0 {
		arrErrors = append(arrErrors, "Invalid Price Value, must be a number and greater than 0")
	}

	if len(arrErrors) > 0 {
		return utils.ServiceResponse{
			Code:    http.StatusBadRequest,
			Message: "Validation error",
			Data:    arrErrors,
		}
	}

	// Create the product
	product := &model.Product{
		Name:  name,
		Stock: stock,
		Price: price,
		IsAvailable: true,

	}

	

	err = s.productRepository.Create(product)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error creating product",
			Data:    err.Error(),
		}
	}

	return utils.ServiceResponse{
		Code:    http.StatusCreated,
		Message: "Product added successfully",
		Data:    product,
	}
}

func (s *ProductService) FindByID(id int) utils.ServiceResponse {
	// Fetch product by ID from repository
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusNotFound,
			Message: "Product with ID " + strconv.Itoa(id) + " not found",
			Data:    nil,
		}
	}

	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Product fetched successfully",
		Data:    product,
	}
}

func (s *ProductService) Update(id int, productData map[string]string) utils.ServiceResponse {
	// Cari produk berdasarkan ID
	existingProduct, err := s.productRepository.FindByID(id)
	if err != nil {
		// Jika produk tidak ditemukan, berikan response Not Found
		if err == sql.ErrNoRows {
			return utils.ServiceResponse{
				Code:    http.StatusNotFound,
				Message: "Product with ID " + strconv.Itoa(id) + " not found",
				Data:    nil,
			}
		}
		// Jika ada error lain saat fetch product
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error fetching product with ID " + strconv.Itoa(id),
			Data:    err.Error(),
		}
	}

	// Ambil data input dari productData
	name := productData["name"]
	stockStr := productData["stock"]
	priceStr := productData["price"]

	// Perbarui nama jika tidak kosong
	if name != "" {
		existingProduct.Name = name
	}

	// Perbarui stock jika tidak kosong dan valid
	if stockStr != "" {
		stock, err := strconv.Atoi(stockStr)
		if err != nil || stock < 0 {
			return utils.ServiceResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid Stock Value, must be a number and greater than 0",
				Data:    nil,
			}
		}
		existingProduct.Stock = stock
	}

	// Perbarui harga jika tidak kosong dan valid
	if priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err != nil || price < 0 {
			return utils.ServiceResponse{
				Code:    http.StatusBadRequest,
				Message: "Invalid Price Value, must be a number and greater than 0",
				Data:    nil,
			}
		}
		existingProduct.Price = price
	}

	// Update produk menggunakan repository
	err = s.productRepository.Update(existingProduct)
	if err != nil {
		// Jika terjadi error saat update, berikan response Internal Server Error
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error updating product with ID " + strconv.Itoa(id),
			Data:    err.Error(),
		}
	}

	// Jika berhasil, berikan response sukses
	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Product updated successfully",
		Data:    existingProduct,
	}
}


func (s *ProductService) Delete(id int) utils.ServiceResponse {
	// Cek apakah produk dengan ID tersebut ada di database
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		// Jika tidak ditemukan, berikan response bahwa produk tidak ada
		if err == sql.ErrNoRows {
			return utils.ServiceResponse{
				Code:    http.StatusNotFound,
				Message: "Product with ID " + strconv.Itoa(id) + " not found",
				Data:    nil,
			}
		}
		// Jika ada error lain saat fetch product
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error checking product existence",
			Data:    err.Error(),
		}
	}

	// Jika produk ada, lanjutkan dengan proses delete (soft delete)
	err = s.productRepository.Delete(product.ID)
	if err != nil {
		return utils.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error deleting product",
			Data:    err.Error(),
		}
	}

	// Return success response
	return utils.ServiceResponse{
		Code:    http.StatusOK,
		Message: "Product deleted successfully",
		Data:    nil,
	}
}




