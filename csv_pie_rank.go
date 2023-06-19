package main

var csvPieRankHeader []string = []string{
	"Pie", "Year1Rank", "Year2Rank", "Delta", "TotalRank",
}

type CsvPieRank struct {
	Pie       string
	Year1Rank string
	Year2Rank string
	Delta     string
	TotalRank string
}
