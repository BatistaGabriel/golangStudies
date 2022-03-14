package main

import "fmt"

func main() {
	fmt.Print("O comando print")
	fmt.Print(" vai pritnar tudo na mesma linha.")

	fmt.Println(" O comando println começa na linha anterior")
	fmt.Println("e no final quebra para uma nova linha.")

	x := 3.141516
	// fmt.Println("O valor de x é" + x)
	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "hello"

	fmt.Printf("\n%d %f %1.f %t %s", a, b, b, c, d)
}
