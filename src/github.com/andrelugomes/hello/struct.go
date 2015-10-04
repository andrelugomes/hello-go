package main


import "fmt"

//Estruturas são tipo Objetos
type coordenada struct{
	x,y,z int64
}

//Para a função retornar... deve colocar no final o TIPO do Retorno, no caso "coordenada"
func criarCoordenada(c coordenada, x,y,z int64) coordenada {
	c.x = x
	c.y = y
	c.z = z	
	return c
}

func main() {
	//Criando a Estrutura
	var c coordenada
	c = criarCoordenada(c,1,2,3)
	fmt.Println(c)
}