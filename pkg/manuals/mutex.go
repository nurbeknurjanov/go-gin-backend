package manuals

import "fmt"

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

/*func RunTest() {
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
}*/

func merge(channels ...<-chan int) (mergedChannel chan int) {
	//var mergedChannel = make(chan int)

	for _, channel := range channels {
		for num := range channel {
			mergedChannel <- num
		}
	}
	return
}

func RunTest() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	channel3 := make(chan int)

	go func() {
		defer close(channel1)
		channel1 <- 1
		channel1 <- 11
	}()
	go func() {
		defer close(channel2)
		channel2 <- 2
	}()
	go func() {
		defer close(channel3)
		channel3 <- 3
		channel3 <- 30
		channel3 <- 33
	}()

	channels := merge(channel1, channel2, channel3)
	for num := range channels {
		fmt.Println(num)
	}

}
