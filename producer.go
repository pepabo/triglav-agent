package triglav

type Producer struct {
	queue *Queue
}

func (self *Producer) Run() {
	message := &Message{
		Tag: "triglav",
		Body: map[string]interface{}{
			"key": "value",
		},
	}

	self.queue.Push(message)
	self.queue.Quit()
}
