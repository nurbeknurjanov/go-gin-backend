package manuals

import (
	"context"
	"fmt"
	"time"
)

func RunTest() {
	bctx := context.Background()
	ctxValue := context.WithValue(bctx, "test", "testValue")
	ctxCancel, cancel := context.WithCancel(bctx)
	_ = cancel

	deadlineCtx, cancelDead := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancelDead() // Освобождаем ресурсы

	go func() {
		//time.Sleep(1 * time.Second)
		select {
		case <-ctxCancel.Done(): //подслушивает
			fmt.Println("Operation cancelled", ctxValue.Value("test"))
		case <-deadlineCtx.Done():
			fmt.Println("Deadline exceeded:", deadlineCtx.Err())
		case <-time.After(3 * time.Second):
			fmt.Println("3 second passed")
		}

	}()

	//cancel()
	time.Sleep(5 * time.Second)
}
