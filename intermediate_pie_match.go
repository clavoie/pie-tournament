package main

import (
	"log"
	"sort"
	"strings"
)

type intermediatePieMatch struct {
	Id          string
	Year        int
	Bracket     int
	BracketType string
	Round       int
	PieA        string
	PieAVotes   int
	PieB        string
	PieBVotes   int
}

type intermediatePieMatches struct {
	Sortable []*intermediatePieMatch
}

var sortableIntermediatePieMatches = &intermediatePieMatches{
	Sortable: make([]*intermediatePieMatch, 0, 128),
}

func (ipm *intermediatePieMatches) AddIntermediateMatch(year, bracketNumber int, bracketType string, bracket *Bracket) {
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
		Id:          bracket.Poll.ID,
		Year:        year,
		Bracket:     bracketNumber,
		BracketType: bracketType,
		Round:       bracket.RoundNumber,
		PieA:        strings.TrimSpace(choiceA.Text),
		PieB:        strings.TrimSpace(choiceB.Text),
		PieAVotes:   choiceA.Votes,
		PieBVotes:   choiceB.Votes,
	})
}

func (ipm *intermediatePieMatches) ImportAllIntoPies() {
	ipm.Sort()

	for _, intermediatePieMatch := range ipm.Sortable {
		pieA := pies.AddIfMissing(intermediatePieMatch.PieA)
		pieB := pies.AddIfMissing(intermediatePieMatch.PieB)

		pieA.AddByeIfMissing(intermediatePieMatch)
		pieB.AddByeIfMissing(intermediatePieMatch)

		year := intermediatePieMatch.Year
		newMatchA := NewPieMatchFromIntermediate(pieA.NumNonByeMatches(year), intermediatePieMatch.PieAVotes, intermediatePieMatch.PieBVotes, intermediatePieMatch, pieA, pieB)
		newMatchB := NewPieMatchFromIntermediate(pieB.NumNonByeMatches(year), intermediatePieMatch.PieBVotes, intermediatePieMatch.PieAVotes, intermediatePieMatch, pieB, pieA)

		pieA.AddMatch(newMatchA)
		pieB.AddMatch(newMatchB)
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
