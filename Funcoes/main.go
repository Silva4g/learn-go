package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt ...string) string {
		fmt.Println("Funcao f:", txt)
	}
	f("texto da funcao f")
	//f()
}
