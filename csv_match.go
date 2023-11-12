package main

import (
	"fmt"
	"log"
)

var csvMatchHeader []string = []string{
	"MatchId", "Year", "TournamentRound", "PieRound", "Bracket", "BracketType", "Pie", "Votes", "Opponent", "OpponentVotes", "Result",
}

type CsvMatch struct {
	MatchId         string
	Year            string
	TournamentRound string
	PieRound        string
	Bracket         string
	BracketType     string
	Pie             string
	Votes           string
	Opponent        string
	OpponentVotes   string
	Result          string
}

func NewCsvMatch(pm *PieMatch) *CsvMatch {
	if pm.BracketType == "" {
		log.Fatalf("Missing pie bracket for match %v", pm.MatchId)
	}

	if pm.Result == Bye {
		return &CsvMatch{
			MatchId:         pm.MatchId,
			Year:            fmt.Sprint(pm.Year),
			TournamentRound: fmt.Sprint(pm.TournamentRound),
			PieRound:        fmt.Sprint(pm.MatchNumberForPie),
			Bracket:         fmt.Sprint(pm.BracketNumber),
			BracketType:     pm.BracketType,
			Pie:             pm.Pie.Name,
			Votes:           fmt.Sprint(pm.VotesForPie),
			Result:          pm.Result.String(),
		}
	}

	return &CsvMatch{
		MatchId:         pm.MatchId,
		Year:            fmt.Sprint(pm.Year),
		TournamentRound: fmt.Sprint(pm.TournamentRound),
		PieRound:        fmt.Sprint(pm.MatchNumberForPie),
		Bracket:         fmt.Sprint(pm.BracketNumber),
		BracketType:     pm.BracketType,
		Pie:             pm.Pie.Name,
		Votes:           fmt.Sprint(pm.VotesForPie),
		Opponent:        pm.Opponent.Name,
		OpponentVotes:   fmt.Sprint(pm.VotesForOpponent),
		Result:          pm.Result.String(),
	}
}

func (cm *CsvMatch) ToCsv() []string {
	return []string{
		cm.MatchId, cm.Year, cm.TournamentRound, cm.PieRound, cm.Bracket, cm.BracketType, cm.Pie, cm.Votes, cm.Opponent, cm.OpponentVotes, cm.Result,
	}
}
