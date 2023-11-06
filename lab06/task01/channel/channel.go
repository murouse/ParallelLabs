package channel

import (
	"fmt"
	"sync"
)

type Messager interface {
	Key() string
}

type Channel struct {
	mu    *sync.Mutex
	queue []Messager
	recvx uint
	sendx uint
}

func New(capacity uint) *Channel {
	return &Channel{
		mu:    &sync.Mutex{},
		queue: make([]Messager, capacity),
	}
}

func (c *Channel) Send(elem Messager) {
	if c.queue[c.sendx] != nil {
		fmt.Printf("Message %q is lost\n", elem.Key())
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	c.queue[c.sendx] = elem
	c.sendx += 1
	c.sendx %= uint(len(c.queue))
}

func (c *Channel) Receive() Messager {
	if c.queue[c.recvx] == nil {
		return nil
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	elem := c.queue[c.recvx]
	c.queue[c.recvx] = nil
	c.recvx += 1
	c.recvx %= uint(len(c.queue))
	return elem
}
