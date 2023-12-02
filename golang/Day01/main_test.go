package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLineToNumberCalibration1(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(12, lineToNumberCalibration1("1abc2"))
	assert.Equal(38, lineToNumberCalibration1("pqr3stu8vwx"))
	assert.Equal(15, lineToNumberCalibration1("a1b2c3d4e5f"))
	assert.Equal(77, lineToNumberCalibration1("treb7uchet"))
}

func TestLineToNumberCalibration2(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(29, lineToNumberCalibration2("two1nine"))
	assert.Equal(83, lineToNumberCalibration2("eightwothree"))
	assert.Equal(13, lineToNumberCalibration2("abcone2threexyz"))
	assert.Equal(24, lineToNumberCalibration2("xtwone3four"))
	assert.Equal(42, lineToNumberCalibration2("4nineeightseven2"))
	assert.Equal(14, lineToNumberCalibration2("zoneight234"))
	assert.Equal(76, lineToNumberCalibration2("7pqrstsixteen"))
}
