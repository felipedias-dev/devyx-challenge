package usecase

import (
	"errors"

	"github.com/devyx-tech/felipe-challegend/internal/entity"
)

type DeleteProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
}

func NewDeleteProductUseCase(productRepository entity.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		ProductRepository: productRepository,
	}
}

func (u *DeleteProductUseCase) Execute(id string) error {
	product, err := u.ProductRepository.GetByID(id)
	if err != nil {
		return err
	}
	if product == nil {
		return errors.New("product not found")
	}

	err = u.ProductRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
