package main

import "fmt"
import "runtime"
import "time"

func main() {

	numero := 3

	fmt.Print("O número ", numero, " se escreve assim: ")

	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("Você é da família do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("Sim!")
	default:
		fmt.Println("Deixa pra lá...")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("14h da marugada eu acordo...")
	default:
		fmt.Println("Preciso de café...")
	}

	numero = 10
	fmt.Println("Este numero cabe num dígito?")
	switch {
	case numero >= 10 && numero < 100:
		fmt.Println("Serve dois dígitos?")
	case numero >= 100:
		fmt.Println("Não dá...")
	}
}
