package triglav

import (
	"log"
)

type Worker interface {
	Work()
}

func NewWorker(message *Message) (worker Worker) {
	switch message.Tag {
	case "triglav.update.host":
		worker = NewUpdateHostWorker(message)
	default:
		worker = NewNullWorker(message)
	}
	return
}

type NullWorker struct {
	message *Message
}

func NewNullWorker(message *Message) (worker Worker) {
	worker = &NullWorker{message: message}
	return
}

func (self *NullWorker) Work() {
	log.Println(self.message.String())
}

type UpdateHostWorker struct {
	message *Message
}

func NewUpdateHostWorker(message *Message) (worker Worker) {
	worker = &UpdateHostWorker{message: message}
	return
}

func (self *UpdateHostWorker) Work() {
	log.Println(self.message.String())
}
