package manuals

import (
	"fmt"
	"reflect"
)

func countMe() {
}
func RunBuffer() {
	var a int = 123123
	b := int64(a)

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}
