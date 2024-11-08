package utils

import (
	"reflect"
	"testing"

	"github.com/devyx-tech/felipe-challegend/internal/usecase/dto"
)

func TestSortProductsByPrice(t *testing.T) {
	tests := []struct {
		name     string
		given    []dto.ListProductsOutputDTO
		expected []dto.ListProductsOutputDTO
	}{
		{
			name: "Success with one product",
			given: []dto.ListProductsOutputDTO{
				{ID: "1", Name: "Product 1", Price: 100},
			},
			expected: []dto.ListProductsOutputDTO{
				{ID: "1", Name: "Product 1", Price: 100},
			},
		},
		{
			name: "Success with two given",
			given: []dto.ListProductsOutputDTO{
				{ID: "1", Name: "Product 1", Price: 100},
				{ID: "2", Name: "Product 2", Price: 50},
			},
			expected: []dto.ListProductsOutputDTO{
				{ID: "2", Name: "Product 2", Price: 50},
				{ID: "1", Name: "Product 1", Price: 100},
			},
		},
		{
			name: "Success with three given",
			given: []dto.ListProductsOutputDTO{
				{ID: "1", Name: "Product 1", Price: 100},
				{ID: "2", Name: "Product 2", Price: 50},
				{ID: "3", Name: "Product 3", Price: 75},
			},
			expected: []dto.ListProductsOutputDTO{
				{ID: "2", Name: "Product 2", Price: 50},
				{ID: "3", Name: "Product 3", Price: 75},
				{ID: "1", Name: "Product 1", Price: 100},
			},
		},
		{
			name: "Success with 10 given",
			given: []dto.ListProductsOutputDTO{
				{ID: "1", Name: "Product 1", Price: 100},
				{ID: "2", Name: "Product 2", Price: 50},
				{ID: "3", Name: "Product 3", Price: 75},
				{ID: "4", Name: "Product 4", Price: 25},
				{ID: "5", Name: "Product 5", Price: 150},
				{ID: "6", Name: "Product 6", Price: 200},
				{ID: "7", Name: "Product 7", Price: 10},
				{ID: "8", Name: "Product 8", Price: 300},
				{ID: "9", Name: "Product 9", Price: 500},
				{ID: "10", Name: "Product 10", Price: 1},
			},
			expected: []dto.ListProductsOutputDTO{
				{ID: "10", Name: "Product 10", Price: 1},
				{ID: "7", Name: "Product 7", Price: 10},
				{ID: "4", Name: "Product 4", Price: 25},
				{ID: "2", Name: "Product 2", Price: 50},
				{ID: "3", Name: "Product 3", Price: 75},
				{ID: "1", Name: "Product 1", Price: 100},
				{ID: "5", Name: "Product 5", Price: 150},
				{ID: "6", Name: "Product 6", Price: 200},
				{ID: "8", Name: "Product 8", Price: 300},
				{ID: "9", Name: "Product 9", Price: 500},
			},
		},
		{
			name: "Success with 20 given",
			given: []dto.ListProductsOutputDTO{
				{ID: "1", Name: "Product 1", Price: 100},
				{ID: "2", Name: "Product 2", Price: 4500},
				{ID: "3", Name: "Product 3", Price: 75},
				{ID: "4", Name: "Product 4", Price: 5500},
				{ID: "5", Name: "Product 5", Price: 150},
				{ID: "6", Name: "Product 6", Price: 2000},
				{ID: "7", Name: "Product 7", Price: 200},
				{ID: "8", Name: "Product 8", Price: 2500},
				{ID: "9", Name: "Product 9", Price: 10},
				{ID: "10", Name: "Product 10", Price: 5000},
				{ID: "11", Name: "Product 11", Price: 1000},
				{ID: "12", Name: "Product 12", Price: 500},
				{ID: "13", Name: "Product 13", Price: 1},
				{ID: "14", Name: "Product 14", Price: 4000},
				{ID: "15", Name: "Product 15", Price: 1500},
				{ID: "16", Name: "Product 16", Price: 300},
				{ID: "17", Name: "Product 17", Price: 3000},
				{ID: "18", Name: "Product 18", Price: 50},
				{ID: "19", Name: "Product 19", Price: 3500},
				{ID: "20", Name: "Product 20", Price: 25},
			},
			expected: []dto.ListProductsOutputDTO{
				{ID: "13", Name: "Product 13", Price: 1},
				{ID: "9", Name: "Product 9", Price: 10},
				{ID: "20", Name: "Product 20", Price: 25},
				{ID: "18", Name: "Product 18", Price: 50},
				{ID: "3", Name: "Product 3", Price: 75},
				{ID: "1", Name: "Product 1", Price: 100},
				{ID: "5", Name: "Product 5", Price: 150},
				{ID: "7", Name: "Product 7", Price: 200},
				{ID: "16", Name: "Product 16", Price: 300},
				{ID: "12", Name: "Product 12", Price: 500},
				{ID: "11", Name: "Product 11", Price: 1000},
				{ID: "15", Name: "Product 15", Price: 1500},
				{ID: "6", Name: "Product 6", Price: 2000},
				{ID: "8", Name: "Product 8", Price: 2500},
				{ID: "17", Name: "Product 17", Price: 3000},
				{ID: "19", Name: "Product 19", Price: 3500},
				{ID: "14", Name: "Product 14", Price: 4000},
				{ID: "2", Name: "Product 2", Price: 4500},
				{ID: "10", Name: "Product 10", Price: 5000},
				{ID: "4", Name: "Product 4", Price: 5500},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortProductsByPrice(tt.given); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SortProductsByPrice() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
