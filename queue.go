package triglav

type Queue struct {
	queue chan *Message
	quit  chan bool
}

func (self *Queue) Push(message *Message) {
	self.queue <- message
}

func (self *Queue) Pop() (message *Message) {
	message = <-self.queue
	return
}

func (self *Queue) Quit() {
	self.quit <- true
}
