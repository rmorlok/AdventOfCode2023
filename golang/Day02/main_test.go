package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineToGame(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Game{
		id: 1, draws: []Draw{
			{red: 4, green: 0, blue: 3},
			{red: 1, green: 2, blue: 6},
			{red: 0, green: 2, blue: 0},
		},
	}, lineToGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"))
	assert.Equal(Game{
		id: 2, draws: []Draw{
			{red: 0, green: 2, blue: 1},
			{red: 1, green: 3, blue: 4},
			{red: 0, green: 1, blue: 1},
		},
	}, lineToGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue"))
	assert.Equal(Game{
		id: 3, draws: []Draw{
			{red: 20, green: 8, blue: 6},
			{red: 4, green: 13, blue: 5},
			{red: 1, green: 5, blue: 0},
		},
	}, lineToGame("Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red"))
	assert.Equal(Game{
		id: 4, draws: []Draw{
			{red: 3, green: 1, blue: 6},
			{red: 6, green: 3, blue: 0},
			{red: 14, green: 3, blue: 15},
		},
	}, lineToGame("Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red"))
	assert.Equal(Game{
		id: 5, draws: []Draw{
			{red: 6, green: 3, blue: 1},
			{red: 1, green: 2, blue: 2},
		},
	}, lineToGame("Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"))
}

func TestStrToDraw(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(Draw{red: 3, green: 1, blue: 6}, strToDraw("1 green, 3 red, 6 blue"))
	assert.Equal(Draw{red: 6, green: 3, blue: 1}, strToDraw("6 red, 1 blue, 3 green"))
}
