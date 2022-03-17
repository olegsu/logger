package logger

import (
	"fmt"
	"time"
)

type Logger struct {
	msg chan string
}

type Stringer interface {
	String() string
}

func New() Logger {
	l := Logger{
		msg: make(chan string),
	}
	go l.process()
	return l
}

func (l *Logger) Info(msg string, kvs ...Stringer) {
	if len(kvs)%2 != 0 {
		panic("not good")
	}
	l.msg <- msg
}

func (l *Logger) process() {
	msgs := []string{}
	for {
		select {
		case m := <-l.msg:
			msgs = append(msgs, m)
		case <-time.NewTicker(5 * time.Second).C:
			for _, msg := range msgs {
				fmt.Println(msg)
			}
		}
	}
}
