package manuals

import (
	"fmt"
	"math"
)

func RunFloats() {

	var a float64 = 10
	var b float64 = 3
	c := a / b
	//fmt.Println("c", c) //3.3333333333333335
	//fmt.Printf("c is  %d\n", c) //%!d(float64=3.3333333333333335)

	//d := int8(c)
	//fmt.Println("d", d) //3
	fmt.Println("4", math.Ceil(c))
	fmt.Println("3", math.Floor(c))

	fmt.Println("5", math.Round(4.5))
	fmt.Println("4", math.Round(4.4))

	fmt.Println("4.43", math.Round(4.43123*100)/100)
	fmt.Println("4.46", math.Round(4.45523*100)/100)

	fmt.Printf("%.2f\n", 12.3456)
}
