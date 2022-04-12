package main

import "fmt"

func main() {
	/*
	 entendendo quais são os valores
	 default de cada tipo
	*/
	var a int
	var b float64
	var c bool
	var d string

	//ponteiro de inteiro
	var e *int //nil é o valor null para golang

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
