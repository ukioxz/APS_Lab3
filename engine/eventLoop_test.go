package engine

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestEventLoop(s *testing.T) {

  res1 := &toPrint{arg: "hello"}
  res2 := &toPrint{arg: "jinx is jinx"}
  res3 := splitText("3,goda,do,zavoda", ",")

  eventLoop := new(EventLoop)
  eventLoop.Start()
  assert.Equal(s, false, eventLoop.stopRequest)
  assert.Equal(s, 0, len(eventLoop.messageQueue.com))

  eventLoop.Post(res1)
  eventLoop.Post(res2)
  eventLoop.Post(res3)

  assert.Equal(s, 3, len(eventLoop.messageQueue.com))

  eventLoop.AwaitFinish()
  assert.Equal(s, true, eventLoop.stopRequest)
  assert.Equal(s, 0, len(eventLoop.messageQueue.com))

  assert.Equal(s, res1, &toPrint{arg: "hello"})
  assert.Equal(s, res2, &toPrint{arg: "jinx is jinx"})
  assert.Equal(s, res3, splitText("3,goda,do,zavoda", ","))
}
