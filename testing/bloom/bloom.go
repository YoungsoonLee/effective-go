package bloom

import "testing"

// NumBytes returns the number of bytes that can hold n bits.
func NumBytes(nBits int) int {
	return (nBits + 7) / 8
}

// FuzzNumBytes is a fuzz test for NumBytes.
func FuzzNumBytes(f *testing.F) {
	f.Add(0)

	fn := func(t *testing.T, nBits int) {
		// ignore negative values
		if nBits < 0 {
			return
		}

		nBytes := NumBytes(nBits)
		ok := (nBytes*8 >= nBits) && (nBits >= (nBytes-1)*8)
		if !ok {
			t.Fatalf("NumBytes(%d) = %d", nBits, nBytes)
		}
	}

	f.Fuzz(fn)
}
