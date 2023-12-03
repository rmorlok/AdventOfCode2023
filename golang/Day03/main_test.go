package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNumbers(t *testing.T) {
	assert := assert.New(t)

	data := strings.Split(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "\n")
	s := createSchematic(data)
	ns := s.extractNumbers()

	assert.Equal(10, len(ns))

	assert.Equal(467, ns[0].val)
	assert.Equal([]rune{'*'}, ns[0].adjacentSymbols())

	assert.Equal(114, ns[1].val)
	assert.Equal([]rune{}, ns[1].adjacentSymbols())

	assert.Equal(35, ns[2].val)
	assert.Equal([]rune{'*'}, ns[2].adjacentSymbols())

	assert.Equal(633, ns[3].val)
	assert.Equal([]rune{'#'}, ns[3].adjacentSymbols())

	assert.Equal(4361, sumOfNoAdjacent(s))
}

func TestNumbersReal(t *testing.T) {
	assert := assert.New(t)

	data := strings.Split(`121......992...............*.......%585....814............936.......102..#353.........*.....140.........*..434..301..................%..315.
45......$..651....$...............................*.........526.............41......*......302...................530..........@819.......463`, "\n")

	s := createSchematic(data)
	ns := s.extractNumbers()

	assert.Equal(19, len(ns))

	assert.Equal(121, ns[0].val)
	assert.Equal([]rune{}, ns[0].adjacentSymbols())

	assert.Equal(463, ns[18].val)
	assert.Equal([]rune{}, ns[18].adjacentSymbols())
}

func TestSymbols(t *testing.T) {
	assert := assert.New(t)

	data := strings.Split(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`, "\n")
	s := createSchematic(data)
	syms := s.extractSymbols()

	assert.Equal(6, len(syms))

	assert.Equal('*', syms[0].val)
	assert.Equal(2, len(syms[0].adjacentNumbers()))

	assert.Equal('#', syms[1].val)
	assert.Equal(1, len(syms[1].adjacentNumbers()))

	assert.Equal('*', syms[2].val)
	assert.Equal(1, len(syms[2].adjacentNumbers()))

	assert.Equal('*', syms[5].val)
	assert.Equal(2, len(syms[5].adjacentNumbers()))

	assert.Equal(467835, sumOfGearRatios(s))
}
