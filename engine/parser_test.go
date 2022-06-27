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

func TestParseSplit1(s *testing.T) {
  commandTxt := "split go,the,best ,"
	res:= Parse(commandTxt)
  res2:= splitText("go,the,best", ",")
	assert.Equal(s, res2, res)
}

func TestParseSplit2(s *testing.T) {
  commandTxt := "splitttt go,the,best ,"
	res:= Parse(commandTxt)
	assert.Equal(s, &toPrint{arg: "Error"}, res)
}

func TestParseSplit3(s *testing.T) {
  commandTxt := "split go,the,best . , 5"
	res:= Parse(commandTxt)
	assert.Equal(s, &toPrint{arg: "Error"}, res)
}

func TestParseSplit4(s *testing.T) {
  commandTxt := "split '3''goda''do''zavoda' ''"
	res:= Parse(commandTxt)
  res2:= splitText("'3''goda''do''zavoda'", "''")
	assert.Equal(s, res2, res)
}

func TestParseSplit5(s *testing.T) {
  commandTxt := ""
	res:= Parse(commandTxt)
	assert.Equal(s, &toPrint{arg: "Error. Too less arguments."}, res)
}
