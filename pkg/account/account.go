package account

import (
	"crypto/ecdsa"
	"encoding/base64"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/keng42/xwallet/pkg/chains/atom"
	"github.com/keng42/xwallet/pkg/chains/bch"
	"github.com/keng42/xwallet/pkg/chains/btc"
	"github.com/keng42/xwallet/pkg/chains/crw"
	"github.com/keng42/xwallet/pkg/chains/csc"
	"github.com/keng42/xwallet/pkg/chains/eos"
	"github.com/keng42/xwallet/pkg/chains/fio"
	"github.com/keng42/xwallet/pkg/chains/hns"
	"github.com/keng42/xwallet/pkg/chains/rsk"
	"github.com/keng42/xwallet/pkg/chains/swtc"
	"github.com/keng42/xwallet/pkg/chains/trx"
	"github.com/keng42/xwallet/pkg/chains/xrp"
	"github.com/keng42/xwallet/pkg/chains/xwc"
	"github.com/stellar/go/keypair"
)

type Account struct {
	Path string

	Address        string
	PublicKey      []byte
	PrivateKey     []byte
	PublicKeyText  string
	PrivateKeyText string

	PublicKeyHex  string
	PrivateKeyHex string
	PrivateKeyWIF string

	EcdsaPub  *ecdsa.PublicKey
	EcdsaPriv *ecdsa.PrivateKey
}

func (acc *Account) ToEthereum() {
	address := crypto.PubkeyToAddress(*acc.EcdsaPub)
	acc.Address = address.Hex()
	acc.PublicKeyText = "0x" + acc.PublicKeyHex
	acc.PrivateKeyText = "0x" + acc.PrivateKeyHex
}

func (acc *Account) ToTron() (err error) {
	address := crypto.PubkeyToAddress(*acc.EcdsaPub)
	acc.Address, err = trx.ConvertAddr(address.Hex())
	if err != nil {
		return
	}

	acc.PublicKeyText = acc.PublicKeyHex
	acc.PrivateKeyText = acc.PrivateKeyHex
	return
}

func (acc *Account) ToRsk(chainID int) {
	acc.Address = rsk.ConvertAddr(crypto.PubkeyToAddress(*acc.EcdsaPub).Hex(), chainID)
	acc.PublicKeyText = "0x" + acc.PublicKeyHex
	acc.PrivateKeyText = "0x" + acc.PrivateKeyHex
}

func (acc *Account) ToHandshake() (err error) {
	acc.Address, err = hns.ConvertAddr(acc.PublicKey)
	return
}

func (acc *Account) ToStellar(kp *keypair.Full) (err error) {
	acc.Address = kp.Address()
	acc.PublicKeyText = kp.Address()
	acc.PrivateKeyText = kp.Seed()
	return
}

func (acc *Account) ToRipple() (err error) {
	acc.Address, err = xrp.ConvertAddr(acc.Address)
	if err != nil {
		return
	}

	acc.PrivateKeyText, err = xrp.ConvertPriv(acc.PrivateKeyText)
	return
}

func (acc *Account) ToJingtum() (err error) {
	acc.Address, err = swtc.ConvertAddr(acc.Address)
	if err != nil {
		return
	}

	acc.PrivateKeyText, err = swtc.ConvertPriv(acc.PrivateKeyText)
	return
}

func (acc *Account) ToCasinoCoin() (err error) {
	acc.Address, err = csc.ConvertAddr(acc.Address)
	if err != nil {
		return
	}

	acc.PrivateKeyText, err = csc.ConvertPriv(acc.PrivateKeyText)
	return
}

func (acc *Account) ToBitcoinCash(net *btc.NetworkParams, style bch.AddrStyle) (err error) {
	switch style {
	case bch.ADDR_STYLE_CASH:
		acc.Address, err = bch.ConvertToCashAddr(acc.PublicKey, net.GetRaw())
		return
	case bch.ADDR_STYLE_BITPAY:
		acc.Address, err = bch.ConvertToBitpayAddr(acc.PublicKey, net.GetRaw())
		return
	default:
		acc.Address, err = btc.P2PKH(acc.PublicKey, net)
		return
	}
}

func (acc *Account) ToSimpleLedger(net *btc.NetworkParams) (err error) {
	acc.Address, err = bch.ConvertToSlpCashAddr(acc.PublicKey, net.GetRaw())
	return
}

func (acc *Account) ToCrown() (err error) {
	acc.Address, err = crw.ConvertAddr(acc.Address)
	return
}

func (acc *Account) ToEOSIO() (err error) {
	acc.Address = ""
	acc.PublicKeyText = eos.ConvertPub(acc.PublicKey)
	acc.PrivateKeyText = eos.ConvertPriv(acc.PrivateKey)
	return
}

func (acc *Account) ToFIO() (err error) {
	acc.Address = ""
	acc.PublicKeyText = fio.ConvertPub(acc.PublicKey)
	acc.PrivateKeyText = fio.ConvertPriv(acc.PrivateKey)
	return
}

func (acc *Account) ToCosmos() (err error) {
	hrp := "cosmos"
	acc.Address, err = atom.ConvertAddr(acc.PublicKey, hrp)
	if err != nil {
		return
	}

	acc.PublicKeyText, err = atom.ConvertPub(acc.PublicKey, hrp)
	if err != nil {
		return
	}

	acc.PrivateKeyText = base64.StdEncoding.EncodeToString(acc.PrivateKey)
	return
}

func (acc *Account) ToTHORChain() (err error) {
	hrp := "thor"
	acc.Address, err = atom.ConvertAddr(acc.PublicKey, hrp)
	if err != nil {
		return
	}

	acc.PublicKeyText = acc.PublicKeyHex
	acc.PrivateKeyText = acc.PrivateKeyHex
	return
}

func (acc *Account) ToWhitecoin() (err error) {
	acc.Address = xwc.ConvertAddr(acc.PublicKey)
	acc.PublicKeyText = xwc.ConvertPub(acc.PublicKey)
	acc.PrivateKeyText = xwc.ConvertPriv(acc.PrivateKey)
	return
}

func (acc *Account) ToTerra() (err error) {
	hrp := "terra"
	acc.Address, err = atom.ConvertAddr(acc.PublicKey, hrp)
	if err != nil {
		return
	}

	acc.PublicKeyText = acc.PublicKeyHex
	acc.PrivateKeyText = acc.PrivateKeyHex
	return
}

func (acc *Account) ToStarname() (err error) {
	hrp := "star"
	acc.Address, err = atom.ConvertAddr(acc.PublicKey, hrp)
	if err != nil {
		return
	}

	acc.PublicKeyText, err = atom.ConvertPub(acc.PublicKey, hrp)
	if err != nil {
		return
	}

	acc.PrivateKeyText = base64.StdEncoding.EncodeToString(acc.PrivateKey)
	return
}
