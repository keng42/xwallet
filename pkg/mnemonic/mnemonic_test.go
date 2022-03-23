package mnemonic_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/keng42/xwallet/pkg/mnemonic"
	"github.com/stretchr/testify/require"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

var mne = "term mushroom resemble heavy calm tribe leader aim coyote polar during neglect"

func TestA(t *testing.T) {
	mne, err := mnemonic.FromEntropy("df723add353209d09fa02a31b4ed10ca")
	require.Nil(t, err)
	fmt.Println("===== mnemonic:", mne)

	seed := bip39.NewSeed(mne, "")
	fmt.Println("===== seed:", hex.EncodeToString(seed))

	masterKey, err := bip32.NewMasterKey(seed)
	require.Nil(t, err)

	fmt.Println("===== masterKey hex:", hex.EncodeToString(masterKey.Key))
	fmt.Println("===== masterKey str:", masterKey.String())

	bip32Key, err := masterKey.NewChildKey(0)
	require.Nil(t, err)

	fmt.Println("===== bip32Key hex:", hex.EncodeToString(bip32Key.Key))
	fmt.Println("===== bip32Key str:", bip32Key.String())
	fmt.Println("===== bip32Key pub:", bip32Key.PublicKey().String())

	key0, err := bip32Key.NewChildKey(0)
	require.Nil(t, err)

	fmt.Println("===== key0 hex:", hex.EncodeToString(key0.Key))
	fmt.Println("===== key0 str:", key0.String())
	fmt.Println("===== key0 pub:", key0.PublicKey().String())
	fmt.Println("===== key0 pub hex:", hex.EncodeToString(key0.PublicKey().Key))
	fmt.Println("===== key0 priv base58", bip32.BitcoinBase58Encoding.EncodeToString(key0.Key))

	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), key0.Key)
	wif, err := btcutil.NewWIF(privKey, &chaincfg.MainNetParams, true)
	require.Nil(t, err)

	fmt.Println("===== wif:", wif.String())

	addrPubKey, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.MainNetParams)
	require.Nil(t, err)

	fmt.Println("===== addr:", addrPubKey.EncodeAddress())
}
