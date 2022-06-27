package engine

import (
	"fmt"
	"strings"
)

type toPrint struct {
	arg string
}

type toSplit struct {
	arg1, arg2 string
}
func printText(arg string) *toPrint {
	return &toPrint{
		arg: arg,
	}
}
func (p *toPrint) Execute(loop Handler) {
	fmt.Println(p.arg)
}

func splitText(arg1 string, arg2 string) *toSplit {
	return &toSplit{
		arg1: arg1,
		arg2: arg2,
	}
}

func (split *toSplit) Execute(h Handler) {
	arrPart := strings.Split(split.arg1, split.arg2)
	res := strings.Join(arrPart, "\n")
	h.Post(printText(res))
}
