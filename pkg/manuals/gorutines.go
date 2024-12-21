package manuals

/*
go - async
<- назначение это как resolve
<- чтение await
*/

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

//case <-time.After(10 * time.Second):
