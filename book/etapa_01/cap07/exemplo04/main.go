package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Olá, meu nome é", p.Name)
}

type Robot struct {
	Serie int
}

type Androide struct {
	Person
	Model string
}

func main() {
	meuAndroid := Androide{}
	meuAndroid.Talk()
}
