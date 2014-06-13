package triglav

type Producer struct {
	queue *Queue
}

var generatorTags = []string{
	"triglav.update.host",
}

func (self *Producer) Run(options map[string]interface{}) {
	quit := make(chan struct{})
	producerQueue := &Queue{
		queue: make(chan *Message),
		quit:  quit,
	}
	defer producerQueue.Quit()

	for _, tag := range generatorTags {
		generator := NewGenerator(tag, producerQueue)
		go generator.Generate(options)
	}

	for {
		message := producerQueue.Pop()
		self.produce(message)
	}
}

func (self *Producer) produce(message *Message) {
	self.queue.Push(message)
}
