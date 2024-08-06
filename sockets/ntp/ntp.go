package main

import (
	"encoding/binary"
	"net"
	"time"
)

// NTPMessage is an NTP message.
type NTPMessage struct {
	VNMode           uint8
	Stratum          uint8
	Poll             uint8
	Precision        uint8
	RootDelay        uint32
	RootDispersion   uint32
	RedID            uint32
	RefTimeSec       uint32
	RefTimeFrac      uint32
	OrigTimeSec      uint32
	OrigTimeFrac     uint32
	ReceivedTimeSec  uint32
	ReceivedTimeFrac uint32
	TransmitTimeSec  uint32
	TransmitTimeFrac uint32
}

func ntpDelta() time.Duration {
	unixEpoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	ntpEpoch := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	return ntpEpoch.Sub(unixEpoch)
}

var (
	// NTPDelta is the difference between NTP and Unix time.
	NTPDelta = ntpDelta()
	zeroTime time.Time
)

// TransmitTime returns the transmit time.
func (m *NTPMessage) TransmitTime() time.Time {
	sec := int64(m.TransmitTimeSec)
	nanos := (int64(m.TransmitTimeFrac) * 1e9) >> 32
	return time.Unix(sec, nanos).Add(NTPDelta)
}

// CurrentTime returns the current time from NTP host.
func CurrentTime(host string) (time.Time, error) {
	conn, err := net.Dial("udp", host)
	if err != nil {
		return zeroTime, err
	}

	defer conn.Close()

	msg := NTPMessage{
		VNMode: 0b00011011, // li = 0, vn = 3, and mode = 3
	}

	if err := binary.Write(conn, binary.BigEndian, &msg); err != nil {
		return zeroTime, err
	}

	if err := binary.Read(conn, binary.BigEndian, &msg); err != nil {
		return zeroTime, err
	}

	return msg.TransmitTime(), nil
}

func main() {

}
