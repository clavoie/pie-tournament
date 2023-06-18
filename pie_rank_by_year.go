package main

type PieRanks map[int][]*PieRank

func NewPieRanks() PieRanks {
	return make(PieRanks)
}

var pieRanks = NewPieRanks()

func (prs PieRanks) Calculate() {
	for year := range pieMatches.matchesByYear {
		pieRanks[year] = make([]*PieRank, 0, 128)
	}
}
