package manuals

import "fmt"

func RunSlices() {
	a := []int{1, 2, 3}
	b := a[:2]
	b = append(b, 4)
	fmt.Println(a)//1,2,4
}