package main

type pessoa struct {
	nome string
	sobrenome string
	idade int
}

func (p pessoa) nomeCompleto() {
	println("Meu nome é " + p.nome + " " + p.sobrenome)
}

func main(){
	pessoa := pessoa{
		nome: "Fulano",
		sobrenome: "De Tal",
		idade: 26,
	}
	pessoa.nomeCompleto()
}