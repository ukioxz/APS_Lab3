package engine

import (
	"fmt"
)

type PrintCommand struct {
	arg string
}

func (p *PrintCommand) Execute(loop Handler) {
  fmt.Println(p.arg)
}
type SplitCommand struct {
	arg string
 }

func (s *SplitCommand) Execute(loop Handler) {

}
