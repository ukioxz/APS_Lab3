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
	qu *queue
	stopSignal  chan struct{}
	stopRequest bool
}

func (qu *queue) empty() bool {
	if(len(qu.com) == 0){
		return false
	}
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

func (qu *queue) pull(cmd Command) {
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
	qu.data = qu.com[1:]
	return result
}

