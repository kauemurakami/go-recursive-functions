[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-recursive-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-recursive-functions/blob/main/README.md)

## Funções recursivas
Funções recursivas basicamente são funções que chamam elas mesmas, nesse caso é como se ela dependesse de outra execução dela mesma.  

### Usando
Vamos fazer uma função que vai retornar o número que está em uma determinada posição da sequemcia de Fibonacci.  [O que é a sequência de Fibonacci](https://en.wikipedia.org/wiki/Fibonacci_sequence)  

### Criando função
O importante em funções recursivas, é que elas tenham uma condição de parada, ou seja, você deve especificar um momento que ela deve parar de chamar ela mesma, para que não haja um estouro de pilha quando ela roda infinitamente.  
```go
package main

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
	// position := uint(1000) // erro pois está empilhando muitas execuções, muita demora e
	//												uma dependendo da outra o que deixa o precess muito pesado

	fmt.Println(fibonacci(position)) //output 55 ou eja a posição 10 tem o valor 55 seguindo a sequência
}
```
Algo a mais que podemos fazer é imprimir cada número da sequência de fibonacci que percorremos.  
```go
...
func main() {
	position := uint(7) 
	// position := uint(1000) erro pois está empilhando muitas execuções,
	//												uma dependendo da outra o que deixa o precess muito pesado

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
    // output
    // 1
    // 1
    // 2
    // 3
    // 5
    // 8 // 8 + 5 é 13 por isso paramos aqui. 
	}

	fmt.Println(fibonacci(position)) //output 13 ou eja a posição 10 tem o valor 55 seguindo a sequência
  //out
}
```

