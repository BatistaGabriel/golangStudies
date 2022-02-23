package main

import (
	"fmt"
	"math"
)

/*
	IMPORTANTE: em GO variáveis tem que ser declaradas e
	utilizadas, caso seja declarada uma variável e não utilizada
	o próprio compilador do GO irá acusar um erro
*/

/*
	neste arquivo vamos explorar como podemos
	em GO trabalhar com constantes e variáveis
*/
func main() {
	const PI float64 = 3.1415 // esta é a declaração mais comum

	/*
		para variáveis o tipo não é obrigatório
		o GO consegue inferir o valor dela assim
		que é realizada a atribuíção do valor para ela
	*/
	var raio = 3.2

	/*
		no entendo esta é a forma mais comum de
		declaração de uma variável, a forma reduzida
		note que para este caso temos a utilização do :=

		:= -> isto indica para o GO que variável ainda não existe
		e que ela deverá ser criada e inicializada com o valor que
		virá logo a frente

		para variáveis existentes apenas do uso do = serve
		para a atribuição do valor
	*/
	area := PI * math.Pow(raio, 2)

	fmt.Println("A área da circunferência é:", area)

	/*
		outras formas de definição de constantes e variáveis
		estas formas são blocos de definição de valores
	*/
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	/*
		declaração de variável in-line
	*/
	var e, f bool = true, false
	fmt.Println(e, f)

	/*
		declaração de variável in-line com a forma reduzida
	*/
	g, h, i := 2, false, "olá!"
	fmt.Println(g, h, i)

	/*
		BONUS:

		no imports é possivel setar alias para a lib caso deseje
		para fazer isto basta apenas adicionar o nome do alias
		antes do nome da lib:

		import(
			"fmt"
			m "math"
		)

		neste caso, quando quiser usar algo da lib math a sintaxe
		ficará assim:

		m.Pow(<arg1>, <arg2>)

		--

		caso queira ter uma lib no arquivo, porém, ele não será usada
		e você queira que o compilador não remova sua referencia adicione
		um _ antes do nome da lib:

		import(
			"fmt"
			_ "math"
		)
	*/
}
