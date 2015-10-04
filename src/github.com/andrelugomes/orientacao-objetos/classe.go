package main


import "fmt"

//Classe Coordenada
type coordenada struct{
	x,y,z int64
}

// Método referente a Classe usa *. Utiliza Ponteiros C++. Neste caso a "c" é um ponteiro pra coordenada
// Se retirar o ponteiro não funciona, pois passaria uma cópia e não alteraria o valor
func (c coordenada) setX(x int64) {
	c.x = x
}

func (c *coordenada) setY(y int64) {
	c.y = y
}

func (c *coordenada) setZ(z int64) {
	c.z = z
}

func criarCoordenada(c *coordenada, x,y,z int64) {
	c.x = x
	c.y = y
	c.z = z	
}

func main() {
	//Criando a Estrutura
	var c coordenada
	//& para passar o endereço da memória
	criarCoordenada(&c,1,1,1)
	c.setX(2)
	c.setY(2)
	c.setZ(2)
	fmt.Println(c)
}