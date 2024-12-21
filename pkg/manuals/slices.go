package manuals

func RunSlices() {
	//добавить в начало
	//a = append([]int{69}, a...)

	/*a = append(a, 0)
	copy(a[1:], a)
	a[0] = 69*/

	//a = append(a[:4], a[5:]...) //1 2 3 4 6 7 8 9 10] удалить 5 элемент

	// Вычисляем индекс середины
	//mid := len(slice) / 2
	// Вставка элемента в середину
	//slice = append(slice[:mid], append([]int{element}, slice[mid:]...)...)

	/*
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		arr2 := arr[4:5]
		arr2 = append(arr2, 60) //arr2 кобойот
		arr - просто меняется состав 6 на 60

		ВАЖНО, зависит от длины arr, не cap
		copy(arr, []int{10, 20, 30}) - меняется состав
	*/

	//arr = arr[1:] // длина норм сокращается
	//arr[:0] - норм становитя пустым
	//arr = append(arr[:0:0], arr[1:]...)
	//arr = append([]int{}, arr[1:]...)
	//arr = append(arr[:0], arr[1:]...)

	//убрать последний - плохой вариант, но работает
	/*copy(arr, arr[1:])
	arr = arr[:9]*/
	//arr = arr[:len(arr)-1]

	// Удаление элемента
	//arr = append(arr[:4], arr[4+1:]...)
	//arr = append(arr[:4], arr[4+2:]...)

	//вставить элемент в начало среза
	//https://chatgpt.com/c/675b91cc-b854-8012-a5e3-57a126037a50
	//arr = append([]int{69}, arr...)

	//arr = append(arr, 0) //1 2 3 4 5 6 7 8 9 10 0
	//сдвиг назад
	//copy(arr[1:], arr) //1 1 2 3 4 5 6 7 8 9 10
	//arr[0] = 69
	/*Если требуется простота, используйте способ 1.
	Для большего контроля и эффективности — способ 2.*/
	//сдвиг через [:] лучше для скорости и памяти
}

/*a := []int{1, 2, 3, 4, 5}
b := a[2:3]
fmt.Println(b)

b[0] = 30
b = append(b, 40)
fmt.Println(a) //[1 2 30 40 5]*/

/*
var slice = make([]byte, 5)
//copy(slice, []byte{65, 66})
copy(slice, "AB")
fmt.Println(string(slice))*/

/*var arr = []int{1, 2, 3}
slice := arr[0:0] - slice := arr[0:0:0] тогда не будет 74 - при этом range будет адекватно работать
fmt.Println(slice)      // пусто
fmt.Println(slice[0:])  // пусто
fmt.Println(slice[0:3]) // а тут выводит!!!!!!!!!!!!!!!!!!!!!!!!!!*/
