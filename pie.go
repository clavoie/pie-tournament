package main

type Pie struct {
	Name          string
	PointsByYear  map[int]int
	MatchesByYear map[int][]*PieMatch
}

func NewPie(name string) *Pie {
	return &Pie{
		Name:          name,
		PointsByYear:  map[int]int{},
		MatchesByYear: map[int][]*PieMatch{},
	}
}

func (p *Pie) AddByeIfMissing(intermediatePieMatch *intermediatePieMatch) {
	year := intermediatePieMatch.Year
	if intermediatePieMatch.Round > 1 && p.NumMatches(year) == 0 {
		p.MatchesByYear[year] = append(p.MatchesByYear[year], NewByePieMatchFromIntermediate(intermediatePieMatch, p))
	}
}

func (p *Pie) AddMatch(pieMatch *PieMatch) {
	year := pieMatch.Year
	p.MatchesByYear[year] = append(p.MatchesByYear[year], pieMatch)
	p.PointsByYear[year] += pieMatch.VotesForPie * pieMatch.MatchNumberForPie
}

func (p *Pie) NumMatches(year int) int {
	return len(p.MatchesByYear[year])
}
