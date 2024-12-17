package manuals

func runChannel(i int, ch chan<- int) {
	ch <- i
}
func RunBuffer() {
	//ch := make(chan int)

	/*go runChannel(1, ch)
	go runChannel(2, ch)
	go runChannel(3, ch)*/

	/*for i := range ch {
		fmt.Println(i)
	}*/

	/*for {
		select {
		case i := <-ch:
			fmt.Println("i->", i)
		default:
			break
		}
		break
	}*/

}
