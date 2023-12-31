$pieRanks = @{}
$pieToRound = @{}

Function Get-Result {
    Param(
        [Parameter(Mandatory = $true)]
        [int]$pie
        ,
        [Parameter(Mandatory = $true)]
        [int]$opponent
    )

    if ($pie -lt $opponent) {
        return "LOSS"
    }

    if ($pie -gt $opponent) {
        return "WIN"
    }

    return "TIE"
}


Get-ChildItem *.json | ForEach-Object {
    $foundYear = $_.FullName -match 'year-(\d+)'

    if (-not $foundYear) {
        throw "Could not find year for: $($_.FullName)"
    }

    $year = [int]$Matches[1]

    Get-Content $_.FullName | 
    convertfrom-json | 
    Select-Object -expandproperty reduxAsyncConnect | 
    Select-Object -ExpandProperty response | 
    Select-Object -ExpandProperty data | 
    Select-Object -ExpandProperty brackets | 
    Where-Object { -not $_.dummy } |
    Sort-Object -Property roundNumber |
    ForEach-Object { 
        [pscustomobject]@{ 
            Round   = $_.roundNumber
            Year    = $year
            Id      = $_.poll._id
            Results = ($_.poll.choices | ForEach-Object { [pscustomobject]@{ 
                        Pie   = $_.text
                        Votes = $_.votes 
                    } } ) 
        } }
} | ForEach-Object {
    $pieA = $_.Results[0]
    $pieB = $_.Results[1]

    if (-not $pieToRound[$pieA.Pie]) {
        $pieToRound[$pieA.Pie] = 0
    }

    $pieToRound[$pieA.Pie]++
    $pieARound = $pieToRound[$pieA.Pie]

    if (-not $pieToRound[$pieB.Pie]) {
        $pieToRound[$pieB.Pie] = 0
    }

    $pieToRound[$pieB.Pie]++
    $pieBRound = $pieToRound[$pieB.Pie]

    # $pieAPoints = $pieARound * $_.Year * $pieA.Votes
    # $pieBPoints = $pieBRound * $_.Year * $pieB.Votes
    $pieAPoints = $pieARound * $pieA.Votes
    $pieBPoints = $pieBRound * $pieB.Votes

    if ($pieARound -eq 1 -and $_.Round -gt 1) {
        [pscustomobject]@{
            MatchId         = ""
            TournamentRound = $_.Round - 1
            PieRound        = 0
            Year            = $_.Year
            Pie             = $pieA.Pie
            Votes           = 0
            Points          = 0
            Opponent        = ""
            OpponentVotes   = 0
            OpponentPoints  = ""
            Result          = "BYE"
        }
    }

    if ($pieBRound -eq 1 -and $_.Round -gt 1) {
        [pscustomobject]@{
            MatchId         = ""
            TournamentRound = $_.Round - 1
            PieRound        = 0
            Year            = $_.Year
            Pie             = $pieB.Pie
            Votes           = 0
            Points          = 0
            Opponent        = ""
            OpponentVotes   = 0
            OpponentPoints  = ""
            Result          = "BYE"
        }
    }

    [pscustomobject]@{
        MatchId         = $_.Id
        TournamentRound = $_.Round
        PieRound        = $pieARound
        Year            = $_.Year
        Pie             = $pieA.Pie
        Votes           = $pieA.Votes
        Points          = $pieAPoints
        Opponent        = $pieB.Pie
        OpponentVotes   = $pieB.Votes
        OpponentPoints  = $pieBPoints
        Result          = (Get-Result $pieA.Votes $pieB.Votes)
    }

    [pscustomobject]@{
        MatchId         = $_.Id
        TournamentRound = $_.Round
        PieRound        = $pieBRound
        Year            = $_.Year
        Pie             = $pieB.Pie
        Votes           = $pieB.Votes
        Points          = $pieBPoints
        Opponent        = $pieA.Pie
        OpponentVotes   = $pieA.Votes
        OpponentPoints  = $pieAPoints
        Result          = (Get-Result $pieB.Votes $pieA.Votes)
    }

    $pieRanks[$pieA.Pie] += $pieAPoints
    $pieRanks[$pieB.Pie] += $pieBPoints
} | Export-Csv -Path ".\pie-tourny-year-1.csv" -NoTypeInformation

$pieRanks.GetEnumerator() | Sort-Object  -Property Value -Descending | ForEach-Object {
    [pscustomobject]@{
        Name       = $_.Name
        Importance = $_.Value
    }
} | Export-Csv -Path ".\pie-tourny-year-1-rankings.csv" -NoTypeInformation