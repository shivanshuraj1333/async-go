package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup, cancel context.CancelFunc, fail bool) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped\n", id)
			return
		default:
			if fail {
				fmt.Printf("Worker %d failed, cancelling others\n", id)
				cancel() // cancel others
				return
			}
			fmt.Printf("Worker %d is working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		// simulate failure in worker 1
		go worker(ctx, i, &wg, cancel, i == 1)
	}

	wg.Wait()
	fmt.Println("All workers done or cancelled.")
}
