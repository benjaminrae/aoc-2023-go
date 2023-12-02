package day2

import (
	"adventofcode/pkg"
	"log"
	"reflect"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Reveal struct {
	Blue  int
	Red   int
	Green int
}

type Game struct {
	Reveals []Reveal
	Id      int
}

func LineToGame(line string) Game {
	var game = Game{}

	var gameAndContents = strings.Split(line, ": ")
	var gameID = gameAndContents[0]
	var gameContents = gameAndContents[1]

	var id, err = strconv.Atoi(strings.Split(gameID, " ")[1])

	if err != nil {
		log.Fatal(err)
	}

	game.Id = id

	var rawReveals = strings.Split(gameContents, "; ")
	var reveals = []Reveal{}

	for _, rawReveal := range rawReveals {
		var rawColors = strings.Split(rawReveal, ", ")
		var reveal = Reveal{0, 0, 0}

		for _, rawColor := range rawColors {
			var totalAndColor = strings.Split(rawColor, " ")
			var rawTotal = totalAndColor[0]
			var color = totalAndColor[1]

			total, err := strconv.Atoi(rawTotal)
			if err != nil {
				log.Fatal(err)
			}

			caser := cases.Title(language.English)

			field := reflect.ValueOf(&reveal).Elem().FieldByName(caser.String(color))
			if field.IsValid() && field.Kind() == reflect.Int {
				field.SetInt(int64(total))
			} else {
				log.Printf("Invalid field or non-int field: %s", color)
			}

		}

		reveals = append(reveals, reveal)
	}

	game.Reveals = reveals

	return game
}

func (g Game) GetHighestReveal() Reveal {
	var reveal = Reveal{0, 0, 0}

	for _, r := range g.Reveals {
		if r.Blue > reveal.Blue {
			reveal.Blue = r.Blue
		}

		if r.Red > reveal.Red {
			reveal.Red = r.Red
		}

		if r.Green > reveal.Green {
			reveal.Green = r.Green
		}
	}

	return reveal
}

func Solve(inputFile string, maxReveal Reveal) int {
	var input = pkg.GetFileContents(inputFile)

	var lines = strings.Split(input, "\n")

	var validGames = []Game{}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var game = LineToGame(line)

		var highestReveal = game.GetHighestReveal()

		if highestReveal.Blue <= maxReveal.Blue &&
			highestReveal.Red <= maxReveal.Red &&
			highestReveal.Green <= maxReveal.Green {
			validGames = append(validGames, game)
		}
	}

	var sum = 0

	for _, game := range validGames {
		sum += game.Id
	}

	return sum
}
