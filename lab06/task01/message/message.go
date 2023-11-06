package message

import "strconv"

type Message struct {
	Text      string
	Num       uint
	Timestamp int64
}

func New(text string, num uint, timestamp int64) *Message {
	return &Message{Text: text, Num: num, Timestamp: timestamp}
}

func (m *Message) Key() string {
	return strconv.Itoa(int(m.Num))
}
