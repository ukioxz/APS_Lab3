package engine

import (
	"strings"
)

func Parse(in string) Command {
	fields := strings.Fields(in)

	if len(fields) < 2 {
		return &ErrorMessage("incorrect number of args")
	}

	name := fields[0]
	args := fields[1:]

	switch name {
	case "print":
		return PrintCommand(parts[1]))
	case "split":
		arg1 := args[0]

		arg2 := args[1]

		return NewSplitCommand(arg1, arg2)
	default:
		return &ErrorMessage("Error command does not exist (" + in + ")")
	}
}
