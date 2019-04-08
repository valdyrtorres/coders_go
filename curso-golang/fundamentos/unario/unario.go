package main

import "fmt"

func main() {
	x := 1
	y := 1

	// Apenas forma pos-fixada(postfix)
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou y = y - 1
	fmt.Println(y)

	// Não é possível em GO
	//fmt.Println(x == y--)
}
