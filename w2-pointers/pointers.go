package main

func main() {
	// pointer var
	var pointer *int
	pointer = new(int)
	*pointer = 4
	println(pointer)
	println(*pointer)

	// modify by pointer
	a := 4
	b := &a
	*b = 5
	c := *b
	println(a)
	println(c)
}
