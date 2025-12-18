package main

import (
	"fmt"
	"sync"
	"time"
)

func work(t <-chan int, wg *sync.WaitGroup, index int) {
	// dung wait khi done
	defer wg.Done()
	for task := range t {
		fmt.Printf("Work %d\t with task %d\n", index, task)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// khoi tao 1 waitinggroup
	var wg sync.WaitGroup

	// khoi tao bien
	var (
		numberMember = 5
		numberTask   = 100
	)

	// khoi tao 1 chanel (buffer channel)
	tasks := make(chan int, 10)

	// khoi tao cac worker
	for i := 1; i <= numberMember; i++ {
		wg.Add(1)
		go work(tasks, &wg, i) // tao 1 gorounetine tach khoi gorountine main
	}

	// gui task
	for i := 1; i <= numberTask; i++ {
		tasks <- i
	}

	close(tasks) // thong bao task da done
	wg.Wait()

	fmt.Println("All tasks done")
}
