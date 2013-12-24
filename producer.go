package triglav

type Producer struct {
	queue *Queue
}

func (self *Producer) Run(options map[string]interface{}) {
	quit := make(chan bool)
	producerQueue := &Queue{
		queue: make(chan *Message),
		quit:  quit,
	}
	generator := NewGenerator("triglav.update.host", producerQueue)

	go generator.Generate(options)

	for {
		message := producerQueue.Pop()
		self.produce(message)
	}
}

func (self *Producer) produce(message *Message) {
	self.queue.Push(message)
}
