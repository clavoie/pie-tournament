package main

// should it know the year>>
type PieMatchesByPie map[*Pie][]*PieMatch

func NewPieMatchesByPie() PieMatchesByPie {
	return make(PieMatchesByPie)
}

func (pmbp PieMatchesByPie) Add(pm *PieMatch) {
	pmbp[pm.Pie] = append(pmbp[pm.Pie], pm)
}

func (pmbp PieMatchesByPie) AddByeIfMissing(year, bracketNumber int, bracketType string, bracket *Bracket, p *Pie) {
	if bracket.RoundNumber > 1 && pmbp.NumMatches(p) == 0 {
		pmbp[p] = append(pmbp[p], NewByePieMatch(year, bracketNumber, bracketType, bracket, p))
	}
}

func (pmbp PieMatchesByPie) AddByeIfMissingFromIntermediate(intermediatePieMatch *intermediatePieMatch, p *Pie) {
	if intermediatePieMatch.Round > 1 && pmbp.NumMatches(p) == 0 {
		pmbp[p] = append(pmbp[p], NewByePieMatchFromIntermediate(intermediatePieMatch, p))
	}
}

func (pmbp PieMatchesByPie) NumMatches(p *Pie) int {
	count := 0
	for _, pm := range pmbp[p] {
		if pm.Result == Bye {
			continue
		}

		count++
	}

	return count
}
