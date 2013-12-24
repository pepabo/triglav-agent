package triglav

import (
	. "github.com/r7kamura/gospel"
	"reflect"
	"testing"
)

func TestGenerator(t *testing.T) {
	Describe(t, "generator.NewGenerator(name)", func() {
		Context("when the `name` is triglav.update.host", func() {
			generator := NewGenerator("triglav.update.host", &Queue{})

			It("should return UpdateHostGenertor", func() {
				Expect(reflect.TypeOf(generator).String()).To(Equal, "*triglav.UpdateHostGenertor")
			})
		})

		Context("when the `name` is not either of above", func() {
			generator := NewGenerator("no.such.generator", &Queue{})

			It("should return NullGenerator", func() {
				Expect(reflect.TypeOf(generator).String()).To(Equal, "*triglav.NullGenerator")
			})
		})
	})
}
