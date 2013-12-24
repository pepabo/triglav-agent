package triglav

type Producer struct {
	queue *Queue
}

func (self *Producer) Run() {
	message := &Message{
		Tag: "triglav.update.host",
		Body: map[string]interface{}{
			"key": "value",
		},
	}

	self.queue.Push(message)
}
