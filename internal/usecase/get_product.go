package usecase

import (
	"github.com/devyx-tech/felipe-challegend/internal/entity"
	"github.com/devyx-tech/felipe-challegend/internal/usecase/dto"
)

type GetProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewGetProductUseCase(productRepository entity.ProductRepositoryInterface) *GetProductUseCase {
	return &GetProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *GetProductUseCase) Execute(id string) (dto.GetProductOutputDTO, error) {
	product, err := u.ProductRepository.GetByID(id)
	if err != nil {
		return dto.GetProductOutputDTO{}, err
	}

	output := dto.GetProductOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

	return output, nil
}
