// Package swtc provides utilities for SWTC - Jingtum.
package swtc

import (
	"encoding/hex"
	"log"

	"github.com/eknkc/basex"
)

var (
	EncodingSWTC *basex.Encoding
	EncodingBTC  *basex.Encoding
)

func init() {
	var err error
	EncodingSWTC, err = basex.NewEncoding("jpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65rkm8oFqi1tuvAxyz")
	if err != nil {
		log.Fatal(err)
	}

	EncodingBTC, err = basex.NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	if err != nil {
		log.Fatal(err)
	}
}

// ConvertAddr converts legacy btc address to swtc address
func ConvertAddr(s string) (addr string, err error) {
	var b []byte
	b, err = EncodingBTC.Decode(s)
	if err != nil {
		return
	}
	addr = EncodingSWTC.Encode(b)

	return
}

// ConvertPriv converts wif private key to swtc private key
func ConvertPriv(s string) (priv string, err error) {
	var b []byte
	b, err = EncodingBTC.Decode(s)
	if err != nil {
		return
	}
	priv = hex.EncodeToString(b)[2:66]

	return
}
