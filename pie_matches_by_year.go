package main

type PieMatchesByYear map[int]PieMatchesByPie

func NewPieMatchesByYear() PieMatchesByYear {
	return make(PieMatchesByYear)
}

func (pmby PieMatchesByYear) AddIfMissing(year int) PieMatchesByPie {
	matches, exists := pmby[year]

	if exists {
		return matches
	}

	matches = NewPieMatchesByPie()
	pmby[year] = matches
	return matches
}
