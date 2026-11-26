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

// generateRandomElements возвращает слайс из size случайных целых чисел от 1 до 100.
// Если size меньше или равен нулю, возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	var array []int
	array = make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(100) + 1
	}
	return array
}

// maxChunks делит слайс data на CHUNKS частей, находит максимум каждой части
// в отдельной горутине и возвращает максимум среди них.
// Для пустого слайса возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if data == nil {
		return 0
	}
	maxs := make([]int, CHUNKS)
	chunks := len(data) / CHUNKS
	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		start := i * chunks
		end := start + chunks
		if i == CHUNKS-1 {
			end = len(data)
		}
		wg.Add(1)
		go func(chunk []int, index int) {
			defer wg.Done()
			maxs[index] = maximum(chunk)
		}(data[start:end], i)

	}
	wg.Wait()
	return maximum(maxs)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
