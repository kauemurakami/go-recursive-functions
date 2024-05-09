package main

import "fmt"

// 1 2 3 4 5 6  7 posições
// 1 1 2 3 5 8 13 .... // valores
func fibonacci(position uint) uint {
	if position <= 1 { // caso chegue no valor 1 ou zero paramos o loop
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	position := uint(7)
	// position := uint(1000) erro pois está empilhando muitas execuções,
	//												uma dependendo da outra o que deixa o precess muito pesado

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println(fibonacci(position)) //output 13 ou eja a posição 7 tem o valor 13 seguindo a sequência

}
