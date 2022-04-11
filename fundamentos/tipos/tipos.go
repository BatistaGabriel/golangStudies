package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// só positivos = uint8(byte) uint16(short) uint32(int) uint64(long)
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// valores com sinal = int8(byte) int16(short) int32(int) int64(long)
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo da variável i1 é", reflect.TypeOf(i1))

	// representação do mapeamento da tabela unicode (int32)
	var i2 rune = 'a'
	fmt.Println("O valor de i2 é", i2)
	fmt.Println("O tipo da vairável i2 é", reflect.TypeOf(i2))

	// números reais = float32 e float64
	var x float32 = 49.99 // ou x = float32(49.99)
	fmt.Println("O tipo da variável x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal do valor 49.99 é", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Oi eu sou Goku"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com multiplas linhas
	s2 := `olá
	meu
	nome
	é
	Gibimba`
	fmt.Println(s2)
	fmt.Println("o tamanho da string é", len(s2))

	// golang tem char??
	// nope
	char := 'a'
	fmt.Println("O tipo da variável char é", reflect.TypeOf(char))
	fmt.Println(char)
}
