package manuals

import "fmt"

func divide(a, b int) {
	defer func() {
		fmt.Println("Error happened")
		if err := recover(); err != nil {
			fmt.Println("fixed->", err.(error).Error())
		}
	}() //defer will work even if error happens
	fmt.Println(a / b)        //error happens
	fmt.Println("Not called") //that's why this is not called
}
func RunDefer() {
	divide(1, 0)
	fmt.Println("Continue")
}

type X struct {
	V int
}

// x *X does matter
func (x X) S() {
	fmt.Println(x.V)
}
func RunDefer2() {
	x := X{1}   // or &X{1}, no matter
	defer x.S() //запомнит все свойства x
	defer func() { x.S() }()

	//return можем так манипулировать
	x.V = 2
}

func f2() (x int) {
	defer func() {
		x += 90
	}()
	x = 1
	return
}

func RunDefer3() {
	a := 1
	//defer fmt.Println(a)//запоминает аргументы, запомнит и выведет 1
	defer func() { //тут аргументов нет, во внутрь не лезет
		fmt.Println(a)
	}()
	a++
}
