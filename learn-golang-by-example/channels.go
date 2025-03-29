package main

import (
	"fmt"
	"sync"
	"time"
)

func sum(s []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, v := range s {
		sum += v
	}

	c <- sum
}

func main() {
	// list := []int{7, 3, 2, 1}
	// c := make(chan int)
	// var sync sync.WaitGroup
	// sync.Add(2)
	// go sum(list[:len(list)/2], c, &sync)
	// go sum(list[len(list)/2:], c, &sync)

	// go func() {
	// 	sync.Wait()
	// 	close(c)
	// }()

	// for x := range c {
	// 	// x := <-c
	// 	fmt.Println(x)
	// }

	test_lockmutext()

}

// func test_channel() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 1
// 	}()

// 	time.Sleep(time.Second * 3)
// 	fmt.Println(<-ch)
// }

func fibonaci_test_range_close_buffer(ch chan int, n int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y

	}

	close(ch)
}

func test() {
	ch := make(chan int)
	go fibonaci_test_range_close_buffer(ch, 10)
	for i := range ch {
		fmt.Println(i)
	}
}

func process(ch, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Cancel")
			return
		}
	}

}

func test_select() {
	ch, quit := make(chan int), make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}

		quit <- 1
	}()

	process(ch, quit)
}

func test_lockmutext() {
	n := 0
	lock := sync.Mutex{}

	for i := 0; i < 5; i++ {
		go func() {
			for k := 1; k <= 3000; k++ {
				lock.Lock()
				n++
				lock.Unlock()
			}
		}()
	}
	time.Sleep(time.Second * 10)

	fmt.Println(n)
}

type MapConcurrency struct {
	store map[string]int
	mu    sync.Mutex
}

func (m *MapConcurrency) Increase(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.store[key]; !ok {
		return false
	}

	m.store[key] += 1
	return true
}

func (m *MapConcurrency) Get(key string) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.store[key]; !ok {
		return 0
	}

	return m.store[key]
}
func test_concurrency_map() {

}
