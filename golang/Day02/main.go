package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	draws []Draw
}

func (g *Game) FewestRequired() Draw {
	var r Draw
	for _, d := range g.draws {
		if d.red > r.red {
			r.red = d.red
		}

		if d.green > r.green {
			r.green = d.green
		}

		if d.blue > r.blue {
			r.blue = d.blue
		}
	}

	return r
}

func (g *Game) PossibleWith(t Draw) bool {
	for _, d := range g.draws {
		if d.green > t.green || d.red > t.red || d.blue > t.blue {
			return false
		}
	}

	return true
}

type Draw struct {
	red   int
	green int
	blue  int
}

func (d Draw) Power() int {
	return d.blue * d.green * d.red
}

func main() {
	sum1, sum2 := 0, 0
	sum1Set := Draw{red: 12, green: 13, blue: 14}

	readFile, err := os.Open("data/input.txt")
	defer readFile.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		l := fileScanner.Text()
		g := lineToGame(l)
		if g.PossibleWith(sum1Set) {
			sum1 += g.id
		}
		sum2 += g.FewestRequired().Power()
	}

	fmt.Printf("Sum ids (1): %d\n", sum1)
	fmt.Printf("Total powers (2): %d\n", sum2)
}

func lineToGame(l string) Game {
	result := strings.Split(l, ": ")
	gamePlusId, rest := result[0], result[1]

	result = strings.Split(gamePlusId, " ")
	gameId, _ := strconv.Atoi(result[1])

	draws := []Draw{}
	for _, s := range strings.Split(rest, "; ") {
		draws = append(draws, strToDraw(s))
	}

	return Game{
		id:    gameId,
		draws: draws,
	}
}

func strToDraw(s string) Draw {
	var d Draw
	for _, colorCount := range strings.Split(s, ", ") {
		result := strings.Split(colorCount, " ")
		count, _ := strconv.Atoi(result[0])
		color := result[1]

		switch color {
		case "red":
			d.red = count
		case "green":
			d.green = count
		case "blue":
			d.blue = count
		}
	}

	return d
}
