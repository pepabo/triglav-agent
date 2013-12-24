package triglav

type Producer struct {
	queue *Queue
}

func (self *Producer) Run() {
	quit := make(chan bool)
	producerQueue := &Queue{
		queue: make(chan *Message),
		quit: quit,
	}
	generator := NewGenerator("triglav.update.host", producerQueue, map[string]interface{}{"interval":1})

	go generator.Generate()

	for {
		message := producerQueue.Pop()
		self.produce(message)
	}
}

func (self *Producer) produce(message *Message) {
	self.queue.Push(message)
}
