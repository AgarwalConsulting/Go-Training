func A() {
	a := 10
	B()
	...
}

func B() {
	b := 20
	C()
	...
}

func C() {
	c := 30
	fmt.Println(a, b, c) // ?
}

func main() {
	A()
}
