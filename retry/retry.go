package retry

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

func Example() {
	fileName := "notExist.go"
	openFile := func() error {
		_, err := os.Open(fileName)
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	go Retry(openFile, RetryInterval(1*time.Second), RetryTimes(100), Context(ctx)) //nolint:errcheck
	time.Sleep(10 * time.Second)
	cancel()
}

const (
	DEFAULT_RETRY_INTERVAL = 3 * time.Second
	DEFAULT_RETRY_TIMES    = 3
)

type RetryConfig struct {
	Context       context.Context
	RetryInterval time.Duration
	RetryTimes    uint
}

type RetryFun func() error
type RetryOption func(*RetryConfig)

func Retry(retryFunc RetryFun, opt ...RetryOption) (err error) {
	config := &RetryConfig{
		Context:       context.TODO(),
		RetryTimes:    DEFAULT_RETRY_TIMES,
		RetryInterval: DEFAULT_RETRY_INTERVAL,
	}
	for _, opt := range opt {
		opt(config)
	}
	for i := uint(0); i < config.RetryTimes; i++ {
		err = retryFunc()
		if err != nil {
			select {
			case <-time.After(config.RetryInterval):
				fmt.Printf("Retry times %d\n", i)
			case <-config.Context.Done():
				fmt.Println("Retry cancelled")
				return
			}
		} else {
			return nil
		}
	}
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
	return
}

func Context(ctx context.Context) RetryOption {
	return func(config *RetryConfig) {
		config.Context = ctx
	}
}

func RetryTimes(retryTimes uint) RetryOption {
	return func(config *RetryConfig) {
		config.RetryTimes = retryTimes
	}
}

func RetryInterval(retryInterval time.Duration) RetryOption {
	return func(config *RetryConfig) {
		config.RetryInterval = retryInterval
	}
}
