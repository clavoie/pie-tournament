package main

import "fmt"

var csvMatchHeader []string = []string{
	"MatchId", "Year", "TournamentRound", "PieRound", "Pie", "Votes", "Opponent", "OpponentVotes",
}

type CsvMatch struct {
	MatchId         string
	Year            string
	TournamentRound string
	PieRound        string
	Pie             string
	Votes           string
	Opponent        string
	OpponentVotes   string
}

func NewCsvMatch(pm *PieMatch) *CsvMatch {
	if pm.Result == Bye {
		return &CsvMatch{
			MatchId:         pm.MatchId,
			Year:            fmt.Sprint(pm.Year),
			TournamentRound: fmt.Sprint(pm.TournamentRound),
			PieRound:        fmt.Sprint(pm.MatchNumberForPie),
			Pie:             pm.Pie.Name,
			Votes:           fmt.Sprint(pm.VotesForPie),
		}
	}

	return &CsvMatch{
		MatchId:         pm.MatchId,
		Year:            fmt.Sprint(pm.Year),
		TournamentRound: fmt.Sprint(pm.TournamentRound),
		PieRound:        fmt.Sprint(pm.MatchNumberForPie),
		Pie:             pm.Pie.Name,
		Votes:           fmt.Sprint(pm.VotesForPie),
		Opponent:        pm.Opponent.Name,
		OpponentVotes:   fmt.Sprint(pm.VotesForOpponent),
	}
}

func (cm *CsvMatch) ToCsv() []string {
	return []string{
		cm.MatchId, cm.Year, cm.TournamentRound, cm.PieRound, cm.Pie, cm.Votes, cm.Opponent, cm.OpponentVotes,
	}
}
