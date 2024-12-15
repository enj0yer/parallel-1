package processing

import (
	"fmt"
	"math"
	"sync"
)

type Converter[T int, E any] struct {
	items   []T
	applier func(T) (E, error)
}

func (c *Converter[T, E]) ProcessSequentially() ([]E, error) {
	result, err := c.processChunk(c.items)
	if err != nil {
		return nil, fmt.Errorf("error while sequentially processing data: %v", err)
	}
	return result, nil
}

func (c *Converter[T, E]) ProcessSimultaneously(threads int) ([]E, error) {
	var result []E = make([]E, len(c.items))
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	chunkSize := int(math.Ceil(float64(len(c.items)) / float64(threads)))

	for i := 0; i < threads; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if start > len(c.items) {
			start = len(c.items)
		}
		if end > len(c.items) {
			end = len(c.items)
		}

		wg.Add(1)
		go func(start int, end int) {
			defer wg.Done()
			buffer, err := c.processChunk(c.items[start:end])
			if err != nil {
				fmt.Printf("error while processing chunk from %d to %d: %v", start, end, err)
				return
			}
			mutex.Lock()
			copy(result[start:end], buffer)
			mutex.Unlock()
		}(start, end)
	}
	wg.Wait()
	return result, nil
}

func (c *Converter[T, E]) processChunk(items []T) ([]E, error) {
	var result []E = make([]E, len(items))

	for i, item := range items {
		res, err := c.applier(item)
		if err != nil {
			return nil, fmt.Errorf("error while converting value %v: %w", item, err)
		}
		result[i] = res
	}
	return result, nil
}

func NewConverter[T int, E any](items []T, applier func(T) (E, error)) *Converter[T, E] {
	return &Converter[T, E]{items: items, applier: applier}
}
