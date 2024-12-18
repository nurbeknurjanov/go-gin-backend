package manuals

import (
	"context"
	"fmt"
	"time"
)

func listenChannels(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done")
			return
		}
	}
}
func RunBuffer() {
	root := context.Background()
	ctxValue := context.WithValue(root, "q", "123123")
	ctx, cancel := context.WithTimeout(ctxValue, 5*time.Second)
	stop := context.AfterFunc(ctx, func() {
		fmt.Println("Before done")
	})

	go listenChannels(ctx)

	_, _ = cancel, stop
	/*cancel()
	stop()*/

	fmt.Println("ctx", ctx.Value("q"))
}
