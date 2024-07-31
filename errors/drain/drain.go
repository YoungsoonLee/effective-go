package main

import (
	"fmt"
	"log"
	"time"
)

// Message is a message.
type Message struct {
	Time time.Time
	Type string
	Data []byte
}

func drain(ch <-chan Message, handler func(Message)) {
	for msg := range ch {
		msg.Time = time.Now()
		//go handler(msg)
		safelyGo(func() {
			handler(msg)
		})
	}
}

func testHandler(msg Message) {
	ts := msg.Time.Format("15:04:05")
	log.Printf("%s [%s] %x", ts, msg.Type, msg.Data[:20])
}

func main() {
	ch := make(chan Message)

	// Populate the channel with some messages.
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			msg := Message{
				Type: "test",
				Data: []byte(fmt.Sprintf("message %d", i)),
			}
			ch <- msg
		}
	}()

	drain(ch, testHandler)

	time.Sleep(time.Second)
	fmt.Println("done")
}

// safelyGo will run the given function in a goroutine and recover from any panics.
func safelyGo(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v", r)
			}
		}()
		f()
	}()
}
