package engine

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestParsePrint1(s *testing.T) {
  commandTxt := "print he"
	res:= Parse(commandTxt)
	assert.Equal(s, &toPrint{arg: "he"}, res)
}

func TestParsePrint2(s *testing.T) {
  commandTxt := "print jinx is jinx"
	res:= Parse(commandTxt)
	assert.Equal(s, &toPrint{arg: "Error"}, res)
}

func TestParsePrint3(s *testing.T) {
  commandTxt := "printErr shark2"
	res:= Parse(commandTxt)
	assert.Equal(s, &toPrint{arg: "Error"}, res)
}
