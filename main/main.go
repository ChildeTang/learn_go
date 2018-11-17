package main

import "fmt"

func main() {
	fmt.Println("Hello Word!!!")
	fmt.Println("Hello Word!!!")
	fmt.Println("Hello Word!!!")

	var i int
	i = 10
	fmt.Println("i =", i)

	str := `123
123
  123
	123`
	fmt.Println(str)

	var a int
	var b float32
	var c float64
	var d bool
	var e string
	fmt.Printf("%d %.5f %.8f %v %v\n",a, b, c, d, e)

	fmt.Printf("%T",&i)

	m := make(map[string] map[string] map[int] byte, 10)
	p := make(map[string] int, 10)
	fmt.Printf("%v\n", m)
	fmt.Printf("%v\n", p)
}