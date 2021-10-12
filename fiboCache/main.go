package main

import (
	"fmt"
	"sync"
)

type MemoryFunc func(key int) (interface{}, error)

type MemoryFuncResult struct {
	value interface{}
	err   error
}

type Memory struct {
	f    MemoryFunc
	lock sync.Mutex

	cache map[int]MemoryFuncResult
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()

	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
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

	fibo := []int{1, 2, 3, 4, 555, 555, 67, 99999999999, 333333, 4444556}

	var wg sync.WaitGroup

	for _, n := range fibo {

		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			value, _ := cache.Get(index)
			fmt.Println(value)
		}(n)

	}

	wg.Wait()
}
