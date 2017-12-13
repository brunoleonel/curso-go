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
