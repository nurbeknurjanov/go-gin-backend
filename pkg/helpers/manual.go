package helpers

import (
	"fmt"
)

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

func Say1(ch chan int) {
	ch <- 10
	fmt.Println("1 input")
	ch <- 20
	fmt.Println("2 input")
}

func GoRun() {
	ch := make(chan int)
	go Say1(ch)
	//fun 3 say start
	//fun 3.5 say middle despite ch
	fmt.Println("1 output", <-ch)
	fmt.Println("2 second", <-ch)
}

/*func say(ch chan int) {
	ch <- 1
	fmt.Println("After 1")
	ch <- 2
	fmt.Println("After 2")
	ch <- 3
	fmt.Println("After 3")
	ch <- 4
	fmt.Println("After 4")
	ch <- 5
	fmt.Println("After 5")
	ch <- 6
	fmt.Println("After 6")
	ch <- 7
	fmt.Println("After 7")
	ch <- 8
	fmt.Println("After 8")
	ch <- 9
	fmt.Println("After 9")
	ch <- 10
	fmt.Println("After 10")
}

func GoRun() {
	ch := make(chan int)
	go say(ch)
	fmt.Println("el1", <-ch)
	fmt.Println("el2", <-ch)
	fmt.Println("el3", <-ch)
	fmt.Println("el4", <-ch)
	fmt.Println("el5", <-ch)
	fmt.Println("el6", <-ch)
	fmt.Println("el7", <-ch)
	fmt.Println("el8", <-ch)
	fmt.Println("el9", <-ch)
	fmt.Println("el10", <-ch)

	time.Sleep(2 * time.Second)
}*/

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
