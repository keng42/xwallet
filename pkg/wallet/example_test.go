package wallet_test

import (
	"fmt"
	"log"

	"github.com/keng42/xwallet/pkg/chains/bch"
	"github.com/keng42/xwallet/pkg/chains/btc"
	"github.com/keng42/xwallet/pkg/networks"
	"github.com/keng42/xwallet/pkg/wallet"
)

func ExampleWallet_NewAddress_legacy() {
	nw, err := networks.New("BTC - Bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	acc, err := w.NewAddress(nw, "", 0, wallet.SEGWIT_TYPE_DISABLE)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("address: %s\n", acc.Address)
	fmt.Printf("public key: %s\n", acc.PublicKeyText)
	fmt.Printf("private key: %s\n", acc.PrivateKeyText)

	// Output:
	// address: 1NaFxCzxCJyYK6fP4kn936bSbpxvNTaBe6
	// public key: 02fb207bd0e7a3dff762a91a70d54fa65674b74aa6bc1e302deffc7bce5163643c
	// private key: L4t3tg2jB5XvDJ7s3vsprkHX6xdQUPcxkSeYyvAytLjqZhruRRaQ
}

func ExampleWallet_NewAddress_segwit() {
	nw, err := networks.New("BTC - Bitcoin")
	if err != nil {
		log.Fatal(err)
	}

	acc, err := w.NewAddress(nw, "", 0, wallet.SEGWIT_TYPE_P2WSH)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("address: %s\n", acc.Address)
	fmt.Printf("public key: %s\n", acc.PublicKeyText)
	fmt.Printf("private key: %s\n", acc.PrivateKeyText)

	// Output:
	// address: bc1qpypdshwmyydmnc0hapgam2alj7yms7z6m35fqr27cfzkme3vucnq9v9t08
	// public key: 02fb207bd0e7a3dff762a91a70d54fa65674b74aa6bc1e302deffc7bce5163643c
	// private key: L4t3tg2jB5XvDJ7s3vsprkHX6xdQUPcxkSeYyvAytLjqZhruRRaQ
}

func ExampleWallet_NewAddress_bitpay() {
	nw, err := networks.New("BCH - Bitcoin Cash")
	if err != nil {
		log.Fatal(err)
	}

	acc, err := w.NewAddress(nw, "", 0, 0)
	if err != nil {
		log.Fatal(err)
	}

	cashAddr := acc.Address
	acc.ToBitcoinCash(&btc.MainNetParams, bch.ADDR_STYLE_BITPAY)
	bitpayAddr := acc.Address

	fmt.Printf("cash address: %s\n", cashAddr)
	fmt.Printf("bitpay address: %s\n", bitpayAddr)
	fmt.Printf("public key: %s\n", acc.PublicKeyText)
	fmt.Printf("private key: %s\n", acc.PrivateKeyText)

	// Output:
	// cash address: bitcoincash:qz0r7yu7ns40xm0yda8dlftc445a5km57gnqthn900
	// bitpay address: CWtctzVwTbBJPCnyrV9W35JXPSdBJPJ7Gd
	// public key: 022bb58fb4028ff449e9af7c1e9da8c16577b0d54bcdd32f90e8684fc8cc8757e0
	// private key: L2ZeRGjztRuYRaKYPQJxSKMKivfig8UKZbsr6Kb98HykZCQbjjCw
}
