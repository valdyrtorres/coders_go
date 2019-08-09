package main

import "fmt"

func multiplicacao(a, b int) int {
	return a + b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

// Obs Não há implementação nativa de go para map, reduce e filter
// esses recursos aqui podem ajudar a implementar esses tipos de função

func main() {
	resultado := exec(multiplicacao, 3, 4)
	fmt.Println(resultado)
}
