package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/devyx-tech/felipe-challegend/internal/entity"
	"github.com/devyx-tech/felipe-challegend/internal/usecase"
	"github.com/devyx-tech/felipe-challegend/internal/usecase/dto"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewProductHandler(productRepository entity.ProductRepositoryInterface) *ProductHandler {
	return &ProductHandler{
		ProductRepository: productRepository,
	}
}

// Create product godoc
// @Summary Create a product
// @Description Create a product
// @Tags products
// @Accept json
// @Produce json
// @Param input body dto.CreateProductInputDTO true "Product data"
// @Success 201 {object} dto.CreateProductOutputDTO "Product created"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /products [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateProductInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createProduct := usecase.NewCreateProductUseCase(h.ProductRepository)
	output, err := createProduct.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// List products godoc
// @Summary List products
// @Description List products
// @Tags products
// @Accept json
// @Produce json
// @Param page query string false "Page number"
// @Param limit query string false "Items per page"
// @Success 200 {array} dto.ListProductsOutputDTO "Products found"
// @Failure 500 "Internal server error"
// @Router /products [get]
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	listProducts := usecase.NewListProductsUseCase(h.ProductRepository)
	output, err := listProducts.Execute(queryParams)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Get product by ID godoc
// @Summary Get product by ID
// @Description Get product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} dto.GetProductOutputDTO "Product found"
// @Failure 404 "Product not found"
// @Failure 500 "Internal server error"
// @Router /products/{id} [get]
func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	listProducts := usecase.NewGetProductUseCase(h.ProductRepository)
	output, err := listProducts.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update product godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param input body dto.UpdateProductInputDTO true "Product data"
// @Success 204 "Product updated"
// @Failure 400 "Bad request"
// @Failure 500 "Internal server error"
// @Router /products/{id} [put]
func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var input dto.UpdateProductInputDTO
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	input.ID = id

	updateProduct := usecase.NewUpdateProductUseCase(h.ProductRepository)
	err = updateProduct.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Delete product godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 "Product deleted"
// @Failure 500 "Internal server error"
// @Router /products/{id} [delete]
func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	deleteProduct := usecase.NewDeleteProductUseCase(h.ProductRepository)
	err := deleteProduct.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
