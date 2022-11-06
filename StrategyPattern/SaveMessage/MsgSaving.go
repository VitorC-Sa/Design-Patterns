package savemessage

import "fmt"

type message struct {
	originName   string
	dstName      string
	msg          string
	outputMethod OutputBehavior
}

func NewMessage(msg, origin, dst string) *message {
	return &message{
		originName:   origin,
		dstName:      dst,
		msg:          msg,
		outputMethod: Log{},
	}
}

func (m *message) SetMessage(msg, origin, dst string) {
	m.originName = origin
	m.dstName = dst
	m.msg = msg
}

func (m *message) SetOutputMethod(ot OutputBehavior) { m.outputMethod = ot }

func (m message) SendMessage() {
	r := fmt.Sprintf("%s: Dear %s, %s", m.originName, m.dstName, m.msg)

	m.outputMethod.Save(r)
}
