package consumer

import (
	"ParallelLabs/lab06/task01/channel"
	"fmt"
	"time"
)

type Consumer struct {
	name string
	dur  time.Duration
}

func New(name string, dur time.Duration) *Consumer {
	return &Consumer{name: name, dur: dur}
}

func (c *Consumer) AsyncConsume(ch *channel.Channel) {
	go func() {
		for {
			if msg := ch.Receive(); msg != nil {
				fmt.Printf("%s consume %+v\n", c.name, msg)
			} else {
				fmt.Printf("%s didn't find anything\n", c.name)
			}
			time.Sleep(c.dur)
		}
	}()
}
