package main

import (
	"ParallelLabs/lab06/task01/channel"
	"ParallelLabs/lab06/task01/consumer"
	"ParallelLabs/lab06/task01/producer"
	"time"
)

func main() {
	ch := channel.New(3)
	{
		prd := producer.New("Alice", time.Millisecond*100)
		prd.AsyncProduce(ch)
	}

	{
		cns1 := consumer.New("Bob", time.Millisecond*200)
		cns1.AsyncConsume(ch)

		cns2 := consumer.New("Tom", time.Millisecond*200)
		cns2.AsyncConsume(ch)
	}

	var done chan struct{}
	<-done
}
