package triglav

type Agent struct {
}

func (self *Agent) Run(options map[string]interface{}) {
	quit := make(chan struct{})
	queue := &Queue{
		queue: make(chan *Message),
		quit:  quit,
	}
	consumer := &Consumer{
		queue: queue,
	}
	producer := &Producer{
		queue: queue,
	}

	go consumer.Run(options)
	go producer.Run(options)

	<-quit
}
