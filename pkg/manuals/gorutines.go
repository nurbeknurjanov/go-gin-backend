package manuals

import "fmt"

/*
go - async
<- назначение это как resolve
<- чтение await
*/

/*func GoRun() {
	ch := make(chan int)
	go func(ch chan int) {
		ch <- 10
		fmt.Println("1 input")
		ch <- 20
		fmt.Println("2 input")
	}(ch)
	fmt.Println("1 output", <-ch)
	fmt.Println("2 second", <-ch)
}*/

/*func GoRun() {
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("After 1")
		ch <- 2
		fmt.Println("After 2")
		ch <- 3
		fmt.Println("After 3")
	}()
	fmt.Println("el1", <-ch)
	fmt.Println("After el1")
	fmt.Println("el2", <-ch)
	fmt.Println("After el2")
	fmt.Println("el3", <-ch)
	fmt.Println("After el3")

	time.Sleep(2 * time.Second)
}*/

/*
	func Sum(n int, ch chan<- int) {
		ch <- n * n
	}

	func GoRun() {
		ch := make(chan int)
		go Sum(3, ch)
		fmt.Println(<-ch)
	}
*/
/*func Sum(ch chan int) {
	n := <-ch
	ch <- n * n
}
func GoRun() {
	ch := make(chan int)
	go Sum(ch) //без этого зависнет
	ch <- 3 //так как не сможет записать
	fmt.Println("result", <-ch)
}*/

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

/*func Listen(ch chan int) {
	defer close(ch)
	for i := range ch {
		fmt.Println(i)
	}
}
func Run() {
	ch := make(chan int)
	go Listen(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}*/

/*func Say(ch chan int) {
	defer close(ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func Run() {
	ch := make(chan int)
	go Say(ch)
	for i := range ch {
		fmt.Println(i)
	}
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

/*func fibonacci(c chan int, enough chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-enough:
			fmt.Println("quit")
			return
		}
	}
}
func RunTest() {
	c := make(chan int)
	enough := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		enough <- true
	}()

	fibonacci(c, enough)
}*/

func fillData(data, exit chan int) {
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
