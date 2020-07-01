package main

import (
	"context"
	"fmt"
	"time"
)

func func2(ctx context.Context) {
	// 从ctx中通过key取出值，并检测是不是string类型
	if v, ok := ctx.Value("thisiskey").(string); ok {
		fmt.Println("v in func2: ", v)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("routine fun2 interrupt... err: ", ctx.Err())
			return
		}
		fmt.Println("in func2 routine.")
	}
}

func TestContext() {
	messages := make(chan int, 10)

	// producer
	for i := 0; i < 10; i++ {
		messages <- i
	}

	key1 := "thisiskey"
	value1 := "thisisvalue"
	ctxv := context.WithValue(context.Background(), key1, value1)
	ctx, cancelThis := context.WithTimeout(ctxv, 5*time.Second)
	// ctx2, cancel2 := context.WithCancel(ctxv)

	// consumer
	go func(ctx context.Context) {
		go func2(ctx)

		ticker := time.NewTicker(1 * time.Second) // 每1秒钟会触发一次ticker
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("child process interrupt...")
				return
			default:
				fmt.Printf("send message: %d\n", <-messages)
			}
		}
	}(ctx)

	defer close(messages)
	defer cancelThis() // 在context.WithTimeout使用时，不是在这里取消ctx，而是释放资源。

	fmt.Println("main is waiting for ctx.Done.")
	select {
	case <-ctx.Done():
		// time.Sleep(1 * time.Second)
		fmt.Println("main process exit!")
	}

	return
}
