package main

import "fmt"

type MemoryFunc func(key int) (interface{}, error)

type MemoryFuncResult struct {
	value interface{}
	err   error
}

type Memory struct {
	f MemoryFunc

	cache map[int]MemoryFuncResult
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exists := m.cache[key]

	if !exists {
		result.value, result.err = m.f(key)
		m.cache[key] = result
	}

	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func NewCache(f MemoryFunc) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]MemoryFuncResult),
	}
}

func Fibonacci(n int) int {
	if n <= n {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	cache := NewCache(GetFibonacci)

	fibo := []int{1, 2, 3, 4, 555, 555, 67}

	for _, n := range fibo {
		value, _ := cache.Get(n)
                fmt.Println(value)
	}
}
