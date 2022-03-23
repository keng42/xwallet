package hash

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

// Rmd160 calculates the hash ripemd160(b).
func Rmd160(buf []byte) []byte {
	h := ripemd160.New()
	h.Write(buf)
	return h.Sum(nil)
}

// DoubleSha256 calculates the hash sha256(sha256(b)).
func DoubleSha256(buf []byte) []byte {
	b := sha256.Sum256(buf)
	b = sha256.Sum256(b[:])
	return b[:]
}

func Reversed(buf []byte) []byte {
	b := make([]byte, len(buf))
	l := len(b) - 1
	for i, v := range buf {
		b[l-i] = v
	}

	return b
}
