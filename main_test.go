package main

import (
	"testing"

	"github.com/stretchr/testify/assert" //  использую testify для проверок
)

func TestGenerateRandomElements(t *testing.T) {
	//  табличные тесты
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"10 элементов", 10, 10},
		{"нулевой размер", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			assert.Equal(t, tt.expected, len(data))
			for _, v := range data {
				assert.GreaterOrEqual(t, v, 0)
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	// табличные тесты + assert
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"возрастающий", []int{1, 2, 3, 4, 5}, 5},
		{"один элемент", []int{10}, 10},
		{"смешанные", []int{100, 99, 1, 55}, 100},
		{"пустой", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maximum(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaxChunks(t *testing.T) {
	// табличные тесты с аналогичными кейсами, как для maximum
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"обычный", []int{5, 9, 1, 3, 20, 4, 15, 10, 8, 6}, 20},
		{"один элемент", []int{42}, 42},
		{"пустой", []int{}, 0},
		{"равные значения", []int{7, 7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxChunks(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
