package manuals

import (
	"context"
	"fmt"
	"time"
)

func RunContext() {
	root := context.Background()

	ctxCancel, cancel := context.WithCancel(root)
	_ = cancel

	deadlineCtx, cancelDead := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	defer cancelDead() // Освобождаем ресурсы

	go func() {
		//time.Sleep(1 * time.Second)
		select {
		case <-ctxCancel.Done(): //подслушивает
			fmt.Println("Operation cancelled")
		case <-deadlineCtx.Done():
			fmt.Println("Deadline exceeded:", deadlineCtx.Err())
		case <-time.After(3 * time.Second):
			fmt.Println("3 second passed")
		}

	}()

	//cancel()
	time.Sleep(5 * time.Second)
}

func listenChannels(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}
	}
}
func MainContext() {
	root := context.Background()
	ctxValue := context.WithValue(root, "q", "123123")
	ctxTimeout, cancel := context.WithTimeout(ctxValue, 5*time.Second)
	stop := context.AfterFunc(ctxTimeout, func() {
		fmt.Println("Before done")
	})

	go listenChannels(ctxTimeout)

	_, _ = cancel, stop
	/*cancel()
	stop()*/

	time.Sleep(10 * time.Second)
	fmt.Println("ctxTimeout", ctxTimeout.Value("q"))
}
