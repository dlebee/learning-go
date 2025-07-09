package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Tell WaitGroup this worker is done when function exits

	fmt.Printf("Worker %d: started\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("Worker %d: finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(5)

	go worker(1, &wg)
	go worker(2, &wg)
	go worker(3, &wg)
	go worker(4, &wg)
	go worker(5, &wg)

	fmt.Println("Main: waiting for workers to finish...")
	wg.Wait()

	fmt.Println("Main: all workers completed. Exiting.")
}
