// Package nano provides utilities for NANO - Nano.
package nano

import (
	"crypto/ed25519"
	"encoding/base32"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/keng42/go-bip32-ed25519"
	"github.com/keng42/xwallet/pkg/chains/nano/edwards25519"
	"github.com/keng42/xwallet/pkg/utilities/hash"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/blake2b"
)

// nano uses a non-standard base32 character set.
var EncodingNano = base32.NewEncoding("13456789abcdefghijkmnopqrstuwxyz")

// NewAddress creates a new address, public key and private key by specify mnemonic and path.
func NewAddress(mne string, pwd string, path string, index uint32) (addr string, pub, priv []byte, err error) {
	seed := bip39.NewSeed(mne, pwd)

	dpath, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return
	}

	// NOTE: The reason of using "github.com/keng42/go-bip32-ed25519"
	// (fork from "github.com/tyler-smith/go-bip32")
	// is that the hashKey in bip32.NewMasterKey is different.

	key, err := bip32.NewMasterKey(seed)
	if err != nil {
		return
	}

	for _, p := range dpath {
		key, err = key.NewChildKey(p)
		if err != nil {
			return
		}
	}

	key, err = key.NewChildKey(index)
	if err != nil {
		return
	}

	priv = key.Key
	pub = ConvertPub(priv)
	addr, err = ConvertAddr(pub)
	return
}

// ConvertPub converts bip32 private key bytes to ed25519.PublicKey
func ConvertPub(buf []byte) (pub ed25519.PublicKey) {
	pub = make([]byte, ed25519.PublicKeySize)

	digest := blake2b.Sum512(buf)
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	// NOTE: We've forked golang's ed25519 implementation
	// to use blake2b instead of sha3
	// https://github.com/frankh/nano/blob/master/address/address.go

	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], digest[:])
	edwards25519.GeScalarMultBase(&A, &hBytes)
	var pubBuf [32]byte
	A.ToBytes(&pubBuf)

	copy(pub, pubBuf[:])
	return
}

// ConvertAddr converts ed25519.PublicKey to nano address.
func ConvertAddr(pub ed25519.PublicKey) (addr string, err error) {
	// calc checksum
	h, err := blake2b.New(5, nil)
	if err != nil {
		return
	}
	h.Write(pub)
	checksum := h.Sum(nil)
	checksum = hash.Reversed(checksum)

	// Pubkey is 256bits, base32 must be multiple of 5 bits
	// to encode properly.
	// Pad the start with 0's and strip them off after base32 encoding
	buf := append([]byte{0, 0, 0}, pub...)
	buf = append(buf, checksum...)
	address := EncodingNano.EncodeToString(buf)[4:]

	addr = string("nano_" + address)
	return
}
