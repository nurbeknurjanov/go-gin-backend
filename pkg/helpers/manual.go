package helpers

import (
	"fmt"
	"time"
)

type Test struct {
	ID   *int    `json:"id"`
	Name *string `json:"name,omitempty"`
	Age  *string `json:"age"`
}

func Say(ch chan int) {
	fmt.Println("3 Say start")
	ch <- 10
	fmt.Println("3.5 Say middle") // ???? why ??? должно было быть 4.5
	//Потому что скопом пихаются все значения в поток, разом функция выполняется
	ch <- 20 //Но 7 не выходит, потому что ch он как await
	fmt.Println("7 Say end")
}

func GoRun() {
	ch := make(chan int)
	fmt.Println("1 start main")
	go Say(ch)
	fmt.Println("2 middle main")
	//fun 3 say start
	fmt.Println("4 first", <-ch)
	fmt.Println("5 second", <-ch)
	//fun 6 say end
	//fmt.Println("no third", <-ch)
	fmt.Println("6 end main")
	/*for index, value range <-ch {

	}*/

	time.Sleep(2 * time.Second)
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
