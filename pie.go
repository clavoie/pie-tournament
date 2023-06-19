package main

import "fmt"

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

func (p *Pie) ConvertTiesToTiesWL() {
	exceptions := map[int][]string{
		2: {"Strawberry Blackberry", "Atlantic Beach", "Black Bottom Mocha", "Bourbon Chocolate Pecan"},
	}
	for year, pieMatches := range p.MatchesByYear {
		for index, pieMatch := range pieMatches {
			if pieMatch.Result != Tie {
				continue
			}

			if (index + 1) >= len(pieMatches) {
				exs, hasException := exceptions[year]
				exFound := false

				if hasException {
					for _, ex := range exs {
						if p.Name == ex {
							exFound = true
							break
						}
					}
				}

				if exFound {
					pieMatch.Result = TieWin
				} else {
					pieMatch.Result = TieLoss
				}
			} else {
				pieMatch.Result = TieWin
			}
		}
	}
}

func (p *Pie) NumMatches(year int) int {
	return len(p.MatchesByYear[year])
}

func (p *Pie) NumNonByeMatches(year int) int {
	count := 0

	for _, match := range p.MatchesByYear[year] {
		if match.Result == Bye {
			continue
		}

		count++
	}

	return count
}

func (p *Pie) CsvPieRankRecord() []string {
	return []string{
		p.Name,
		fmt.Sprint(p.PointsByYear[1]),
		fmt.Sprint(p.PointsByYear[2]),
		fmt.Sprint(p.PointsByYear[2] - p.PointsByYear[1]),
		fmt.Sprint(p.PointsByYear[2] + p.PointsByYear[1]),
	}
}
