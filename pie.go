package main

type Pie struct {
	Name string
}

func NewPie(name string) *Pie {
	return &Pie{name}
}
