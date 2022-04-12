package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)

	fmt.Println(notaFinal)

	// cuidado ao concatenar inteiros em strings
	fmt.Println("Teste " + string(123)) // 123 no caso é o código da table ascii para o valor {

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string pra int
	num, _ := strconv.Atoi("123") // usando _ significa que vamos ignorar o segundo retorno a função Atoi
	fmt.Println(num - 122)

	b, _ := strconv.ParseBool("1")
	if b {
		fmt.Println("Verdadeiro")
	}
}
