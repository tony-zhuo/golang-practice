package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	Test(ctx)
}

func Test(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(i)
			i++
			time.Sleep(1 * time.Second)
		}
	}
}
