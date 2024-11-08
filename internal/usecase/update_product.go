package usecase

import (
	"github.com/devyx-tech/felipe-challegend/internal/entity"
	"github.com/devyx-tech/felipe-challegend/internal/usecase/dto"
)

type UpdateProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewUpdateProductUseCase(productRepository entity.ProductRepositoryInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *UpdateProductUseCase) Execute(input dto.UpdateProductInputDTO) error {
	product := &entity.Product{
		ID:    input.ID,
		Name:  input.Name,
		Price: input.Price,
	}

	err := u.ProductRepository.Update(product)
	if err != nil {
		return err
	}

	return nil
}
