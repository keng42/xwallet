// Package atom provides utilities for ATOM - Cosmos.
package atom

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcutil/bech32"
)

// ConvertAddr converts legacy public key to ATOM address
func ConvertAddr(buf []byte, hrp string) (addr string, err error) {
	var b []byte
	b, err = bech32.ConvertBits(btcutil.Hash160(buf), 8, 5, true)
	if err != nil {
		return
	}

	addr, err = bech32.Encode(hrp, b)

	return
}

// ConvertPub converts legacy public key to ATOM public key
func ConvertPub(buf []byte, hrp string) (pub string, err error) {
	// AminoSecp256k1 public key prefix and length
	buf = append([]byte{0xeb, 0x5a, 0xe9, 0x87, 0x21}, buf...)

	var b []byte
	b, err = bech32.ConvertBits(buf, 8, 5, true)
	if err != nil {
		return
	}

	pub, err = bech32.Encode(hrp+"pub", b)

	return
}
