// Package eos provides utilities for EOS - EOSIO.
package eos

import (
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/keng42/xwallet/pkg/utilities/hash"
)

// ConvertPub converts legacy public key to EOS public key
func ConvertPub(buf []byte) (pub string) {
	checksum := hash.Rmd160(buf)[:4]

	buf = append(buf, checksum...)
	pub = "EOS" + base58.Encode(buf)

	return
}

// ConvertPriv converts legacy private key to EOS private key
func ConvertPriv(buf []byte) (priv string) {
	buf = append([]byte{0x80}, buf...)

	sum1 := sha256.Sum256(buf)
	sum2 := sha256.Sum256(sum1[:])
	checksum := sum2[:4]

	buf = append(buf, checksum...)
	priv = base58.Encode(buf)

	return
}
