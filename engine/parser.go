package engine

import (
	"strings"
)

func Parse(commandLine string) Command {
	parts := strings.Fields(commandLine)
	switch parts[0] {
	case "print":
		return &toPrint{arg: parts[1]}
	case "split":
		arg1 := parts[1]
		arg2 := parts[2]
		return splitText(arg1, arg2)
	default:
		return &toPrint{arg: "Error"}
	}
}
