package main

import "fmt"

type produto struct {
	nome    string
	preço   float64
	estoque int
}

func swapp(ptr *produto, x float64) {
	ptr.preço = x
}

func main() {

	p := produto{nome: "Bicicleta", preço: 300, estoque: 20}

	fmt.Println("Informações antes da troca:")
	fmt.Println("O nome do produto é:", p.nome)
	fmt.Println("O preço do produto é:", p.preço)
	fmt.Println("A quantidade de produtos em estoque é:", p.estoque)

	swapp(&p, 450)
	fmt.Println("Informações depois da troca:")
	fmt.Println("O nome do produto é:", p.nome)
	fmt.Println("O preço do produto é:", p.preço)
	fmt.Println("A quantidade de produtos em estoque é:", p.estoque)

}
