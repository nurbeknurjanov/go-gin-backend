package manuals

import (
	"fmt"
	"sync"
)

/*func RunTest() {
	var myMap = map[string]int{}

	mu := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			//fmt.Println("out", i)
			mu.Lock()
			if i == 999 {
				fmt.Println("in-------------", i)
			}
			myMap["A"] = i
			mu.Unlock()
		}()
	}

	time.Sleep(3 * time.Second)
}*/

func RunTest() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("out", i)
		}()
	}

	wg.Wait()
	fmt.Println("all done")
}
