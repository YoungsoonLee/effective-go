package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Ride is a ride record.
type Ride struct {
	ID       int
	Time     time.Time
	Duration time.Duration
	Distance float64
	Price    float64
}

// UnmarshalRide returns a Ride from serialized data.
func UnmarshalRide(data []byte, ride *Ride) error {
	r := bytes.NewReader(data)
	return NewDecoder(r).DecodeRide(ride)
}

func main() {
	data := []byte(`{
        "id": 83,
        "time": "2020-03-01T14:32:05Z",
        "duration": 900000000000,
        "distance": 2.7,
        "price": 22.3
    }`)

	var r Ride
	if err := UnmarshalRide(data, &r); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%+v\n", r)
	}

}

// Decoder is an example decoder
type Decoder struct {
	dec *json.Decoder
}

// NewDecoder returns a new Decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{json.NewDecoder(r)}
}

// DecodeRide decodes a Ride.
func (d *Decoder) DecodeRide(r *Ride) error {
	return d.dec.Decode(r)
}
