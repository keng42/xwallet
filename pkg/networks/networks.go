package networks

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type Network struct {
	Name            string `json:"name"`
	Network         string `json:"network"`
	HdCoin          int64  `json:"hdCoin"`
	SegwitAvailable bool   `json:"segwitAvailable"`
	Bip49available  bool   `json:"bip49available"`

	Params Params `json:"_"`
}

type Params struct {
	BaseNetwork   string `json:"baseNetwork"`
	MessagePrefix string `json:"messagePrefix"`
	Bip32         Bip32  `json:"bip32"`
	PubKeyHash    int64  `json:"pubKeyHash"`
	ScriptHash    int64  `json:"scriptHash"`
	Wif           int64  `json:"wif"`
	Bech32        string `json:"bech32"`

	P2wpkh       *Params `json:"p2wpkh,omitempty"`
	P2wpkhInP2sh *Params `json:"p2wpkhInP2sh,omitempty"`
	P2wsh        *Params `json:"p2wsh,omitempty"`
	P2wshInP2sh  *Params `json:"p2wshInP2sh,omitempty"`
}

type Bip32 struct {
	Public  int64 `json:"public"`
	Private int64 `json:"private"`
}

//go:embed params.json
var paramsjson []byte

//go:embed networks.json
var networksjson []byte

var NetworkParams map[string]Params
var NetworksMap map[string]Network
var Networks []Network

func init() {
	var err error

	err = json.Unmarshal(paramsjson, &NetworkParams)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(networksjson, &Networks)
	if err != nil {
		log.Fatal(err)
	}

	NetworksMap = make(map[string]Network, len(Networks))

	for i, nw := range Networks {
		params, ok := NetworkParams[nw.Network]
		if ok {
			Networks[i].Params = params
		} else {
			// fmt.Println("=====", nw.Name, nw.Network)
		}

		NetworksMap[nw.Name] = Networks[i]
	}
}

func New(name string) (nw Network, err error) {
	var ok bool
	nw, ok = NetworksMap[name]
	if !ok {
		err = errors.New("unknow network")
		return
	}
	return
}

func (n *Network) IsGRS() bool {
	return n.Name == "GRS - Groestlcoin" ||
		n.Name == "GRS - Groestlcoin Testnet"
}

func (n *Network) IsELA() bool {
	return n.Name == "ELA - Elastos"
}

func (n *Network) IsEthereum() bool {
	name := n.Name
	return (name == "ETH - Ethereum") ||
		(name == "ETC - Ethereum Classic") ||
		(name == "EWT - EnergyWeb") ||
		(name == "PIRL - Pirl") ||
		(name == "MIX - MIX") ||
		(name == "MOAC - MOAC") ||
		(name == "MUSIC - Musicoin") ||
		(name == "POA - Poa") ||
		(name == "EXP - Expanse") ||
		(name == "CLO - Callisto") ||
		(name == "DXN - DEXON") ||
		(name == "ELLA - Ellaism") ||
		(name == "ESN - Ethersocial Network") ||
		(name == "VET - VeChain") ||
		(name == "ERE - EtherCore") ||
		(name == "BSC - Binance Smart Chain")
}

func (n *Network) IsRsk() bool {
	var name = n.Name
	return name == "R-BTC - RSK" || name == "tR-BTC - RSK Testnet"
}

func (n *Network) IsStellar() bool {
	return n.Name == "XLM - Stellar"
}

func (n *Network) IsNano() bool {
	return n.Name == "NANO - nano"
}

func (n *Network) Bip44Path() string {
	return fmt.Sprintf("m/44'/%d'/0'/0", n.HdCoin)
}

func (n *Network) Bip49Path() string {
	return fmt.Sprintf("m/49'/%d'/0'/0", n.HdCoin)
}

func (n *Network) Bip84Path() string {
	return fmt.Sprintf("m/84'/%d'/0'/0", n.HdCoin)
}

func DisplayAll() {
	for _, v := range Networks {
		fmt.Println(v.Name)
	}
}
