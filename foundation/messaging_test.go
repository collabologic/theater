package foundation

import (
	"fmt"
	"testing"
)

type TestElement1 struct {
	TheElement
	TheObserver
	TheNotifier
}

func (te *TestElement1) send() {
	nte := Notifier(te)
	te.Notify(Message{
		TESTMESG,
		te.ID,
		&nte,
		nil,
	})
}

type TestElement2 struct {
	TheElement
	TheObserver
	TheNotifier
	Running bool
}

func (te *TestElement2) receive(m *Message) {
	fmt.Println("RECEIVE!")
	fmt.Println(m)
	te.Running = false
}

const (
	TESTMESG MsgIdentifier = "TEST"
)

var te1 TestElement1
var te2 TestElement2

func Test_Messaging_1(t *testing.T) {
	mq := NewMsgQueue()
	te1.Mq = &mq
	te2.Mq = &mq
	te2.Observe(&mq, TESTMESG, elementID(te1.ID), func(m *Message) { te2.receive(m) })
	te2.Running = true
	go mq.Loop()
	te1.send()

	for te2.Running {
	}

}
