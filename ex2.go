package main

import "fmt"

func swapb(ptr *bool) {

	if *ptr == true {
		*ptr = false
	} else {
		*ptr = true
	}

}

func main() {

	var x bool
	x = true

	fmt.Println("Valores antes da troca:")
	fmt.Println("O conti é lindo!!", x)

	swapb(&x)
	fmt.Println("Valores depois da troca:")
	fmt.Println("O conti não é lindo!!", x)
}
