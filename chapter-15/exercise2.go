package main

import "fmt"
type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func changePersonName(p *pessoa, newNome string){
	(*p).nome = newNome
}
func main(){
	pessoa := pessoa{
		nome: "Fulano",
		sobrenome: "de Tal",
		idade: 26,
	}
	fmt.Println(pessoa)
	changePersonName(&pessoa, "Beltrano")
	fmt.Println(pessoa)
}