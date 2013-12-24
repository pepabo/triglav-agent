package triglav

import (
	"log"
	"reflect"
	"time"
)

type Generator interface {
	Generate()
}

func NewGenerator(name string, queue *Queue, args map[string]interface{}) (generator Generator) {
	switch name {
	case "triglav.update.host":
		generator = NewUpdateHostGenerator(queue, args)
	default:
		generator = NewNullGenerator(queue, args)
	}
	return
}

type NullGenerator struct {
	queue *Queue
	args  map[string]interface{}
}

func NewNullGenerator(queue *Queue, args map[string]interface{}) (generator Generator) {
	generator = &NullGenerator{
		queue: queue,
		args:  args,
	}
	return
}

func (self *NullGenerator) Generate() {

}

type UpdateHostGenertor struct {
	queue *Queue
	args  map[string]interface{}
}

func NewUpdateHostGenerator(queue *Queue, args map[string]interface{}) (generator Generator) {
	generator = &UpdateHostGenertor{
		queue: queue,
		args:  args,
	}
	return
}

func (self *UpdateHostGenertor) Generate() {
	sec := reflect.ValueOf(self.args["interval"])

	if sec.Kind() != reflect.Int {
		log.Panic("[triglav.update.host] `interval` must be an Int value.")
		return
	}

	ticker := time.NewTicker(time.Duration(sec.Int()) * time.Second).C

	for {
		select {
		case <-ticker:
			message := &Message{
				Tag: "triglav.update.host",
				Body: map[string]interface{}{
					"key": "value",
				},
			}
			self.queue.Push(message)
		}
	}
}
