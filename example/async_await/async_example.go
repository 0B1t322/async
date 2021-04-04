package main

import (
	"context"
	"fmt"
	"time"

	"github.com/0B1t322/async"
)

func main() {
	var (
		num int
		err error
	)

	future := async.Exec(func() {
		num, err = method()
	})

	if err := future.Await(context.Background()); err != nil {
		panic(err)
	}

	fmt.Printf("num: %v, error: %v\n", num, err) // num: 20, err: <nil>
}

func method() (int, error) {
	time.Sleep(time.Second)
	return 20, nil
}