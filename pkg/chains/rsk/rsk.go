// Package rsk provides utilities for R-BTC - RSK
package rsk

import (
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

// ConvertAddr converts legacy eth address to rsk address
func ConvertAddr(s string, chainID int) (addr string) {
	// verify length and prefix
	s = strings.ToLower(s)
	if !common.IsHexAddress(s) {
		return
	}
	if strings.HasPrefix(s, "0x") {
		s = s[2:]
	}

	// compute checksum
	prefix := ""
	if chainID > 0 {
		prefix = strconv.Itoa(chainID) + "0x"
	}

	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(prefix + s))
	hash := sha.Sum(nil)

	var c byte
	buf := []byte("0x" + s)

	for i := 2; i < len(buf); i++ {
		c = buf[i]
		if !(('0' <= c && c <= '9') || ('a' <= c && c <= 'f')) {
			return
		}

		hashByte := hash[(i-2)/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if c > '9' && hashByte > 7 {
			buf[i] -= 32
		}
	}

	addr = string(buf)

	return
}
