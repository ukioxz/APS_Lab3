/*package main

import (
	"bufio"
	"log"
	"os"
	"github.com/ukioxz/APS_Lab4/engine"
)

func main() {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()

	if input, err := os.Open(instructions.txt); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := Parse(commandLine) // parse the line to get a Command (should be created)
		eventLoop.Post(cmd)
  }
}
eventLoop.AwaitFinish()
}*/
package main

import (
  "bufio"
  "os"
  "github.com/ukioxz/APS_Lab4/engine"
  "strings"
)

func main() {
  eventLoop := new(engine.EventLoop)
  eventLoop.Start()

  if input, err := os.Open("instructions.txt"); err == nil {
    defer input.Close()
    scanner := bufio.NewScanner(input)
    for scanner.Scan() {
      commandLine := scanner.Text()
      cmd := Parse(commandLine) // parse the line to get a Command (should be created)
      eventLoop.Post(cmd)
    }
  }
eventLoop.AwaitFinish()
}
