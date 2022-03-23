// Package xrp provides utilities for XRP - Ripple.
package xrp

import (
	"encoding/hex"
	"log"

	// TODO code review
	"github.com/eknkc/basex"
)

var (
	EncodingXRP *basex.Encoding
	EncodingBTC *basex.Encoding
)

func init() {
	var err error
	EncodingXRP, err = basex.NewEncoding("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
	if err != nil {
		log.Fatal(err)
	}

	EncodingBTC, err = basex.NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
	if err != nil {
		log.Fatal(err)
	}
}

// ConvertAddr converts legacy btc address to xrp address
func ConvertAddr(s string) (addr string, err error) {
	var b []byte
	b, err = EncodingBTC.Decode(s)
	if err != nil {
		return
	}
	addr = EncodingXRP.Encode(b)

	return
}

// ConvertPriv converts wif private key to xrp private key
func ConvertPriv(s string) (priv string, err error) {
	var b []byte
	b, err = EncodingBTC.Decode(s)
	if err != nil {
		return
	}
	priv = hex.EncodeToString(b)[2:66]

	return
}
