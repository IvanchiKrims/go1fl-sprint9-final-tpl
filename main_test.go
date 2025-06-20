package main

import "testing"

// Тестирует генерацию случайных чисел
func TestGenerateRandomElements(t *testing.T) {
	data := generateRandomElements(10)
	if len(data) != 10 {
		t.Errorf("Ожидалось 10 элементов, получено %d", len(data))
	}
	for _, v := range data {
		if v < 0 {
			t.Errorf("Найдено отрицательное значение: %d", v)
		}
	}

	// Проверка генерации при нулевом размере
	empty := generateRandomElements(0)
	if len(empty) != 0 {
		t.Errorf("Ожидался пустой слайс, получено %d", len(empty))
	}
}

// Тестирует функцию maximum на разных входных данных
func TestMaximum(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{10}, 10},
		{[]int{100, 99, 1, 55}, 100},
		{[]int{}, 0}, // пустой слайс
	}

	for _, test := range tests {
		result := maximum(test.input)
		if result != test.expected {
			t.Errorf("maximum(%v) = %d; ожидалось %d", test.input, result, test.expected)
		}
	}
}

// Тестирует функцию maxChunks
func TestMaxChunks(t *testing.T) {
	data := []int{5, 9, 1, 3, 20, 4, 15, 10, 8, 6}
	expected := 20
	result := maxChunks(data)
	if result != expected {
		t.Errorf("maxChunks(%v) = %d; ожидалось %d", data, result, expected)
	}
}
