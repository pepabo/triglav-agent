package triglav

import (
	"log"
	"reflect"
)

type Consumer struct {
	queue   *Queue
	oneShot bool
}

func (self *Consumer) Run(options map[string]interface{}) {
	oneShot := reflect.ValueOf(options["one-shot"])

	if oneShot.Kind() != reflect.Bool {
		log.Panic("[consumer] `one-shot` must be a Bool value.")
		return
	}

	for {
		message := self.queue.Pop()

		if oneShot.Bool() {
			self.consume(message)
			self.queue.Quit()
			return
		} else {
			go self.consume(message)
		}
	}
}

func (self *Consumer) consume(message *Message) {
	worker := NewWorker(message)
	worker.Work()
}
