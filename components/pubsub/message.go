package pubsub

import (
	"fmt"
	"time"
)

type Topic string

type Message struct {
	id        string
	channel   Topic
	data      interface{}
	createdAt time.Time
}

func NewMessage(channel Topic, data interface{}) *Message {
	now := time.Now().UTC()
	return &Message{
		id:        fmt.Sprintf("%d", now.UnixNano()),
		channel:   channel,
		data:      data,
		createdAt: now,
	}
}

func (m *Message) String() string {
	return fmt.Sprintf(" Message : %s", m.channel)
}

func (m *Message) SetChannel(topic Topic) {
	m.channel = topic
}

func (m *Message) Channel() Topic {
	return m.channel
}

func (m *Message) Data() interface{} {
	return m.data
}
