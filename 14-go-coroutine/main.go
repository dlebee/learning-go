package main

import (
	"fmt"
	"time"
)

func doWork() {
	fmt.Println("DOINGGGGGG IN BACKGROUND")
}

func main() {
	go doWork()
	fmt.Println("main thread 1...")
	time.Sleep(1 * time.Second)
	fmt.Println("main thread 2...")
	time.Sleep(1 * time.Second)
	fmt.Println("main thread 3...")
}
