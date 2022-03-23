package xlm_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	bip32ed25519 "github.com/keng42/go-bip32-ed25519"
	"github.com/keng42/xwallet/pkg/chains/xlm"
	"github.com/stellar/go/exp/crypto/derivation"
	"github.com/stellar/go/keypair"
	"github.com/stretchr/testify/require"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

var mne = "term mushroom resemble heavy calm tribe leader aim coyote polar during neglect"

func TestA(t *testing.T) {
	kp, err := xlm.NewAddress(mne, "", "m/44'/148'", bip32.FirstHardenedChild+0)
	require.Nil(t, err)
	require.Equal(t, "GB3P6OD356E3WAKHZ2N5A6PHXF2EY6NR7K6JAQJ52K2QB4FFRUI4IHME", kp.Address())
	require.Equal(t, "GB3P6OD356E3WAKHZ2N5A6PHXF2EY6NR7K6JAQJ52K2QB4FFRUI4IHME", kp.Address())
	require.Equal(t, "SBX5QYW7G4C33KTLH3S3IIOISPSK6YR2JO3ZS73LGLMNMPV7TRQNGWBB", kp.Seed())
}

func TestB(t *testing.T) {
	seed := bip39.NewSeed(mne, "")

	// derive key with 'xlm derivation'
	dkey, err := derivation.DeriveForPath("m/44'/148'/0'", seed)
	require.Nil(t, err)

	// derive key with bip32-ed25519
	bkey, err := bip32.NewMasterKey(seed)
	require.Nil(t, err)
	bkey, err = bkey.NewChildKey(44 + bip32.FirstHardenedChild)
	require.Nil(t, err)
	bkey, err = bkey.NewChildKey(148 + bip32.FirstHardenedChild)
	require.Nil(t, err)
	bkey, err = bkey.NewChildKey(0 + bip32.FirstHardenedChild)
	require.Nil(t, err)

	// derive key with bip32-ed25519
	ekey, err := bip32ed25519.NewMasterKey(seed)
	require.Nil(t, err)
	ekey, err = ekey.NewChildKey(44 + bip32.FirstHardenedChild)
	require.Nil(t, err)
	ekey, err = ekey.NewChildKey(148 + bip32.FirstHardenedChild)
	require.Nil(t, err)
	ekey, err = ekey.NewChildKey(0 + bip32.FirstHardenedChild)
	require.Nil(t, err)

	// difference between dkey and bkey
	rawSeed := dkey.RawSeed()
	dSeed := hex.EncodeToString(rawSeed[:])
	dPriv := hex.EncodeToString(dkey.Key)
	bPriv := hex.EncodeToString(bkey.Key)
	ePriv := hex.EncodeToString(ekey.Key)

	fmt.Println("=====     xlm seed", dSeed)
	fmt.Println("=====     xlm priv", dPriv)
	fmt.Println("=====     btc priv", bPriv)
	fmt.Println("===== ed25519 priv", ePriv)

	require.Equal(t, "6fd862df3705bdaa6b3ee5b421c893e4af623a4bb7997f6b32d8d63ebf9c60d3", dSeed, "xlm seed")
	require.Equal(t, "6fd862df3705bdaa6b3ee5b421c893e4af623a4bb7997f6b32d8d63ebf9c60d3", dPriv, "xlm priv")
	require.Equal(t, "c5efadb1ea80f9e64ed0ae1e8786d2d5de3f89f3dd7abb1f174f8a3eb6a666a6", bPriv, "btc priv")
	require.Equal(t, "6fd862df3705bdaa6b3ee5b421c893e4af623a4bb7997f6b32d8d63ebf9c60d3", ePriv, "ed25519 priv")
	// NOTE: The difference is that the hashKey is different in NewMasterKey.
	// They are 'Bitcoin seed' and 'ed25519 seed' respectively.

	kp, err := keypair.FromRawSeed(rawSeed)
	require.Nil(t, err)
	fmt.Println("===== kp addr", kp.Address())
	fmt.Println("===== kp seed", kp.Seed())
}
