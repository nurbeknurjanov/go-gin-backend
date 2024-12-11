package manuals

type Foo struct{}

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

func RunInterfaces() {
	var a AB = &Foo{}
	var c BC = a.(BC)
	c.C()
}

/*var a *int
fmt.Println(I(a) == a, a, I(a)) */
