package main

type PieRank struct {
	Pie       *Pie
	Year1Rank int
	Year2Rank int
	Delta     int
	TotalRank int
}

type PieRankSorter struct {
	Year int
}

// func (prs *PieRankSorter) Len() int { return len(pies) }
// func (prs *PieRankSorter) Swap(i, j int) {
// 	pies[i], pies[j] = pies[j], pies[i]
// }
// func (prs *PieRankSorter) Less(i, j int) bool {

// }
