package manuals

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

/*
func me(i int, wg *sync.WaitGroup) {
	m["A"] = i
	fmt.Println(m)
	wg.Done()
}
func RunBuffer() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go me(i, &wg)
		wg.Wait()
	}
}
*/

/*func merge(channels ...<-chan int) <-chan int {
	var mergedChannel = make(chan int)

	go func() {
		defer close(mergedChannel)
		for _, channel := range channels {
			for num := range channel {
				mergedChannel <- num
			}
		}
	}()

	return mergedChannel
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

}*/

/*var count int64 = 0
func me() {
	//count++
	//atomic.AddInt64(&count, 1)
	//atomic.StoreInt64(&count, atomic.LoadInt64(&count)+1) //not work
}
func RunBuffer() {
	for i := 0; i < 1000; i++ {
		go me()
	}
	time.Sleep(time.Second * 1)
	fmt.Println(count)
}*/
