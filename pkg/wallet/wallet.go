package wallet

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/keng42/xwallet/pkg/account"
	"github.com/keng42/xwallet/pkg/chains/bch"
	"github.com/keng42/xwallet/pkg/chains/btc"
	"github.com/keng42/xwallet/pkg/chains/nano"
	"github.com/keng42/xwallet/pkg/chains/xlm"
	"github.com/keng42/xwallet/pkg/mnemonic"
	"github.com/keng42/xwallet/pkg/networks"
	"github.com/stellar/go/keypair"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	Mnemonic  string
	Password  string
	MasterKey *bip32.Key
}

// TODO save wallet to file with keystore

type SegwitType int

const (
	SEGWIT_TYPE_DISABLE SegwitType = iota
	SEGWIT_TYPE_P2WPKH
	SEGWIT_TYPE_P2WPKH_P2SH
	SEGWIT_TYPE_P2WSH
	SEGWIT_TYPE_P2WSH_P2SH
)

var ErrUnsupportedNetwork = errors.New("unsupported network")

// NewAddress generates new address, public key and private key for specify path.
func (w *Wallet) NewAddress(
	network networks.Network, path string, index uint32, segwitType SegwitType,
) (acc account.Account, err error) {

	if !bip39.IsMnemonicValid(w.Mnemonic) {
		err = mnemonic.ErrInvalidMnemonic
		return
	}

	net := parseNetwork(network)

	// default path for bip44
	if path == "" {
		path = network.Bip44Path()

		if network.IsStellar() || network.IsNano() {
			path = fmt.Sprintf("m/44'/%d'", network.HdCoin)
			if index < bip32.FirstHardenedChild {
				index += bip32.FirstHardenedChild
			}
		}
	}

	// TODO cache keys
	key, err := mnemonic.GetBip32ExtendedKey(w.Mnemonic, w.Password, path)
	if err != nil {
		return
	}
	key, err = key.NewChildKey(index)
	if err != nil {
		return
	}

	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), key.Key)
	wif, err := btcutil.NewWIF(privKey, net.GetRaw(), true)
	if err != nil {
		return
	}

	if index >= bip32.FirstHardenedChild {
		acc.Path = fmt.Sprintf("%s/%d'", path, index-bip32.FirstHardenedChild)
	} else {
		acc.Path = fmt.Sprintf("%s/%d", path, index)
	}

	acc.PublicKey = key.PublicKey().Key
	acc.PrivateKey = key.Key

	acc.PublicKeyHex = hex.EncodeToString(acc.PublicKey)
	acc.PrivateKeyHex = hex.EncodeToString(acc.PrivateKey)
	acc.PrivateKeyWIF = wif.String()

	acc.PublicKeyText = acc.PublicKeyHex
	acc.PrivateKeyText = acc.PrivateKeyWIF

	acc.EcdsaPub = pubKey.ToECDSA()
	acc.EcdsaPriv = privKey.ToECDSA()

	acc.Address, err = btc.P2PKH(pubKey.SerializeCompressed(), net)
	if err != nil {
		return
	}

	// ===== ethereum
	if network.IsEthereum() {
		acc.ToEthereum()
		return
	}

	// ===== other networks
	switch network.Name {
	case "TRX - Tron":
		err = acc.ToTron()
		return

	case "R-BTC - RSK":
		acc.ToRsk(30)
		return
	case "tR-BTC - RSK Testnet":
		acc.ToRsk(31)
		return

	case "HNS - Handshake":
		err = acc.ToHandshake()
		return

	case "XLM - Stellar":
		var kp *keypair.Full
		kp, err = xlm.NewAddress(w.Mnemonic, w.Password, path, index)
		if err != nil {
			return
		}
		acc.ToStellar(kp)
		return

	case "NANO - Nano":
		addr, pub, priv, _err := nano.NewAddress(w.Mnemonic, w.Password, path, index)
		if _err != nil {
			err = _err
			return
		}
		acc.Address = addr
		acc.PublicKeyText = hex.EncodeToString(pub)
		acc.PrivateKeyText = hex.EncodeToString(priv)
		return

	case "NAS - Nebulas":
		err = ErrUnsupportedNetwork
		return

	case "XRP - Ripple":
		err = acc.ToRipple()
		return

	case "SWTC - Jingtum":
		err = acc.ToJingtum()
		return

	case "CSC - CasinoCoin":
		err = acc.ToCasinoCoin()
		return

	case "BCH - Bitcoin Cash":
		err = acc.ToBitcoinCash(net, bch.ADDR_STYLE_CASH)
		return

	case "SLP - Simple Ledger Protocol":
		err = acc.ToSimpleLedger(net)
		return

	case "ZBC - ZooBlockchain":
		err = ErrUnsupportedNetwork
		return

	}

	// Segwit addresses are different
	errSegwit := fmt.Errorf("segwit not available for %s", network.Name)
	switch segwitType {
	case SEGWIT_TYPE_P2WPKH:
		if network.Params.P2wpkh == nil {
			err = errSegwit
			return
		}
		acc.Address, err = btc.P2WPKH(acc.PublicKey, net)
		break
	case SEGWIT_TYPE_P2WPKH_P2SH:
		if network.Params.P2wpkhInP2sh == nil {
			err = errSegwit
			return
		}
		acc.Address, err = btc.P2WPKHInP2SH(acc.PublicKey, net)
		break
	case SEGWIT_TYPE_P2WSH:
		if network.Params.P2wsh == nil {
			err = errSegwit
			return
		}
		acc.Address, err = btc.P2WSH(acc.PublicKey, net)
		break
	case SEGWIT_TYPE_P2WSH_P2SH:
		if network.Params.P2wshInP2sh == nil {
			err = errSegwit
			return
		}
		acc.Address, err = btc.P2WSHInP2SH(acc.PublicKey, net)
		break
	default:
		break
	}
	if err != nil {
		return
	}

	switch network.Name {
	case "CRW - Crown":
		err = acc.ToCrown()
		return

	case "EOS - EOSIO":
		err = acc.ToEOSIO()
		return

	case "FIO - Foundation for Interwallet Operability":
		err = acc.ToFIO()
		return

	case "ATOM - Cosmos Hub":
		err = acc.ToCosmos()
		return

	case "RUNE - THORChain":
		err = acc.ToTHORChain()
		return

	case "XWC - Whitecoin":
		err = acc.ToWhitecoin()
		return

	case "LUNA - Terra":
		err = acc.ToTerra()
		return

	case "IOV - Starname":
		err = acc.ToStarname()
		return
	}

	if network.IsGRS() {
		err = ErrUnsupportedNetwork
		return
	}

	if network.IsELA() {
		err = ErrUnsupportedNetwork
		return
	}

	return
}

func parseNetwork(nw networks.Network) *btc.NetworkParams {
	raw := chaincfg.MainNetParams
	if strings.Contains(nw.Name, "Testnet") {
		raw = chaincfg.TestNet3Params
	} else if strings.Contains(nw.Name, "RegTest") {
		raw = chaincfg.RegressionNetParams
	}

	keyID := nw.Params.Wif
	if keyID == 0 {
		keyID = 0x80
	}
	raw.Bech32HRPSegwit = nw.Params.Bech32
	raw.PrivateKeyID = byte(keyID)
	raw.HDCoinType = uint32(nw.HdCoin)

	var privKeyID [4]byte
	binary.BigEndian.PutUint32(privKeyID[:], uint32(nw.Params.Bip32.Private))
	raw.HDPrivateKeyID = privKeyID

	var pubKeyID [4]byte
	binary.BigEndian.PutUint32(pubKeyID[:], uint32(nw.Params.Bip32.Public))
	raw.HDPublicKeyID = pubKeyID

	net := btc.NewNetworkParams(&raw, nw.Params.PubKeyHash, nw.Params.ScriptHash)
	return &net
}
