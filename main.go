package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует слайс случайных чисел заданного размера
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	rand.Seed(time.Now().UnixNano())
	data := make([]int, size)
	for i := range data {
		// убрал ограничение Intn(1_000_000_000) теперь rand.Int() — полный диапазон положительных int
		data[i] = rand.Int()
	}
	return data
}

// maximum возвращает максимальное значение из слайса
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks ищет максимум в слайсе разбивая его на CHUNKS частей и используя горутины
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	results := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		// передаю готовый слайс в горутину а не только индекс i
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		chunk := data[start:end]

		go func(i int, chunk []int) {
			defer wg.Done()
			results[i] = maximum(chunk)
		}(i, chunk)
	}

	wg.Wait()
	return maximum(results)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел...\n", SIZE)
	data := generateRandomElements(SIZE)

	// Однопоточный максимум
	fmt.Println("Ищем максимум в один поток...")
	start1 := time.Now()
	max1 := maximum(data)
	elapsed1 := time.Since(start1).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n", max1, elapsed1)

	// Многопоточный максимум
	fmt.Printf("Ищем максимум в %d потоков...\n", CHUNKS)
	start2 := time.Now()
	max2 := maxChunks(data)
	elapsed2 := time.Since(start2).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d мкс\n", max2, elapsed2)
}
