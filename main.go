package main

import "log"

func main() {
	// fmt.Println(os.Args)

	processAllBracketsInCurrentDir()

	err := pieMatches.WriteAllMatchResults("results-from-go.csv")
	if err != nil {
		log.Fatal(err)
	}

	err = pieMatches.WriteAllRanks("results-from-go-ranks.csv")
	if err != nil {
		log.Fatal(err)
	}
}
