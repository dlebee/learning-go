package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	fmt.Println("Worker: started")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker: received cancel signal -", ctx.Err())
			return
		default:
			fmt.Println("Worker: working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go doWork(ctx)

	fmt.Println("Main: running. Will cancel worker after 5 seconds...")
	time.Sleep(5 * time.Second)

	fmt.Println("Main: sending cancel...")
	cancel() // You control when to stop the goroutine

	// Optional tiny delay to ensure the goroutine prints "Stopped"
	time.Sleep(200 * time.Millisecond)

	fmt.Println("Main: done.")
}
