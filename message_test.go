package triglav

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestCollector(t *testing.T) {
	Describe(t, "message.String()", func() {
		message := &Message{
			Tag: "test",
			Body: map[string]interface{}{
				"key": "value",
			},
		}
		body, err := message.String()

		It("should convert its body to string value", func() {
			Expect(string(body)).To(Equal, "{\"key\":\"value\"}")
		})

		It("should return no errors", func() {
			Expect(err).To(Equal, nil)
		})
	})
}
