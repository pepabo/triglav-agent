package triglav

import (
	. "github.com/r7kamura/gospel"
	"reflect"
	"testing"
)

func TestWorker(t *testing.T) {
	Describe(t, "worker.NewWorker(message)", func() {
		Context("when message.tag == triglav.update.host", func() {
			message := &Message{
				Tag: "triglav.update.host",
				Body: map[string]interface{}{
					"key": "value",
				},
			}
			worker := NewWorker(message)

			It("should return UpdateHostWorker", func() {
				Expect(reflect.TypeOf(worker).String()).To(Equal, "*triglav.UpdateHostWorker")
			})
		})

		Context("when message.tag is not either of above", func() {
			message := &Message{
				Tag: "no.such.tag",
				Body: map[string]interface{}{
					"key": "value",
				},
			}
			worker := NewWorker(message)

			It("should return NullWorker", func() {
				Expect(reflect.TypeOf(worker).String()).To(Equal, "*triglav.NullWorker")
			})
		})
	})
}
