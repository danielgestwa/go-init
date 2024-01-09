package main

import "fmt"

func main() {
	var i int = 999
	var s string = "MyString"
	var b bool = true
	var u uint = 1
	const pepperPl = 200
	test, test2 := "Hello!", 1234
	var c64 complex64 = 10 + 2i
	var r rune = 'A'
	var bt byte = 'a'
	var f float64 = 0.000000001

	fmt.Println(i, s, b, u, test, test2, pepperPl, c64, bt, r, f)
}
