// Package hns provides utilities for HNS - Handshake.
package hns

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"golang.org/x/crypto/blake2b"
)

// ConvertPub converts legacy public key to HNS address
func ConvertAddr(buf []byte) (addr string, err error) {
	// hns wif key id
	// https://github.com/handshake-org/hsd/blob/master/lib/protocol/networks.js

	h, err := blake2b.New(20, nil)
	h.Write(buf)
	witnessProg := h.Sum(nil)

	// TODO support testnet
	addressWPKH, err := btcutil.NewAddressWitnessPubKeyHash(witnessProg, &chaincfg.Params{Bech32HRPSegwit: "hs"})
	if err != nil {
		return
	}

	addr = addressWPKH.EncodeAddress()

	return
}
