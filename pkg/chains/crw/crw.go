// Package crw provides utilities for CRW - Crown.
package crw

import (
	"crypto/sha256"
	"log"

	"github.com/eknkc/basex"
)

var (
	EncodingBTC *basex.Encoding
)

func init() {
	var err error

	EncodingBTC, err = basex.NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	if err != nil {
		log.Fatal(err)
	}
}

// ConvertAddr converts legacy btc address to CRW address
func ConvertAddr(s string) (addr string, err error) {
	b, err := EncodingBTC.Decode(s)
	if err != nil {
		return
	}

	hash160 := make([]byte, 23)
	hash160[0] = 0x01 //C
	hash160[1] = 0x75 //R
	hash160[2] = 0x07 //W
	copy(hash160[3:], b[1:21])

	_checksum := sha256.Sum256(hash160)
	_checksum = sha256.Sum256(_checksum[:])
	checksum := _checksum[:4]

	binaryAddr := append(hash160, checksum...)
	addr = EncodingBTC.Encode(binaryAddr)

	return
}
