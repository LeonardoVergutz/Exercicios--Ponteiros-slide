package main

import "fmt"

func swap(x, y *int) {

	*x, *y = *y, *x

}

func main() {
	x := 10
	y := 5
	fmt.Println("Valores antes da troca:")
	fmt.Println("O valor de X é:", x)
	fmt.Println("O valor de Y é:", y)
	swap(&x, &y)

	fmt.Println("Valores depois da troca:")
	fmt.Println("O valor de X é:", x)
	fmt.Println("O valor de Y é:", y)
}
