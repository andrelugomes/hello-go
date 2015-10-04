package main

import "fmt"

//fortemente tipada

func main() {
	//variavel estatica que é verificada em tempo de compilação
	var variavel int = 666 
	fmt.Println(variavel)

	//variavel atribuida em tempo de execução
	i := int64(5) //não existe conversão de tipos implecitamente - não há polimorfismo de coesão
	fmt.Println(i)
	i = 

	
}