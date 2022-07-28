package main

import (
    "fmt"
    "math"
)

func main(){
    a := 3
    b := 2

    fmt.Println("Soma =>", a+b)
    fmt.Println("Subtração =>", a-b)
    fmt.Println("Divisão =>", a/b)
    fmt.Println("Multiplicação =>", a*b)
    fmt.Println("Módulo =>", a%b)

    //Bitwise

    /*
    * Valor binario de a = 11
    * Valor binario de b = 10
    */

    fmt.Println("E =>", a&b) // Resultado 10
    fmt.Println("Ou =>", a|b) // Resultado 11
    fmt.Println("Xor =>", a^b) // Resultado 01

    c := 3.0
    d := 2.0

    /*
    * Operações usando a classe Math
    */
    fmt.Println("Maior =>", math.Max(float64(a), float64(b)))
    fmt.Println("Menor =>", math.Min(c,d))
    fmt.Println("Exponenciação =>", math.Pow(c, d))
}
