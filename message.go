package triglav

import (
	"encoding/json"
)

type Message struct {
	Tag  string
	Body map[string]interface{}
}

func (self *Message) String() (result []byte, err error) {
	result, err = json.Marshal(self.Body)
	return
}
