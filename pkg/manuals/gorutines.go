package manuals

/*
go - async
<- назначение это как resolve
<- чтение await

// сперва работает основной поток
// 10-20 нанокоманд forовских по крайней мере успеет
// for запуск горутин, они будут асинхронно запускаться, без учета нано
// select {} для блокирования
// канал если передать как параметр, он дублируется, не ссылка
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

/*
func merge2(chans ...<-chan int) <-chan int {
	combined := make(chan int)
	exit := make(chan int8)

	for _, ch := range chans {
		go func(mx <-chan int) {
			for n := range mx {
				combined <- n
			}
			exit <- 0
		}(ch)
	}

	go func() {
		defer close(combined)

		for range chans {
			<-exit
		}
	}()

	return combined
}
func merge(channels ...<-chan int) chan int {
	out := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, c := range channels {
		go func(cloneC <-chan int) {
			//вечное подслушивание
			for v := range cloneC {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func asChan(numbers ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range numbers {
			c <- v
		}
		close(c)
	}()
	return c
}

func main() {
	a := asChan(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := asChan(11, 12, 13, 14, 15, 16, 17, 18, 19, 20)
	c := asChan(21, 22, 23, 24, 25, 26, 27, 28, 29, 30)

	single := merge(a, b, c)

	for v := range single {
		fmt.Println("v", v)
	}

-->
	fmt.Println("The end")
}*/

/*--->loop:
for {
	select {
	case i, ok := <-single:
		if !ok {
			break loop
		}
		fmt.Println("i", i)
		//default:
	}
}*/

/*ch1 := make(chan int)
ch2 := make(chan string)

go func() { ch1 <- 42 }()
//go func() { ch2 <- "Привет, Go!" }()

cases := []reflect.SelectCase{
{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch1)}, // Чтение из ch1
{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch2)}, // Чтение из ch2
}

chosenIndex, value, ok := reflect.Select(cases)
if ok {
fmt.Printf("Канал %d выбран, значение: %v\n", chosenIndex, value.Interface())
} else {
fmt.Printf("Канал %d закрыт\n", chosenIndex)
}*/

//ch := make([]chan int, 1) - [nil]
