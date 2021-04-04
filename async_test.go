package async_test

import (
	"context"
	"testing"
	"time"

	"github.com/0B1t322/async"
)

func TestFunc_async_method(t *testing.T) {
	var (
		num *int = nil
	)

	method := func() *int {
		num := 2
		return &num
	}

	future := async.Exec(func() {
		num = method()
	})

	if err := future.Await(context.Background()); err != nil {
		t.Log(err)
		t.FailNow()
	}

	if num == nil {
		t.Log("Should be not nil")
		t.FailNow()
	}

}

func TestFunc_async_Deadline(t *testing.T) {
	method := func() *int {
		num := 2
		return &num
	}

	future := async.Exec(func() {
		method()
		time.Sleep(time.Second)
	})

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond))

	if err := future.Await(ctx); err != context.DeadlineExceeded {
		t.Log("Should be deadline")
		t.FailNow()
	}


}