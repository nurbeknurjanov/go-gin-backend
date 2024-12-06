package helpers

import "fmt"

type Test struct {
	ID   *int    `json:"id"`
	Name *string `json:"name,omitempty"`
	Age  *string `json:"age"`
}

func TestFunc() {
	fmt.Println("TestFunc")
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
