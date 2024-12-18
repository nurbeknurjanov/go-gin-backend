package manuals

import (
	"fmt"
	"time"
)

func runChannel1(i int, ch chan<- int, closeChannel bool) {
	ch <- i
	if closeChannel {
		close(ch)
	}
}
func RunLoops() {
	ch := make(chan int)

	go runChannel1(1, ch, true)

	loop := true
outerLoop:
	for loop {
		time.Sleep(1 * time.Second)
		select {
		case i, ok := <-ch:
			fmt.Println("i->", i)
			if !ok {
				ch = nil
				//loop = false
				break outerLoop
			}
		}

		if ch == nil {
			//break
			//return
		}
	}

	fmt.Println("done")
}

//break works for for i:=0 i<10 i++ too
