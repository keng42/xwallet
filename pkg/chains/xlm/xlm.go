// Package xlm provides utilities for XLM - Stellar
package xlm

import (
	"github.com/stellar/go/exp/crypto/derivation"
	"github.com/stellar/go/keypair"
	"github.com/tyler-smith/go-bip39"
)

// NewAddress creates a new address, public key and private key by specify mnemonic and path.
func NewAddress(mne string, pwd string, path string, index uint32) (kp *keypair.Full, err error) {
	seed := bip39.NewSeed(mne, pwd)

	key, err := derivation.DeriveForPath(path, seed)
	if err != nil {
		return
	}

	key, err = key.Derive(index)
	if err != nil {
		return
	}

	kp, err = keypair.FromRawSeed(key.RawSeed())

	return
}
