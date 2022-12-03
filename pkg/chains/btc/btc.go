// Package btc provides utilities for BTC - Bitcoin.
package btc

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/base58"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
)

// support multibyte net.PubKeyHashAddrID and net.ScriptHashAddrID
// https://github.com/iancoleman/bip39/issues/94

type NetworkParams struct {
	Raw *chaincfg.Params

	PubKeyHashAddrID []byte
	ScriptHashAddrID []byte
}

var MainNetParams NetworkParams

func init() {
	MainNetParams = NewNetworkParams(nil, 0, 0)
}

func NewNetworkParams(raw *chaincfg.Params, pubKeyHashAddrID int64, scriptHashAddrID int64) (net NetworkParams) {
	net.Raw = raw
	if net.Raw == nil {
		net.Raw = &chaincfg.MainNetParams
	}
	if pubKeyHashAddrID == 0 {
		pubKeyHashAddrID = int64(net.Raw.PubKeyHashAddrID)
	}
	if scriptHashAddrID == 0 {
		scriptHashAddrID = int64(net.Raw.ScriptHashAddrID)
	}

	var buf []byte

	if pubKeyHashAddrID <= 255 {
		buf = []byte{byte(pubKeyHashAddrID)}
	} else {
		buf = make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(pubKeyHashAddrID))
	}
	net.PubKeyHashAddrID = buf

	if scriptHashAddrID <= 255 {
		buf = []byte{byte(scriptHashAddrID)}
	} else {
		buf = make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(scriptHashAddrID))
	}
	net.ScriptHashAddrID = buf

	return
}

func (np *NetworkParams) GetRaw() *chaincfg.Params {
	if np.Raw == nil {
		return &chaincfg.MainNetParams
	}
	return np.Raw
}

// P2PKH pay to public key hash
func P2PKH(buf []byte, net *NetworkParams) (addr string, err error) {
	if net == nil {
		net = &MainNetParams
	}

	addr = CheckEncode(btcutil.Hash160(buf), net.PubKeyHashAddrID)

	return
}

// P2SH pay to script hash
func P2SH(buf []byte, net *NetworkParams) (addr string, err error) {
	if net == nil {
		net = &MainNetParams
	}

	serializedScript, err := txscript.NewScriptBuilder().
		AddOp(txscript.OP_0).
		AddData(btcutil.Hash160(buf)).
		Script()
	if err != nil {
		return
	}

	addr = CheckEncode(btcutil.Hash160(serializedScript), net.ScriptHashAddrID)

	return
}

// P2WPKH pay to witness public key hash
func P2WPKH(buf []byte, net *NetworkParams) (addr string, err error) {
	if net == nil {
		net = &MainNetParams
	}

	addressWPKH, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(buf), net.GetRaw())
	if err != nil {
		return
	}

	addr = addressWPKH.EncodeAddress()

	return
}

// P2WPKHInP2SH pay to witness public key hash nested in pay to script hash
func P2WPKHInP2SH(buf []byte, net *NetworkParams) (addr string, err error) {
	if net == nil {
		net = &MainNetParams
	}

	serializedScriptWPK, err := txscript.NewScriptBuilder().
		AddOp(txscript.OP_0).
		AddData(btcutil.Hash160(buf)).
		Script()
	if err != nil {
		return
	}

	addr = CheckEncode(btcutil.Hash160(serializedScriptWPK), net.ScriptHashAddrID)

	return
}

// P2WSH pay to witness script hash
func P2WSH(buf []byte, net *NetworkParams) (addr string, err error) {
	if net == nil {
		net = &MainNetParams
	}

	addressPK, err := btcutil.NewAddressPubKey(buf, net.GetRaw())
	if err != nil {
		return
	}

	witnessScript, err := txscript.MultiSigScript([]*btcutil.AddressPubKey{addressPK}, 1)
	if err != nil {
		return
	}

	h256 := sha256.Sum256(witnessScript)
	witnessProg := h256[:]

	addressWSH, err := btcutil.NewAddressWitnessScriptHash(witnessProg, net.GetRaw())
	if err != nil {
		return
	}

	addr = addressWSH.EncodeAddress()

	return
}

// P2WSHInP2SH pay to witness script hash nested in pay to script hash
func P2WSHInP2SH(buf []byte, net *NetworkParams) (addr string, err error) {
	if net == nil {
		net = &MainNetParams
	}

	addressPK, err := btcutil.NewAddressPubKey(buf, net.GetRaw())
	if err != nil {
		return
	}

	witnessScript, err := txscript.MultiSigScript([]*btcutil.AddressPubKey{addressPK}, 1)
	if err != nil {
		return
	}

	h256 := sha256.Sum256(witnessScript)
	scripthash := h256[:]

	redeemScript, err := txscript.NewScriptBuilder().
		AddOp(txscript.OP_0).
		AddData(scripthash).
		Script()

	addr = CheckEncode(btcutil.Hash160(redeemScript), net.ScriptHashAddrID)

	return
}

// CheckEncode prepends version bytes and appends a four byte checksum.
func CheckEncode(input []byte, version []byte) string {
	b := make([]byte, 0, len(version)+len(input)+4)
	b = append(b, version...)
	b = append(b, input[:]...)

	var cksum [4]byte
	h := sha256.Sum256(b)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])

	b = append(b, cksum[:]...)

	return base58.Encode(b)
}
