// Package xwc provides utilities for XWC - Whitecoin.
package xwc

import (
	"crypto/sha512"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/keng42/xwallet/pkg/utilities/hash"
)

// ConvertAddr converts legacy public key to XWC address
func ConvertAddr(buf []byte) (addr string) {
	sum := sha512.Sum512(buf)
	addrBuf := hash.Rmd160(sum[:])
	addrBuf = append([]byte{0x35}, addrBuf...)

	checksum := hash.Rmd160(addrBuf)[:4]
	addrBuf = append(addrBuf, checksum...)

	addr = "XWC" + base58.Encode(addrBuf)

	return
}

// ConvertPub converts legacy public key to XWC public key
func ConvertPub(buf []byte) (pub string) {
	checksum := hash.Rmd160(buf)[:4]

	pubBuf := make([]byte, len(buf))
	copy(pubBuf, buf)
	pubBuf = append(pubBuf, checksum...)

	pub = "XWC" + base58.Encode(pubBuf)

	return
}

// ConvertPriv converts legacy private key to XWC private key
func ConvertPriv(buf []byte) (priv string) {
	privBuf := []byte{0x80}
	privBuf = append(privBuf, buf...)

	checksum := hash.DoubleSha256(privBuf)[:4]

	privBuf = append(privBuf, checksum...)

	priv = base58.Encode(privBuf)

	return
}
