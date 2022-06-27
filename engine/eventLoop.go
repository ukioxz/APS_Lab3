package engine

import (
  "sync"
)

type Command interface {
  Execute(handler Handler)
}

type Handler interface {
  Post(cmd Command)
}

type queue struct {
  sync.Mutex
  com []Command
  getSignal chan struct{}
  getStatus bool
}

type EventLoop struct {
  messageQueue *queue
  stopSignal  chan struct{}
  stopRequest bool
}

func (qu *queue) empty() bool {
  return len(qu.com) == 0
}

func (qu *queue) push(cmd Command) {
  qu.Lock()
  defer qu.Unlock()
  qu.com = append(qu.com, cmd)

  if qu.getStatus {
    qu.getStatus = false
    qu.getSignal <- struct{}{}
  }
}

func (qu *queue) pull() Command {
  qu.Lock()
  defer qu.Unlock()
  if qu.empty() {
    qu.getStatus = true
    qu.Unlock()
    <-qu.getSignal
    qu.Lock()
  }
  result := qu.com[0]
  qu.com[0] = nil
  qu.com = qu.com[1:]
  return result
}

func (loop *EventLoop) Start() {                        //create queue
  loop.messageQueue = &queue{getSignal: make(chan struct{})}
  loop.stopSignal = make(chan struct{})
  go func() {
    for !loop.stopRequest || !loop.messageQueue.empty() {
      cmd := loop.messageQueue.pull()
      cmd.Execute(loop)
    }
    loop.stopSignal <- struct{}{}
  }()
}

func (loop *EventLoop) Post(cmd Command) {        //add command to queue
  loop.messageQueue.push(cmd)
}

func (loop *EventLoop) AwaitFinish() {      //all commands are finished
  loop.stopRequest = true
  <-loop.stopSignal
}
