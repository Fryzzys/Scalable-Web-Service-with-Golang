package main

import (
	"fmt"
	"sync"
)

func routineCoba(wg *sync.WaitGroup, mtx *sync.Mutex, i int, data []string, c1 *chan string) {
	defer wg.Done()
	*c1 <- fmt.Sprintf("%v %d", data, i)
}

func routineBisa(wg *sync.WaitGroup, mtx *sync.Mutex, i int, data []string, c2 *chan string) {
	defer wg.Done()
	*c2 <- fmt.Sprintf("%v %d", data, i)
}

func routineCobaBisa(wg *sync.WaitGroup, mtx *sync.Mutex, c1 *chan string, c2 *chan string, cJoin *chan string) {
	defer wg.Done()
	mtx.Lock()
	*cJoin <- <-*c1
	*cJoin <- <-*c2
	mtx.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mtx sync.Mutex
	var cJoin = make(chan string, 2)
	var c1 = make(chan string)
	var c2 = make(chan string)
	var data1 = []string{"coba1", "coba2", "coba3"}
	var data2 = []string{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go routineCoba(&wg, &mtx, i, data1, &c1)
		go routineBisa(&wg, &mtx, i, data2, &c2)
	}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		go routineCobaBisa(&wg, &mtx, &c1, &c2, &cJoin)
	}

	for i := 1; i <= 8; i++ {
		fmt.Println(<-cJoin)
	}
	
	go func() {
		wg.Wait()
		close(cJoin)
	}()
}
