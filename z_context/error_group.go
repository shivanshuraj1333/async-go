package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"time"
)

func main() {
	ctx := context.Background()
	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < 3; i++ {
		i := i // capture loop variable
		g.Go(func() error {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d stopped due to cancellation\n", i)
					return ctx.Err()
				default:
					if i == 1 {
						return fmt.Errorf("Worker %d failed", i)
					}
					fmt.Printf("Worker %d is working...\n", i)
					time.Sleep(500 * time.Millisecond)
				}
			}
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("All workers exited.")
}
