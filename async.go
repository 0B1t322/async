package async

import "context"

// Future interface has the method signature for await
type Future interface {
	Await(ctx context.Context) error
}

type future struct {
	await func(ctx context.Context) error
}

func (f *future) Await(ctx context.Context) error {
	return f.await(ctx)
}

func Exec(f func()) Future {
	done := make(chan struct{})

	go func() {
		defer close(done)
		f()
		done <- struct{}{}
	}()

	return &future{
		await: func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-done:
				return nil
			}
		},
	}
}