package triglav

import (
	"encoding/json"
)

type Message struct {
	Tag  string
	Body map[string]interface{}
}

func (self *Message) String() (result string, err error) {
	bytes, err := json.Marshal(self.Body)

	if err != nil {
		result = "Malformed JSON"
	} else {
		result = string(bytes)
	}

	return
}
