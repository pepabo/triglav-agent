package triglav

import (
	"log"
)

type Consumer struct {
	queue   *Queue
	oneShot bool
}

func (self *Consumer) Run() {
	for {
		message := self.queue.Pop()

		if self.oneShot {
			self.consume(message)
			self.queue.Quit()
			return
		} else {
			go self.consume(message)
		}
	}
}

func (self *Consumer) consume(message *Message) {
	log.Println(message)
}
