package manuals

/*
go - async
<- назначение это как resolve
<- чтение await
*/

/*func createCh(n int) chan int {
	ch := make(chan int)
	go func() {
		ch <- n
	}()
	return ch
}
func GoRun() {
	fmt.Println("result", <-createCh(30))
}
*/

/*func RunTest() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for i := range ch {
		fmt.Println(i)
	}
}*/

/*func RunTest() {
	ch := make(chan int)
	go func() {
		defer fmt.Println("Finished")
		for i := range ch {
			fmt.Println(i)
		}
	}()
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //call finished
	time.Sleep(3 * time.Second)
}*/

/*func goOne(ch chan string) {
	<-ch
}
func goTwo(ch chan string) {
	<-ch
}
func RunTest() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	//go goTwo(ch2)

	select {
	case ch1 <- "A":
		fmt.Println("msg11")
	case ch2 <- "B":
		fmt.Println("msg22")
	}
}*/

/*func fillData(data, exit chan int) {
	x := 0
	for {
		select {
		case data <- x:
			x++
		case <-exit:
			fmt.Println("Exit")
			return
		default:
			fmt.Println("Waiting")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func RunTest() {
	data := make(chan int)
	exit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-data)
		}

		exit <- 0
	}()

	fillData(data, exit)
}
*/

/*var a = make(chan int)
go func() {
	fmt.Println(<-a)
}()
a <- 1*/

//case <-time.After(10 * time.Second):

/*func listenChannels(ctx context.Context) {
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
	//cancel()
	//stop()

	fmt.Println("ctx", ctx.Value("q"))
}*/
