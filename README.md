[![pt-br](https://img.shields.io/badge/language-pt--br-green.svg)](https://github.com/kauemurakami/go-recursive-functions/blob/main/README.pt-br.md)
[![en](https://img.shields.io/badge/language-en-orange.svg)](https://github.com/kauemurakami/go-recursive-functions/blob/main/README.md)

## Recursive functions
Recursive functions are basically functions that call themselves, in this case it is as if it depended on another execution of itself.  

### Using
Let's make a function that will return the number that is in a certain position in the Fibonacci sequence. [What is the Fibonacci sequence](https://en.wikipedia.org/wiki/Fibonacci_sequence)  

### Creating function
The important thing about recursive functions is that they have a stopping condition, that is, you must specify a moment at which it must stop calling itself, so that there is no stack overflow when it runs infinitely.  
```go
package main

// 1 2 3 4 5 6  7 positions
// 1 1 2 3 5 8 13 .... // values
func fibonacci(position uint) uint {
	if position <= 1 { // If it reaches the value 1 or zero, we stop the loop
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	position := uint(7)
	// position := uint(1000) // error because it is piling up many executions, a lot of delay and
	//												one depending on the other, which makes the precess very heavy

	fmt.Println(fibonacci(position)) //output 13 ou eja a posição 7 tem o valor 13 seguindo a sequência
}
```
Something else we can do is print each number in the Fibonacci sequence that we go through.  
```go
...
func main() {
	position := uint(7) 
	// position := uint(1000) error because it is stacking many executions,
	//												one depending on the other, which makes the precess very heavy

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonacci(i))
    // output
    // 1
    // 1
    // 2
    // 3
    // 5
    // 8 // 8 + 5 is 13 so we stop here.
	}

	fmt.Println(fibonacci(position)) //output 13 i.e. position 7 has the value 13 following the sequence
}
```

