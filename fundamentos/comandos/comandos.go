/*
o proposito desta aula é explorar
como podemos via linha de comando
realizar interações com o nosso
código em Go
*/
package main

import "fmt"

func main() {
	/*
	propositamentel vamos deixar a ancora %s
	mas não passaremos o valor, para termos no
	output o seguinte valor:

	Outro programa em %!s(MISSING)
	*/
	fmt.Printf("Outro programa em %s")

	/*
	para analizar melhor este cenário no nosso terminal
	vamos utilizar o comando:

	> go vet comandos.go

	este comando irá realizar analises e encontrar erros
	que podem não sei pegos pelo compilador

	após rodar este comando temos os seguinte output:
	
	# command-line-arguments
	./comandos.go:19:2: Printf format %s reads arg #1, but call has 0 args
	
	para apenas realizar a compilação do fonte -> go build nomeDoArquivo.go
	para realizar a compilação e execução do fonte -> go run nomeDoArquivo.go
	para realizar a instalação de alguma lib -> go get -u github.com/go-sql-driver/mysql

	*/
}
