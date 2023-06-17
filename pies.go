package main

type Pies map[string]*Pie

var pies Pies = make(Pies)

func (ps Pies) AddIfMissing(pieName string) *Pie {
	pie, exists := ps[pieName]

	if exists {
		return pie
	}

	pie = NewPie(pieName)
	ps[pieName] = pie
	return pie
}
