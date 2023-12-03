package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Schematic [][]rune

func (s *Schematic) extractNumbers() []*Number {
	numbers := make([]*Number, 0)
	var start int
	var inProgress string

	for y := 0; y < len(*s); y++ {
		row := (*s)[y]
		for x := 0; x < len(row); x++ {
			v := row[x]

			if v >= '0' && v <= '9' {
				if len(inProgress) == 0 {
					start = x
				}

				inProgress += string(v)

				// End of row
				if x == len(row)-1 {
					val, _ := strconv.Atoi(inProgress)
					numbers = append(numbers, &Number{
						val:       val,
						start:     Point{x: start, y: y},
						end:       Point{x: x, y: y},
						schematic: s,
					})
					inProgress = ""
					start = 0
				}
			} else {
				if len(inProgress) > 0 {
					val, _ := strconv.Atoi(inProgress)
					numbers = append(numbers, &Number{
						val:       val,
						start:     Point{x: start, y: y},
						end:       Point{x: x, y: y},
						schematic: s,
					})
				}
				inProgress = ""
				start = 0
			}
		}
	}

	return numbers
}

func (s *Schematic) extractSymbols() []*Symbol {
	symbols := make([]*Symbol, 0)

	for y := 0; y < len(*s); y++ {
		row := (*s)[y]
		for x := 0; x < len(row); x++ {
			v := row[x]

			if (v < '0' || v > '9') && v != '.' {
				symbols = append(symbols, &Symbol{
					val:       v,
					start:     Point{x: x, y: y},
					end:       Point{x: x, y: y},
					schematic: s,
				})
			}
		}
	}

	return symbols
}

type Point struct {
	x int
	y int
}

type Number struct {
	val       int
	start     Point
	end       Point
	schematic *Schematic
}

type Symbol struct {
	val       rune
	start     Point
	end       Point
	schematic *Schematic
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (n *Number) adjacentSymbols() []rune {
	vals := []rune{}
	for y := max(n.start.y-1, 0); y <= min(n.end.y+1, len(*n.schematic)-1); y++ {
		row := (*n.schematic)[y]
		for x := max(n.start.x-1, 0); x < min(n.end.x+1, len(row)-1); x++ {
			v := row[x]

			if (v < '0' || v > '9') && v != '.' {
				vals = append(vals, v)
			}
		}
	}

	return vals
}

func (s *Symbol) adjacentNumbers() []*Number {
	numbers := make([]*Number, 0)
	var start int
	var inProgress string

	for y := max(s.start.y-1, 0); y <= min(s.end.y+1, len(*s.schematic)); y++ {
		row := (*s.schematic)[y]
		rowStart := max(s.start.x-1, 0)
		for rowStart > 0 && row[rowStart] >= '0' && row[rowStart] <= '9' {
			rowStart--
		}

		rowEnd := min(s.end.x+1, len(row))
		for rowEnd+1 < len(row) && row[rowEnd] >= '0' && row[rowEnd] <= '9' {
			rowEnd++
		}

		for x := rowStart; x <= rowEnd; x++ {
			v := row[x]

			if v >= '0' && v <= '9' {
				if len(inProgress) == 0 {
					start = x
				}

				inProgress += string(v)

				// End of row
				if x == len(row)-1 {
					val, _ := strconv.Atoi(inProgress)
					numbers = append(numbers, &Number{
						val:       val,
						start:     Point{x: start, y: y},
						end:       Point{x: x, y: y},
						schematic: s.schematic,
					})
					inProgress = ""
					start = 0
				}
			} else {
				if len(inProgress) > 0 {
					val, _ := strconv.Atoi(inProgress)
					numbers = append(numbers, &Number{
						val:       val,
						start:     Point{x: start, y: y},
						end:       Point{x: x, y: y},
						schematic: s.schematic,
					})
				}
				inProgress = ""
				start = 0
			}
		}
	}

	return numbers
}

func (s *Symbol) isGear() bool {
	return s.val == '*' && len(s.adjacentNumbers()) == 2
}

func (s *Symbol) gearRatio() int {
	ns := s.adjacentNumbers()
	if len(ns) == 2 {
		return ns[0].val * ns[1].val
	} else {
		return 0
	}
}

func main() {
	data := []string{}

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
		data = append(data, l)
	}

	s := createSchematic(data)
	sum1 := sumOfNoAdjacent(s)
	sum2 := sumOfGearRatios(s)

	fmt.Printf("Sum vals (1): %d\n", sum1)
	fmt.Printf("Sum gear ratios (2): %d\n", sum2)
}

func createSchematic(data []string) *Schematic {
	shematic := Schematic{}
	for _, s := range data {
		shematic = append(shematic, []rune(s))
	}

	return &shematic
}

func sumOfNoAdjacent(s *Schematic) int {
	var sum int
	for _, n := range s.extractNumbers() {
		if len(n.adjacentSymbols()) > 0 {
			sum += n.val
		}
	}

	return sum
}

func sumOfGearRatios(s *Schematic) int {
	var sum int
	for _, s := range s.extractSymbols() {
		if s.isGear() {
			sum += s.gearRatio()
		}
	}

	return sum
}
