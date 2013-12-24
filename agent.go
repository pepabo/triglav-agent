package triglav

type Agent struct {
}

func (self *Agent) Run() {
	quit := make(chan bool)
	queue := &Queue{
		queue: make(chan *Message),
		quit:  quit,
	}
	consumer := &Consumer{
		queue:   queue,
		oneShot: false,
	}
	producer := &Producer{
		queue: queue,
	}

	go consumer.Run()
	go producer.Run()

	<-quit
}
