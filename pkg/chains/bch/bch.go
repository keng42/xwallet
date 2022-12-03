// Package bch provides utilities for BCH - Bitcoin Cash.
package bch

import (
	"errors"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"

	// TODO code review
	"github.com/keng42/bchutil"
)

type AddrStyle int

const (
	ADDR_STYLE_LEGACY AddrStyle = iota
	ADDR_STYLE_CASH
	ADDR_STYLE_BITPAY
)

func init() {
	bchutil.Prefixes["bch-mainnet"] = "bitcoincash"
	bchutil.Prefixes["bch-testnet"] = "bchtest"
	bchutil.Prefixes["bch-regtest"] = "bchreg"
	bchutil.Prefixes["slp-mainnet"] = "simpleledger"
	bchutil.Prefixes["slp-testnet"] = "slptest"
	bchutil.Prefixes["slp-regtest"] = "slpreg"
}

// ConvertToCashAddr converts legacy public key to cash address
func ConvertToCashAddr(buf []byte, net *chaincfg.Params) (addr string, err error) {
	return convertToCashAddr("bch", buf, net)
}

// ConvertToSlpCashAddr converts legacy public key to slp cash address
func ConvertToSlpCashAddr(buf []byte, net *chaincfg.Params) (addr string, err error) {
	return convertToCashAddr("slp", buf, net)
}

// convertToCashAddr converts legacy public key to cash address
func convertToCashAddr(chain string, buf []byte, net *chaincfg.Params) (addr string, err error) {
	if net == nil {
		net = &chaincfg.MainNetParams
	}

	name := chain + "-" + net.Name
	prefix, ok := bchutil.Prefixes[name]
	if !ok {
		err = errors.New("unknown network parameters")
		return
	}

	addressPKH, err := bchutil.NewCashAddressPubKeyHash(btcutil.Hash160(buf), &chaincfg.Params{Name: name})
	if err != nil {
		return
	}

	addr = prefix + ":" + addressPKH.EncodeAddress()

	return
}

// ConvertToBitpayAddr converts legacy public key to bitpay-style address
func ConvertToBitpayAddr(buf []byte, net *chaincfg.Params) (addr string, err error) {
	if net == nil {
		net = &chaincfg.MainNetParams
	}

	addressPKH, err := bchutil.NewBitpayAddressPubKeyHash(btcutil.Hash160(buf), net)
	if err != nil {
		return
	}

	addr = addressPKH.EncodeAddress()

	return
}
