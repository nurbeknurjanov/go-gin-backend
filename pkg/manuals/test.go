package manuals

import "fmt"

type Test struct {
	ID   *int    `json:"id"`
	Name *string `json:"name,omitempty"`
	Age  *string `json:"age"`
}

/*func RunTest() {
	a := 2
	switch {
	default:
		fmt.Println("Default")
	case a == 1:
		fmt.Println(1)
		fmt.Println(1)
	case a == 2:
		fmt.Println(2)
		fmt.Println(2)
	}
	fmt.Println("End")
}*/

/*func RunTest() {
	switch a := 2; a {
	default:
		fmt.Println("Default")
	case 1:
		fmt.Println(1)
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fmt.Println(2)
	}
	fmt.Println("End")
}*/

/*func CallMe(x string, y ...any) {
	fmt.Println(x)
	fmt.Println(y...) //only any
}
func RunTest() {
	y := []any{1, 2, 3}
	CallMe("a", y...) //только слайсы
}*/

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}
func goTwo(ch chan string) {
	ch <- "From goTwo goroutine"
}
func RunTest() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case msg11 := <-ch1:
		fmt.Println("msg11", msg11)
	case msg22 := <-ch2: //stronger
		fmt.Println("msg22", msg22)
	}
}
