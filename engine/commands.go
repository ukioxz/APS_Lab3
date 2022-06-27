package main

import (
	"fmt"
)

type Command interface {
	Execute(handler Handler)
}

type PrintCommand struct {
	arg string
}

func (pc PrintCommand) Execute(h Handler) {
	fmt.Println(string(pc))
}

type SplitCommand struct {
	arg string
 }

/*func (p *PrintCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}*/

func (s *SplitCommand) Execute(loop engine.Handler) {

}
