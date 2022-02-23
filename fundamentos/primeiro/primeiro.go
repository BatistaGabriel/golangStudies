/*
* todos programas em GO roda dentro de um pacote
* para que sejam executados eles ser rodados no pacote
* principal que se chama main

* pacotes -> modulos de códigos que pode ser utilizados
 */
package main

/*
*	a forma como indicamos quais bibliotecas vamos usar
*	dentro de um pacote, similar ao using do C#
*
*	para multiplos imports é possível usar a sintaxe:
*	import (
*		"package_A",
*		"package_B",
*			...
*	)
 */
import "fmt" // -> biblioteca que consegue fazer os output para o usuário

/*
*	a função nomeada como mai vai ser a primeira
*	que a linguagem irá buscar para executar
 */
func main() {
	/*
	*	o método Println ira "cuspir" valores
	*	na tela para o usuário, quebrando linhas
	 */
	fmt.Println("Primeiro")
	fmt.Println("Programa!")
}
