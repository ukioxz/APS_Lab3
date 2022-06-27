package engine

import (
	"strings"
)

func Parse(commandLine string) Command {
	parts := strings.Fields(commandLine)
	if(len(parts) == 0) {
		return &toPrint{arg: "Error. Too less arguments."}
	}

	switch parts[0] {
	case "print":
    if len(parts) == 2 {
      return &toPrint{arg: parts[1]}
    } else {
      return &toPrint{arg: "Error"}
    }
	case "split":
		if len(parts) > 3 {
      return &toPrint{arg: "Error"}
    }
		arg1 := parts[1]
		arg2 := parts[2]
		return splitText(arg1, arg2)
	default:
		return &toPrint{arg: "Error"}
	}
}
