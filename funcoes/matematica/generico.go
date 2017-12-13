package matematica

/*
Calculo é um calculo genérico
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

/*
Multiplicador é a multplicação de um número pelo outro
*/
func Multiplicador(x int, y int) int {
	return x * y
}

/*
Divisor é a divisão de um número pelo outro
*/
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return resultado
}

/*
DivisorComResto retorna a divisão de um número pelo outro com o resto
*/
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}
