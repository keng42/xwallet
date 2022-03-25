# xwallet

A tool for converting BIP39 mnemonic phrases to addresses and private keys, creating and signing transactions.

It's inspired by [iancoleman/bip39](https://github.com/iancoleman/bip39).

[![API Reference](https://pkg.go.dev/badge/github.com/keng42/xwallet)](https://pkg.go.dev/github.com/keng42/xwallet)
[![Build and Create Github Release](https://github.com/keng42/xwallet/actions/workflows/release.yml/badge.svg)](https://github.com/keng42/xwallet/actions/workflows/release.yml)
[![Coverage Status](https://coveralls.io/repos/github/keng42/xwallet/badge.svg)](https://coveralls.io/github/keng42/xwallet)
[![Go Report Card](https://goreportcard.com/badge/github.com/keng42/xwallet)](https://goreportcard.com/report/github.com/keng42/xwallet)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/keng42/xwallet/blob/master/LICENSE)

## Features

- [x] generate mnemonic
- [x] generate addresses and private keys
- [ ] create and sign transactions

## Installation

Download a binary from the [releases page](https://github.com/keng42/xwallet/releases)
and place it in a directory in your $PATH.

### Go users

```
go get github.com/keng42/xwallet
```

## CLI Usage

### Display supported networks

```sh
xwallet network ls
```

### Generate random mnemonic

```sh
xwallet mnemonic new

# specify words length
xwallet mnemonic new --words=24
```

### Generate new addresses

```sh
xwallet address new

# specify options
xwallet address new \
  --network="btc" \
  --mnemonic="term mushroom resemble heavy calm tribe leader aim coyote polar during neglect" \
  --path="bip44" \
  --start=0 \
  --end=2 \
  --harden=false
```

## Networks support address generation

- [x] AC - Asiacoin
- [x] ACC - Adcoin
- [x] AGM - Argoneum
- [x] ANON - ANON
- [x] ARYA - Aryacoin
- [x] AUR - Auroracoin
- [x] AXE - Axe
- [x] BCA - Bitcoin Atom
- [x] BCH - Bitcoin Cash
- [x] BEET - Beetlecoin
- [x] BELA - Belacoin
- [x] BITG - Bitcoin Green
- [x] BLK - BlackCoin
- [x] BND - Blocknode
- [x] BOLI - Bolivarcoin
- [x] BRIT - Britcoin
- [x] BSC - Binance Smart Chain
- [x] BSD - Bitsend
- [x] BST - BlockStamp
- [x] BSV - BitcoinSV
- [x] BTA - Bata
- [x] BTC - Bitcoin
- [x] BTC - Bitcoin RegTest
- [x] BTC - Bitcoin Testnet
- [x] BTCP - Bitcoin Private
- [x] BTCPt - Bitcoin Private Testnet
- [x] BTCZ - Bitcoinz
- [x] BTDX - BitCloud
- [x] BTG - Bitcoin Gold
- [x] BTX - Bitcore
- [x] CCN - Cannacoin
- [x] CDN - Canadaecoin
- [x] CESC - Cryptoescudo
- [x] CLAM - Clams
- [x] CLO - Callisto
- [x] CLUB - Clubcoin
- [x] CMP - Compcoin
- [x] CPU - CPUchain
- [x] CRAVE - Crave
- [x] CRP - CranePay
- [x] CRW - Crown
- [x] CRW - Crown (Legacy)
- [x] CSC - CasinoCoin
- [x] DASH - Dash
- [x] DASH - Dash Testnet
- [x] DFC - Defcoin
- [x] DGB - Digibyte
- [x] DGC - Digitalcoin
- [x] DIVI - DIVI
- [x] DIVI - DIVI Testnet
- [x] DMD - Diamond
- [x] DNR - Denarius
- [x] DOGE - Dogecoin
- [x] DOGEt - Dogecoin Testnet
- [x] DXN - DEXON
- [x] ECN - Ecoin
- [x] EDRC - Edrcoin
- [x] EFL - Egulden
- [x] ELLA - Ellaism
- [x] EMC2 - Einsteinium
- [x] EOS - EOSIO
- [x] ERC - Europecoin
- [x] ERE - EtherCore
- [x] ESN - Ethersocial Network
- [x] ETC - Ethereum Classic
- [x] ETH - Ethereum
- [x] EWT - EnergyWeb
- [x] EXCC - ExchangeCoin
- [x] EXCL - Exclusivecoin
- [x] EXP - Expanse
- [x] FIO - Foundation for Interwallet Operability
- [x] FIRO - Firo (Zcoin rebrand)
- [x] FIX - FIX
- [x] FIX - FIX Testnet
- [x] FJC - Fujicoin
- [x] FLASH - Flashcoin
- [x] FRST - Firstcoin
- [x] FTC - Feathercoin
- [x] GAME - GameCredits
- [x] GBX - Gobyte
- [x] GCR - GCRCoin
- [x] GRC - Gridcoin
- [x] HNC - Helleniccoin
- [x] HNS - Handshake
- [x] HUSH - Hush (Legacy)
- [x] HUSH - Hush3
- [x] INSN - Insane
- [x] IOP - Iop
- [x] IOV - Starname
- [x] IXC - Ixcoin
- [x] JBS - Jumbucks
- [x] KMD - Komodo
- [x] KOBO - Kobocoin
- [x] LBC - Library Credits
- [x] LCC - Litecoincash
- [x] LDCN - Landcoin
- [x] LINX - Linx
- [x] LKR - Lkrcoin
- [x] LTC - Litecoin
- [x] LTCt - Litecoin Testnet
- [x] LTZ - LitecoinZ
- [x] LUNA - Terra
- [x] LYNX - Lynx
- [x] MAZA - Maza
- [x] MEC - Megacoin
- [x] MIX - MIX
- [x] MNX - Minexcoin
- [x] MOAC - MOAC
- [x] MONA - Monacoin
- [x] MONK - Monkey Project
- [x] MUSIC - Musicoin
- [x] NANO - Nano
- [x] NAV - Navcoin
- [x] NEBL - Neblio
- [x] NEOS - Neoscoin
- [x] NIX - NIX Platform
- [x] NLG - Gulden
- [x] NMC - Namecoin
- [x] NRG - Energi
- [x] NRO - Neurocoin
- [x] NSR - Nushares
- [x] NVC - Novacoin
- [x] NYC - Newyorkc
- [x] OK - Okcash
- [x] OMNI - Omnicore
- [x] ONION - DeepOnion
- [x] ONX - Onixcoin
- [x] PART - Particl
- [x] PHR - Phore
- [x] PINK - Pinkcoin
- [x] PIRL - Pirl
- [x] PIVX - PIVX
- [x] PIVX - PIVX Testnet
- [x] POA - Poa
- [x] POSW - POSWcoin
- [x] POT - Potcoin
- [x] PPC - Peercoin
- [x] PRJ - ProjectCoin
- [x] PSB - Pesobit
- [x] PUT - Putincoin
- [x] R-BTC - RSK
- [x] RBY - Rubycoin
- [x] RDD - Reddcoin
- [x] RITO - Ritocoin
- [x] RPD - Rapids
- [x] RUNE - THORChain
- [x] RVN - Ravencoin
- [x] RVR - RevolutionVR
- [x] SAFE - Safecoin
- [x] SCRIBE - Scribe
- [x] SDC - ShadowCash
- [x] SDC - ShadowCash Testnet
- [x] SLM - Slimcoin
- [x] SLM - Slimcoin Testnet
- [x] SLP - Simple Ledger Protocol
- [x] SLR - Solarcoin
- [x] SLS - Salus
- [x] SMLY - Smileycoin
- [x] STASH - Stash
- [x] STASH - Stash Testnet
- [x] STRAT - Stratis
- [x] SUGAR - Sugarchain
- [x] SWTC - Jingtum
- [x] SYS - Syscoin
- [x] tBND - Blocknode Testnet
- [x] THC - Hempcoin
- [x] THT - Thought
- [x] TOA - Toa
- [x] tR-BTC - RSK Testnet
- [x] TRX - Tron
- [x] TSTRAT - Stratis Testnet
- [x] TUGAR - Sugarchain Testnet
- [x] TWINS - TWINS
- [x] TWINS - TWINS Testnet
- [x] UNO - Unobtanium
- [x] USC - Ultimatesecurecash
- [x] USNBT - NuBits
- [x] VASH - Vpncoin
- [x] VET - VeChain
- [x] VIA - Viacoin
- [x] VIA - Viacoin Testnet
- [x] VIVO - Vivo
- [x] VTC - Vertcoin
- [x] WC - Wincoin
- [x] WGR - Wagerr
- [x] XAX - Artax
- [x] XBC - Bitcoinplus
- [x] XLM - Stellar
- [x] XMY - Myriadcoin
- [x] XRP - Ripple
- [x] XUEZ - Xuez
- [x] XVC - Vcash
- [x] XVG - Verge
- [x] XWC - Whitecoin
- [x] XWCC - Whitecoin Classic
- [x] XZC - Zcoin (rebranded to Firo)
- [x] ZCL - Zclassic
- [x] ZEC - Zcash
- [x] ZEN - Horizen
- [ ] ELA - Elastos
- [ ] GRS - Groestlcoin
- [ ] GRS - Groestlcoin Testnet
- [ ] NAS - Nebulas
- [ ] ZBC - ZooBlockchain

## More info

**BIP39** Mnemonic code for generating deterministic keys  
Read more at the [official BIP39 spec](https://github.com/bitcoin/bips/blob/master/bip-0039.mediawiki)

**BIP32** Hierarchical Deterministic Wallets  
Read more at the [official BIP32 spec](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)  
See the demo at [bip32.org](bip32.org)

**BIP44** Multi-Account Hierarchy for Deterministic Wallets  
Read more at the [official BIP44 spec](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki)

**BIP49** Derivation scheme for P2WPKH-nested-in-P2SH based accounts  
Read more at the [official BIP49 spec](https://github.com/bitcoin/bips/blob/master/bip-0049.mediawiki)

**BIP85** Deterministic Entropy From BIP32 Keychains  
Read more at the [official BIP85 spec](https://github.com/bitcoin/bips/blob/master/bip-0085.mediawiki)
