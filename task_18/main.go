package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.RWMutex
	Value int
}

func (c *Counter) Increment() {
	c.RLock()
	defer c.RUnlock()

	c.Value++
}

func main() {
	// В целом делаем почти всё то же самое, что и в предыдущих тасках на конкуретность
	var wg sync.WaitGroup
	counter := &Counter{Value: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value: ", counter.Value)
}
