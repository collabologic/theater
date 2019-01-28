package foundation

import (
	"fmt"
	"testing"
	"time"
)

type TestElement3 struct {
	TheElement
	TheObserver
	TheNotifier
}

func (te *TestElement3) send() {
	nte := Notifier(te)
	te.Notify(Msg{
		TESTMESG,
		te.ID,
		nte,
		nil,
	})
}

type TestElement4 struct {
	TheElement
	TheObserver
	TheNotifier
	Running bool
}

func (te *TestElement4) receive(m *Msg) {
	fmt.Println("RECEIVE!")
	fmt.Println(m)
	te.Running = false
}

const (
	TESTMESG2 MsgIdentifier = "TEST"
)

type TestElement5 struct {
	TheElement
	TheObserver
	TheNotifier
	c int32
}

func (te *TestElement5) Behave() {
	fmt.Printf("tick!:%d\r\n", te.c)
	te.c += 1
}

var te3 TestElement3
var te4 TestElement4
var te5 TestElement5

func Test_Pulsar_1(t *testing.T) {
	mq := NewMsgQueue()
	te3.Mq = &mq
	te4.Mq = &mq
	p := NewPulsar(50)
	p.RegistBehavior(&te5)
	te4.Observe(&mq, TESTMESG, elementID(te3.ID), func(m *Msg) { te4.receive(m) })
	te4.Running = true
	go mq.Loop()
	go p.Loop()
	te3.send()

	time.Sleep(1000 * time.Microsecond)
	te4.Running = false

}
