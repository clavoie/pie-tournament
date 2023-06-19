package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"
)

var filePattern = regexp.MustCompile(`^year-(\d)-bracket-(\d).json$`)

func processAllBracketsInCurrentDir() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		nameMatches := filePattern.FindStringSubmatch(fileName)

		if nameMatches == nil {
			continue
		}

		year, err := strconv.Atoi(nameMatches[1])
		if err != nil {
			log.Fatal(fileName, err)
		}

		bracketNumber, err := strconv.Atoi(nameMatches[2])
		if err != nil {
			log.Fatal(fileName, err)
		}

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			err = file.Close()

			if err != nil {
				log.Fatal(err)
			}
		}()

		decoder := json.NewDecoder(file)
		var autoGenerated AutoGenerated
		err = decoder.Decode(&autoGenerated)

		if err != nil {
			log.Fatal(err)
		}

		brackets := autoGenerated.ReduxAsyncConnect.Response.Data.Brackets

		for _, bracket := range brackets {
			sortableIntermediatePieMatches.AddIntermediateMatch(year, bracketNumber, bracket)
		}
	}

	sortableIntermediatePieMatches.ImportAllIntoPies()
	pies.ConvertTiesToTiesWL()
}
