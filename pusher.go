package pusher

import (
	"log"
	"sync"
	"time"
)

var (
	once = sync.Once{}
	ps   *pusher
)

type pusher struct {
	s *subscribers
}

func NewPusher() pushPolicy {
	once.Do(func() {
		ps = &pusher{}
		ps.s = &subscribers{}
	})
	return ps
}

func (p *pusher) Push(topic string, data any) {
	go func() {
		delay := 1 * time.Second
		if len(p.s.subs) == 0 {
			for {
				log.Println("no subscriber,waiting...")
				if len(p.s.subs) > 0 {
					p.Push(topic, data)
					return
				}
				time.Sleep(delay)
			}
		}
		for _, sub := range p.s.subs {
			if topic == sub.topic {
				sub.msg <- data
			}
		}
	}()
	return
}

func (p *pusher) getSubscribers() subsPolicy {
	return p.s
}
