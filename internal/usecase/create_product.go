package usecase

import (
	"github.com/devyx-tech/felipe-challegend/internal/entity"
	"github.com/devyx-tech/felipe-challegend/internal/usecase/dto"
)

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewCreateProductUseCase(productRepository entity.ProductRepositoryInterface) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *CreateProductUseCase) Execute(input dto.CreateProductInputDTO) (dto.CreateProductOutputDTO, error) {
	product := entity.NewProduct(input.Name, input.Price)

	err := u.ProductRepository.Save(product)
	if err != nil {
		return dto.CreateProductOutputDTO{}, err
	}

	return dto.CreateProductOutputDTO{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}
