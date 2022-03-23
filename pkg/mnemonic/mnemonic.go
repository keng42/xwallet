// Package mnemonic provides functions to generate mnemonic phrases.
package mnemonic

import (
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

var (
	// ErrWordsNumberInvalid is returned when trying to generate mnemonic with invalid words number.
	ErrWordsNumberInvalid = errors.New("mnemonic words number must be [12, 24] and a multiple of 3")

	ErrInvalidMnemonic = errors.New("invalid mnemonic")
)

// Generate will return a mnemonic string constains [wordsNum] words using random entropy.
func Generate(wordsNum int) (string, error) {
	if wordsNum%3 != 0 || wordsNum < 12 || wordsNum > 24 {
		return "", ErrWordsNumberInvalid
	}

	strength := wordsNum / 3 * 32
	entropy, err := bip39.NewEntropy(strength)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(entropy)
}

// FromEntropy will return a mnemonic string for the given hex-encoded entropy.
func FromEntropy(entropy string) (string, error) {
	buf, err := hex.DecodeString(entropy)
	if err != nil {
		return "", err
	}

	return bip39.NewMnemonic(buf)
}

// GetBip32ExtendedKey derives bip32.Key with specify mnemonic and path.
func GetBip32ExtendedKey(mne, pwd, path string) (key *bip32.Key, err error) {
	if !bip39.IsMnemonicValid(mne) {
		err = ErrInvalidMnemonic
		return
	}

	dpath, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return
	}

	seed := bip39.NewSeed(mne, pwd)

	key, err = bip32.NewMasterKey(seed)
	if err != nil {
		return
	}

	for _, p := range dpath {
		key, err = key.NewChildKey(p)
		if err != nil {
			return
		}
	}

	return
}
