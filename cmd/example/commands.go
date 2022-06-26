package main

import (
	"fmt"
)

type PrintCommand struct {
	arg string
 }

 func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
 }

