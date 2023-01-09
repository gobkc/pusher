package pusher

type subscriber struct {
	topic string
	msg   chan interface{}
}

type subscribers struct {
	subs []*subscriber
}

func (s *subscribers) Subs(topic string, f func(cb SubsCallback)) {
	newSub := subscriber{
		topic: topic,
		msg:   make(chan interface{}, 1000),
	}
	s.subs = append(s.subs, &newSub)
	go func() {
		for {
			select {
			case msg := <-newSub.msg:
				cb := &subscriberCallBack{d: msg}
				f(cb)
			}
		}
	}()
}
