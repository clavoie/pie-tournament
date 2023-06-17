package main

import (
	"encoding/csv"
	"log"
	"os"
)

type PieMatches struct {
	matches       []*PieMatch
	matchesByYear PieMatchesByYear
}

var pieMatches *PieMatches = &PieMatches{
	matches:       make([]*PieMatch, 0, 128),
	matchesByYear: NewPieMatchesByYear(),
}

func (pm *PieMatches) AddMatch(year int, bracket *Bracket) {
	poll := bracket.Poll
	if len(poll.Choices) != 2 {
		log.Fatal("Unhandled number of choices", poll.Choices)
	}

	choiceA := poll.Choices[0]
	if choiceA.Votes != choiceA.VoteCount {
		log.Fatal("choice votes do not match", choiceA)
	}
	choiceB := poll.Choices[1]
	if choiceB.Votes != choiceB.VoteCount {
		log.Fatal("choice votes do not match", choiceB)
	}

	pieA := pies.AddIfMissing(choiceA.Text)
	pieB := pies.AddIfMissing(choiceB.Text)
	matchesByPie := pm.matchesByYear.AddIfMissing(year)
	matchesByPie.AddByeIfMissing(year, bracket, pieA)
	matchesByPie.AddByeIfMissing(year, bracket, pieB)
	matchesByPie.Add(NewPieMatch(year, matchesByPie.NumMatches(pieA), bracket, choiceA, choiceB))
	matchesByPie.Add(NewPieMatch(year, matchesByPie.NumMatches(pieB), bracket, choiceB, choiceA))
}

func (pm *PieMatches) WriteAllMatchResults(fileName string) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer func() {
		err = csvFile.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	csvWriter := csv.NewWriter(csvFile)
	err = csvWriter.Write(csvMatchHeader)
	if err != nil {
		return err
	}

	for year := range pm.matchesByYear {
		yearMatches := pm.matchesByYear[year]
		for pie := range yearMatches {
			for _, match := range yearMatches[pie] {
				err = csvWriter.Write(NewCsvMatch(match).ToCsv())

				if err != nil {
					return nil
				}
			}
		}
	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if err != nil {
		return err
	}

	return nil
}
