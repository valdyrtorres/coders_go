package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	// Essa construcao, somente funciona com slice. Nao funfa com array
	//aprovados := [...]string{"Maria", "Pedro", "Rafael", "Guilherme"} <- Nao funfa
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...)
}
