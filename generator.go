package triglav

import (
	"log"
	"reflect"
	"time"
)

type Generator interface {
	Generate(options map[string]interface{})
}

func NewGenerator(name string, queue *Queue) (generator Generator) {
	switch name {
	case "triglav.update.host":
		generator = NewUpdateHostGenerator(queue)
	default:
		generator = NewNullGenerator(queue)
	}
	return
}

type NullGenerator struct {
	queue *Queue
}

func NewNullGenerator(queue *Queue) (generator Generator) {
	generator = &NullGenerator{
		queue: queue,
	}
	return
}

func (self *NullGenerator) Generate(options map[string]interface{}) {

}

type UpdateHostGenertor struct {
	queue *Queue
}

func NewUpdateHostGenerator(queue *Queue) (generator Generator) {
	generator = &UpdateHostGenertor{
		queue: queue,
	}
	return
}

func (self *UpdateHostGenertor) Generate(options map[string]interface{}) {
	sec := reflect.ValueOf(options["update-host-interval"])

	if sec.Kind() != reflect.Int {
		log.Panic("[update.host.generator] `interval` must be an Int value.")
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
