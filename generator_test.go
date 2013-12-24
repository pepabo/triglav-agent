package triglav

import (
	. "github.com/r7kamura/gospel"
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	Describe(t, "generator.NewGenerator(name)", func() {
		Context("when the `name` is triglav.update.host", func() {
			args := make(map[string]interface{})
			args["key"] = "value"
			generator := NewGenerator("triglav.update.host", &Queue{}, args)

			It("should return UpdateHostGenertor", func() {
				Expect(reflect.TypeOf(generator).String()).To(Equal, "*triglav.UpdateHostGenertor")
			})
		})

		Context("when the `name` is not either of above", func() {
			args := make(map[string]interface{})
			args["key"] = "value"
			generator := NewGenerator("no.such.generator", &Queue{}, args)

			It("should return NullGenerator", func() {
				Expect(reflect.TypeOf(generator).String()).To(Equal, "*triglav.NullGenerator")
			})
		})
	})
}
