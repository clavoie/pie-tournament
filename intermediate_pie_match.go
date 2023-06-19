package main

import (
	"log"
	"sort"
)

type intermediatePieMatch struct {
	Id        string
	Year      int
	Bracket   int
	Round     int
	PieA      string
	PieAVotes int
	PieB      string
	PieBVotes int
}

type intermediatePieMatches struct {
	Sortable []*intermediatePieMatch
}

var sortableIntermediatePieMatches = &intermediatePieMatches{
	Sortable: make([]*intermediatePieMatch, 0, 128),
}

func (ipm *intermediatePieMatches) AddIntermediateMatch(year, bracketNumber int, bracket *Bracket) {
	if bracket.Dummy {
		return
	}

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

	ipm.Sortable = append(ipm.Sortable, &intermediatePieMatch{
		Id:        bracket.Poll.ID,
		Year:      year,
		Bracket:   bracketNumber,
		Round:     bracket.RoundNumber,
		PieA:      choiceA.Text,
		PieB:      choiceB.Text,
		PieAVotes: choiceA.Votes,
		PieBVotes: choiceB.Votes,
	})
}

func (ipm *intermediatePieMatches) ImportAllIntoPieMatches() {
	for _, intermediateMatch := range ipm.Sortable {
		pieMatches.ImportFromIntermediate(intermediateMatch)
	}
}

func (ipm *intermediatePieMatches) Sort() { sort.Sort(ipm) }

func (ipm *intermediatePieMatches) Len() int { return len(ipm.Sortable) }
func (ipm *intermediatePieMatches) Swap(i, j int) {
	ipm.Sortable[i], ipm.Sortable[j] = ipm.Sortable[j], ipm.Sortable[i]
}
func (ipm *intermediatePieMatches) Less(i, j int) bool {
	a, b := ipm.Sortable[i], ipm.Sortable[j]

	if a.Year < b.Year {
		return true
	}

	if a.Year > b.Year {
		return false
	}

	// same year

	if a.Round < b.Round {
		return true
	}

	if a.Round > b.Round {
		return false
	}

	// same round

	// if same year and same round, then one is not less than the other
	return false
}
