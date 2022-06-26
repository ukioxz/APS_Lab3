package main

import (
	"fmt"
)

type PrintCommand struct {
	arg string
}

type SplitCommand struct {
	arg string
 } 

func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}

func (s *SplitCommand) Execute(loop engine.Handler) {
	
}
