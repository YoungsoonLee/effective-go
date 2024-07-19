package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"time"
)

// Kind is the type of event.
type Kind string

const (
	// AddKind is the kind of event when a file is added.
	AddKind Kind = "add"
	// CheckoutKind is the kind of event when a file is checked out.
	CheckoutKind Kind = "checkout"
)

// Event is an event that occurs in the system.
type Event interface {
	Kind() Kind
}

// Add is an event when a file is added.
type Add struct {
	Time time.Time
	ID   string
	User string
	Item int // SKU
}

// Kind returns the kind of event.
func (*Add) Kind() Kind {
	return AddKind
}

// Checkout is an event when a file is checked out.
type Checkout struct {
	Time time.Time
	Cart string
	User string
}

// Kind returns the kind of event.
func (*Checkout) Kind() Kind {
	return CheckoutKind
}

func init() {
	gob.Register(&Add{})
	gob.Register(&Checkout{})
}

// Encoder encodes events.
type Encoder struct {
	enc *gob.Encoder
}

// NewEncoder creates a new encoder.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{gob.NewEncoder(w)}
}

// Encode encodes an event.
func (e *Encoder) Encode(ev Event) error {
	return e.enc.Encode(&ev)
}

func eventHandler(r io.Reader) error {
	dec := gob.NewDecoder(r)
	for {
		var ev Event
		if err := dec.Decode(&ev); err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		switch ev.Kind() {
		case AddKind:
			add := ev.(*Add)
			// handle add event
			handleAdd(add)
		case CheckoutKind:
			checkout := ev.(*Checkout)
			// handle checkout event
			handleCheckout(checkout)
		default:
			return fmt.Errorf("unknown event kind: %q", ev.Kind())
		}
	}
}

func handleAdd(a *Add) {
	fmt.Printf("add: %+v\n", a)
}

func handleCheckout(c *Checkout) {
	fmt.Printf("checkout: %+v\n", c)
}
