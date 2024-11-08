package usecase

import (
	"net/url"

	"github.com/devyx-tech/felipe-challegend/internal/entity"
	"github.com/devyx-tech/felipe-challegend/internal/usecase/dto"
	"github.com/devyx-tech/felipe-challegend/pkg/utils"
)

type ListProductsUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewListProductsUseCase(productRepository entity.ProductRepositoryInterface) *ListProductsUseCase {
	return &ListProductsUseCase{
		ProductRepository: productRepository,
	}
}

func (u *ListProductsUseCase) Execute(queryParams url.Values) ([]dto.ListProductsOutputDTO, error) {
	products, err := u.ProductRepository.List(queryParams)
	if err != nil {
		return nil, err
	}

	var output []dto.ListProductsOutputDTO
	for _, product := range products {
		output = append(output, dto.ListProductsOutputDTO{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	outputSorted := utils.SortProductsByPrice(output)

	return outputSorted, nil
}
