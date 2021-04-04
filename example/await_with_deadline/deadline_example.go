package main

import (
	"context"
	"fmt"
	"time"

	"github.com/0B1t322/async"
)

func main() {
	future := async.Exec(func() {
		method()
	})

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))

	if err := future.Await(ctx); err != nil {
		fmt.Println(err) // context deadline exceeded
	}
}

func method() int {
	time.Sleep(2*time.Second)
	return 0
}