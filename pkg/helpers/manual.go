package helpers

import "fmt"

type Test struct {
	ID   *int    `json:"id"`
	Name *string `json:"name,omitempty"`
	Age  *string `json:"age"`
}

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

/*func divide(a, b int) {
	defer func() {
		fmt.Println("Error happened")
		if err := recover(); err != nil {
			fmt.Println("fixed->", err.(error).Error())
		}
	}() //defer will work even if error happens
	fmt.Println(a / b)        //error happens
	fmt.Println("Not called") //that's why this is not called
}
func RunPanic() {
	divide(1, 0)
	fmt.Println("Continue")
}*/

/*func Run() {
	a := []int{1, 2, 3}
	b := a[:2]
	b = append(b, 4)
	fmt.Println(a)//1,2,4
}*/

/*type Foo struct{}
func (*Foo) A() {}
func (*Foo) B() {}
func (*Foo) C() {}
type AB interface {
	A()
	B()
}
type BC interface {
	B()
	C()
}
type GG struct {
}
func Run() {
	var a AB = &Foo{}
	var c BC = a.(BC)//works
}*/

/*type X struct {
	V int
}
func Run() {
	x := X{1}
	//x := new(X)
	defer fmt.Println(x.V) //1 new or without pointer
	x.V = 2
}*/

type X struct {
	V int
}

// x *X does matter
func (x X) S() {
	fmt.Println(x.V)
}
func Run() {
	x := X{1}   // or &X{1}, no matter
	defer x.S() //выведет 1, потому что так запомнит переменную x.S
	x.V = 2
}

/*u := models.User{}
u.ID = 1
u.Name = "Alan"
u.Email = "Alan@mail.ru"*/

/*oe := reflect.ValueOf(&u).Elem()
//o := reflect.ValueOf(u) //поможет только читать, но не записывать
fv := oe.FieldByName("Email")

fv.Set(reflect.ValueOf("Changed@mail.ru").Convert(fv.Type()))
//fv.SetString("Changed@mail.ru")
fmt.Println("fv value", fv.Interface())*/

/*o := reflect.ValueOf(u)
for i := 0; i < o.NumField(); i++ {
	fmt.Println(o.Type().Field(i).Name, o.Field(i).Interface())
}*/

/*fs := http.FileServer(http.Dir("public/upload"))
http.Handle("/public", fs)*/
