package wallet_test

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/keng42/xwallet/pkg/networks"
	"github.com/keng42/xwallet/pkg/wallet"
	"github.com/stellar/go/exp/crypto/derivation"
	"github.com/stretchr/testify/require"
)

type Bip44Addr struct {
	Network string `json:"network"`
	Path    string `json:"path"`
	Address string `json:"address"`
	Pubkey  string `json:"pubkey"`
	Privkey string `json:"privkey"`
}

//go:embed bip44-addrs.json
var bip44addrsjson []byte

var Bip44Addrs map[string][]Bip44Addr

var w = wallet.Wallet{
	Mnemonic: "term mushroom resemble heavy calm tribe leader aim coyote polar during neglect",
}

func init() {
	var arr []Bip44Addr
	err := json.Unmarshal(bip44addrsjson, &arr)
	if err != nil {
		log.Fatal(err)
	}

	Bip44Addrs = make(map[string][]Bip44Addr, len(arr))

	for _, item := range arr {
		_, ok := Bip44Addrs[item.Network]
		if !ok {
			Bip44Addrs[item.Network] = make([]Bip44Addr, 0)
		}
		Bip44Addrs[item.Network] = append(Bip44Addrs[item.Network], item)
	}
}

func TestNewAddress(t *testing.T) {
	testNewAddress(t, "ETH - Ethereum")
	testNewAddress(t, "BTC - Bitcoin")
	testNewAddress(t, "TRX - Tron")
	testNewAddress(t, "R-BTC - RSK")
	testNewAddress(t, "tR-BTC - RSK Testnet")
	testNewAddress(t, "HNS - Handshake")
	testNewAddress(t, "XLM - Stellar")
	testNewAddress(t, "NANO - Nano")
	testNewAddress(t, "NAS - Nebulas")
	testNewAddress(t, "XRP - Ripple")
	testNewAddress(t, "SWTC - Jingtum")
	testNewAddress(t, "CSC - CasinoCoin")
	testNewAddress(t, "BCH - Bitcoin Cash")
	testNewAddress(t, "SLP - Simple Ledger Protocol")
	testNewAddress(t, "ZBC - ZooBlockchain")
	testNewAddress(t, "CRW - Crown")
	testNewAddress(t, "EOS - EOSIO")
	testNewAddress(t, "FIO - Foundation for Interwallet Operability")
	testNewAddress(t, "ATOM - Cosmos Hub")
	testNewAddress(t, "RUNE - THORChain")
	testNewAddress(t, "XWC - Whitecoin")
	testNewAddress(t, "LUNA - Terra")
	testNewAddress(t, "IOV - Starname")
}

func TestNewAddressAll(t *testing.T) {
	for _, n := range networks.Networks {
		testNewAddress(t, n.Name)
	}
}

func testNewAddress(t *testing.T, name string) {
	fmt.Println("===== name:", name)
	network, ok := networks.NetworksMap[name]
	require.True(t, ok)

	path := fmt.Sprintf("m/44'/%d'/0'/0", network.HdCoin)
	idxBase := uint32(0)

	if name == "XLM - Stellar" || name == "NANO - Nano" {
		path = fmt.Sprintf("m/44'/%d'", network.HdCoin)
		idxBase = derivation.FirstHardenedIndex
	}

	addrs, ok := Bip44Addrs[network.Name]
	require.True(t, ok)
	require.GreaterOrEqual(t, len(addrs), 2)

	info, err := w.NewAddress(network, path, idxBase+0, wallet.SEGWIT_TYPE_DISABLE)
	if err != nil {
		if errors.Is(err, wallet.ErrUnsupportedNetwork) || strings.Contains(err.Error(), "use fixed hd path") {
			return
		}
	}
	require.Nil(t, err)
	require.Equal(t, addrs[0].Address, info.Address, "address 0 not match")
	require.Equal(t, addrs[0].Pubkey, info.PublicKeyText, "public key 0 not match")
	require.Equal(t, addrs[0].Privkey, info.PrivateKeyText, "private key 0 not match")

	info, err = w.NewAddress(network, path, idxBase+1, wallet.SEGWIT_TYPE_DISABLE)
	require.Nil(t, err)
	require.Equal(t, addrs[1].Address, info.Address, "address 1 not match")
	require.Equal(t, addrs[1].Pubkey, info.PublicKeyText, "public key 1 not match")
	require.Equal(t, addrs[1].Privkey, info.PrivateKeyText, "private key 1 not match")
}

func BenchmarkNewAddress(b *testing.B) {
	b.ResetTimer()

	network, ok := networks.NetworksMap["ETH - Ethereum"]
	require.True(b, ok)

	path := fmt.Sprintf("m/44'/%d'/0'/0", network.HdCoin)
	idxBase := uint32(0)

	addrs, ok := Bip44Addrs[network.Name]
	require.True(b, ok)
	require.GreaterOrEqual(b, len(addrs), 2)

	for i := 0; i < b.N; i++ {
		_, err := w.NewAddress(network, path, idxBase+uint32(i), wallet.SEGWIT_TYPE_DISABLE)
		require.Nil(b, err)
	}

	// Result:
	// with cachedBip32ExtendedKey     BenchmarkNewAddress-12    	     279	   4189567 ns/op	 2780426 B/op	   25605 allocs/op
	// without cachedBip32ExtendedKey  BenchmarkNewAddress-12    	      96	  12289889 ns/op	 7455375 B/op	   68740 allocs/op
}
