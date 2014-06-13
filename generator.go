package triglav

import (
	"github.com/kentaro/verity"
	"log"
	"reflect"
	"time"
)

type Generator interface {
	Generate(options map[string]interface{})
}

func NewGenerator(tag string, queue *Queue) (generator Generator) {
	switch tag {
	case "triglav.update.host":
		generator = NewUpdateHostGenerator(tag, queue)
	default:
		generator = NewNullGenerator(tag, queue)
	}
	return
}

type NullGenerator struct {
	tag   string
	queue *Queue
}

func NewNullGenerator(tag string, queue *Queue) (generator Generator) {
	generator = &NullGenerator{
		tag:   tag,
		queue: queue,
	}
	return
}

func (self *NullGenerator) Generate(options map[string]interface{}) {

}

type UpdateHostGenertor struct {
	tag   string
	queue *Queue
}

func NewUpdateHostGenerator(tag string, queue *Queue) (generator Generator) {
	generator = &UpdateHostGenertor{
		tag:   tag,
		queue: queue,
	}
	return
}

func (self *UpdateHostGenertor) Generate(options map[string]interface{}) {
	sec := reflect.ValueOf(options["update-host-interval"])

	if sec.Kind() != reflect.Int {
		log.Fatal("[update.host.generator] `interval` must be an Int value.")
		return
	}

	self.generateUpdateHostMessage()

	timer := time.NewTicker(time.Duration(sec.Int()) * time.Second)
	defer timer.Stop()

	for _ = range timer.C {
		self.generateUpdateHostMessage()
	}
}

func (self *UpdateHostGenertor) generateUpdateHostMessage() {
	hostInfo, err := verity.Collect()

	if err != nil {
		log.Fatalf("[update.host.generator] failed to collect host information: %s", err)
	}

	message := &Message{
		Tag:  self.tag,
		Body: hostInfo,
	}

	self.queue.Push(message)
}
