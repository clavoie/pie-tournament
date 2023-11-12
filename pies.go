package main

import "strings"

type Pies map[string]*Pie

var pies Pies = make(Pies)

func (ps Pies) AddIfMissing(pieName string) *Pie {
	pieName = strings.TrimSpace(pieName)
	pie, exists := ps[pieName]

	if exists {
		return pie
	}

	pie = NewPie(pieName)
	ps[pieName] = pie
	return pie
}

func (ps Pies) ConvertTiesToTiesWL() {
	for _, pie := range ps {
		pie.ConvertTiesToTiesWL()
	}
}
