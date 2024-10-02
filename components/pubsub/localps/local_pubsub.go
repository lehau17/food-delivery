package localps

import (
	"context"
	"log"
	"sync"

	"github.com/lehau17/food_delivery/common"
	"github.com/lehau17/food_delivery/components/pubsub"
)

type localpubsub struct {
	messageQueue chan *pubsub.Message
	mapChannel   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewLocalPubsub() *localpubsub {
	p := &localpubsub{
		messageQueue: make(chan *pubsub.Message),
		mapChannel:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}
	p.run()
	return p
}

func (p *localpubsub) Publish(ctx context.Context, channel pubsub.Topic, data *pubsub.Message) error {
	data.SetChannel(channel)
	go func() {
		defer common.AppRecover(ctx)
		p.messageQueue <- data
		log.Println("New event :", data.String())
	}()
	return nil
}

func (p *localpubsub) Subcribe(ctx context.Context, channel pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	p.locker.Lock()
	if _, ok := p.mapChannel[channel]; ok {
		val := append(p.mapChannel[channel], c)
		p.mapChannel[channel] = val
	} else {
		p.mapChannel[channel] = []chan *pubsub.Message{c}
	}
	p.locker.Unlock()
	return c, func() {
		log.Println("Unscuple channel")
		if chans, ok := p.mapChannel[channel]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)
					p.locker.Lock()
					p.mapChannel[channel] = chans
					p.locker.Unlock()
					break
				}
			}
		}
	}

}

func (p *localpubsub) run() error {
	go func() {
		for {
			mess := <-p.messageQueue
			log.Println("Message dequeue", mess.String())

			if subs, ok := p.mapChannel[mess.Channel()]; ok {

				for i := range subs {
					go func(c chan *pubsub.Message) {
						log.Println("Run chay for message")
						c <- mess
					}(subs[i])
				}
			}
		}
	}()
	return nil
}
