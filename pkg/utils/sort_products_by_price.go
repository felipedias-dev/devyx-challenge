package utils

import (
	"math/rand"

	"github.com/devyx-tech/felipe-challegend/internal/usecase/dto"
)

func SortProductsByPrice(products []dto.ListProductsOutputDTO) []dto.ListProductsOutputDTO {
	if len(products) < 2 {
		return products
	}

	left, right := 0, len(products)-1

	pivot := rand.Int() % len(products)

	products[pivot], products[right] = products[right], products[pivot]

	for i := range products {
		if products[i].Price < products[right].Price {
			products[i], products[left] = products[left], products[i]
			left++
		}
	}

	products[left], products[right] = products[right], products[left]

	SortProductsByPrice(products[:left])
	SortProductsByPrice(products[left+1:])

	return products
}
