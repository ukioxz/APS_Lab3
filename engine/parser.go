package engine

import (
  "strings"
)

func Parse(commandLine string) Command {
  parts := strings.Fields(commandLine)
  switch parts[0] {
  case "print":
    return &PrintCommand{arg: parts[1]}
  default:
    return &PrintCommand{arg: "Error"}
  }
}
