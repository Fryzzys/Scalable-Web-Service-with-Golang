package main

import (
	"fmt"
	"sync"
)

func routineTest(wg *sync.WaitGroup, i int, data []string, c chan string) {
	defer wg.Done()
	c <- fmt.Sprintf("%v %d", data, i)
}

func main() {
	var wg sync.WaitGroup
	var c = make(chan string)
	var dataCoba = []string{"coba1", "coba2", "coba3"}
	var dataBisa = []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go routineTest(&wg, i, dataCoba, c)
		go routineTest(&wg, i, dataBisa, c)
	}

	for i := 1; i <= 8; i++ {
		fmt.Println(<-c)
	}
	wg.Wait()
}
