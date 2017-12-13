package main

import (
	"fmt"
	"funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado da multiplicação foi %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Soma, 3, 3)
	fmt.Printf("O resultado da soma é %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 12, 3)
	fmt.Printf("O resultado da divisão é %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("O resultado da divisão é %d e o resto %d\r\n", resultado, resto)
}
