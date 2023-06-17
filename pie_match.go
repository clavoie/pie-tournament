package main

type PieMatchResult int

const (
	Win PieMatchResult = iota
	Loss
	Tie
	Bye
)

type PieMatch struct {
	Pie               *Pie
	VotesForPie       int
	Opponent          *Pie
	VotesForOpponent  int
	Result            PieMatchResult
	MatchId           string
	MatchNumberForPie int
	TournamentRound   int
	Year              int
}

func NewByePieMatch(year int, b *Bracket, p *Pie) *PieMatch {
	return &PieMatch{
		Pie:               p,
		VotesForPie:       0,
		Result:            Bye,
		MatchNumberForPie: 0,
		TournamentRound:   b.RoundNumber - 1,
		Year:              year,
	}
}

func NewPieMatch(year, currentPieMatches int, b *Bracket, pieChoice, opponentChoice *PollChoice) *PieMatch {
	result := Tie

	if pieChoice.Votes > opponentChoice.Votes {
		result = Win
	}
	if pieChoice.Votes < opponentChoice.Votes {
		result = Loss
	}
	return &PieMatch{
		Pie:               pies.AddIfMissing(pieChoice.Text),
		VotesForPie:       pieChoice.Votes,
		Opponent:          pies.AddIfMissing(opponentChoice.Text),
		VotesForOpponent:  opponentChoice.Votes,
		Result:            result,
		MatchId:           b.Poll.ID,
		MatchNumberForPie: currentPieMatches + 1,
		TournamentRound:   b.RoundNumber,
		Year:              year,
	}
}
