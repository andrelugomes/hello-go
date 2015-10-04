package main


import "fmt"

//Estruturas são tipo Objetos
type coordenada struct{
	x,y,z int64
}

//Funcões apresentão retorno

//Uma func sem retorno é um Procedimento


//os tipos do retorno devem ficar emtre ()
func funcaoComDoisRetornos(vetor []int64) (int64, float64) {
	var soma int64

	/*
	Vai da problemas de não ura a variavel "i"
	Então tem que simbolizar o uso com "_"
	for i, valor := range vetor{
		soma += valor
	}*/


	//Tipo um foreach em um array. Onde Valor é o contudo do indece i para cada tipo no vetor
	for _, valor := range vetor{
		soma += valor
	}	

	//soma é int64, por isso tem que fazer um "cast" para float64 
	media := float64(soma)/float64(len(vetor)) //letodo len() ira contar quanto é o size do array
	return soma, media
}

//nesse caso. Informamos qual a variavel que retorna. Assim não precisa criala na função e nem mandar retornar... automaticamente o return vai faze-lo
func funcaoRetornaVariavel(vetor []int64) (soma int64, media float64) {
	
	for _, valor := range vetor{
		soma += valor
	}	
	media = float64(soma)/float64(len(vetor)) 
	return
}
func main() {
	//o vetor na verdade é um Slice, pois não tem range definido
	slice := []int64{1,2,3,4}

	//Adicionar ao array
	slice = append(slice,5)
	soma, media := funcaoRetornaVariavel(slice)
	fmt.Println(soma, media)


}