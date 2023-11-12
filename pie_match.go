package main

//go:generate stringer -type PieMatchResult
type PieMatchResult int

const (
	Win PieMatchResult = iota
	Loss
	Tie
	TieLoss
	TieWin
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
	BracketNumber     int
	BracketType       string
}

func NewByePieMatch(year, bracketNumber int, bracketType string, b *Bracket, p *Pie) *PieMatch {
	return &PieMatch{
		Pie:               p,
		VotesForPie:       0,
		Result:            Bye,
		MatchNumberForPie: 0,
		TournamentRound:   b.RoundNumber - 1,
		Year:              year,
		BracketNumber:     bracketNumber,
		BracketType:       bracketType,
	}
}

func NewByePieMatchFromIntermediate(intermediatePieMatch *intermediatePieMatch, p *Pie) *PieMatch {
	return &PieMatch{
		Pie:               p,
		VotesForPie:       0,
		Result:            Bye,
		MatchNumberForPie: 0,
		TournamentRound:   intermediatePieMatch.Round - 1,
		Year:              intermediatePieMatch.Year,
		BracketNumber:     intermediatePieMatch.Bracket,
		BracketType:       intermediatePieMatch.BracketType,
	}
}

func NewPieMatch(year, bracketNumber, currentPieMatches int, bracketType string, b *Bracket, pieChoice, opponentChoice *PollChoice) *PieMatch {
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
		BracketNumber:     bracketNumber,
		BracketType:       bracketType,
	}
}

func NewPieMatchFromIntermediate(currentPieMatches, votes, opponentVotes int, interintermediatePieMatch *intermediatePieMatch, pie, opponent *Pie) *PieMatch {
	result := Tie

	if votes > opponentVotes {
		result = Win
	}
	if votes < opponentVotes {
		result = Loss
	}

	return &PieMatch{
		Pie:               pie,
		VotesForPie:       votes,
		Opponent:          opponent,
		VotesForOpponent:  opponentVotes,
		Result:            result,
		MatchId:           interintermediatePieMatch.Id,
		MatchNumberForPie: currentPieMatches + 1,
		TournamentRound:   interintermediatePieMatch.Round,
		Year:              interintermediatePieMatch.Year,
		BracketNumber:     interintermediatePieMatch.Bracket,
		BracketType:       interintermediatePieMatch.BracketType,
	}
}
