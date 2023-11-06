package producer

import (
	"ParallelLabs/lab06/task01/channel"
	"ParallelLabs/lab06/task01/message"
	"fmt"
	"time"
)

type Producer struct {
	name string
	dur  time.Duration
}

func New(name string, dur time.Duration) *Producer {
	return &Producer{name: name, dur: dur}
}

func (p *Producer) AsyncProduce(ch *channel.Channel) {
	go func() {
		var num uint
		for {
			num++
			msg := message.New(
				fmt.Sprintf("I'm %s, shall we meet?", p.name),
				num,
				time.Now().UnixMilli(),
			)
			ch.Send(msg)
			time.Sleep(p.dur)
		}
	}()
}
