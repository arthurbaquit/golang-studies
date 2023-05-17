package main

type pessoa struct{}

type Humanos interface {
	Falar()
}

func (p *pessoa) Falar() {
	println("Olá, eu sou uma pessoa!")
}

func dizerAlgumaCoisa(humano Humanos) {
	humano.Falar()
}

func main() {
	p := pessoa{}
	// dizerAlgumaCoisa(p) error
	(&p).Falar()
	p.Falar() // <- This is a shortcut for the line above
	dizerAlgumaCoisa(&p)
}
