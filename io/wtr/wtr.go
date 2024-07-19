package main

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

// DigitsFreq calculates leading digit frequency.
type DigitsFreq struct {
	Freqs map[rune]int
	inNum bool
}

// Write implements io.Writer.
func (d *DigitsFreq) Write(p []byte) (n int, err error) {
	if d.Freqs == nil {
		d.Freqs = make(map[rune]int)
	}

	for _, b := range p {
		if r := rune(b); unicode.IsDigit(r) {
			if !d.inNum {
				d.Freqs[r]++
				d.inNum = true
			}
			continue
		}

		// Not a digit.
		if d.inNum {
			d.inNum = false
		}
	}

	return len(p), nil
}

func main() {
	data := `
	We have 1234
	then 2342
	then 110
	then 37
	`

	var df DigitsFreq
	io.Copy(&df, strings.NewReader(data))

	for r, c := range df.Freqs {
		fmt.Printf("%c -> %d\n", r, c)
	}
}
