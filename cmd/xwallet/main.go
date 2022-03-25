package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/keng42/xwallet/pkg/chains/bch"
	"github.com/keng42/xwallet/pkg/mnemonic"
	"github.com/keng42/xwallet/pkg/networks"
	"github.com/keng42/xwallet/pkg/utilities/info"
	"github.com/keng42/xwallet/pkg/wallet"
	"github.com/tyler-smith/go-bip32"

	// TODO code review
	"github.com/olekukonko/tablewriter"

	// TODO code review
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Usage: "A cryptocurrency wallet tool",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("version") {
				fmt.Printf("v%s BuildTime(%s) GitRev(%s)\n", info.Version, info.BuildTime, info.GitRev)
			} else {
				cli.ShowAppHelp(c)
			}
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "network",
				Aliases: []string{"n", "nw"},
				Usage:   "network utilities",
				Subcommands: []*cli.Command{
					{
						Name:  "ls",
						Usage: "list all networks",
						Action: func(c *cli.Context) error {
							networks.DisplayAll()
							return nil
						},
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "mnemonic",
				Aliases: []string{"m", "mne"},
				Usage:   "mnemonic utilities",
				Subcommands: []*cli.Command{
					{
						Name:  "new",
						Usage: "create new mnemonic",
						Flags: []cli.Flag{
							&cli.Int64Flag{
								Name:    "words",
								Value:   12,
								Aliases: []string{"w"},
							},
						},
						Action: CreateNewMnemonicAction,
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "address",
				Aliases: []string{"a", "addr"},
				Usage:   "address utilities",
				Subcommands: []*cli.Command{
					{
						Name:  "new",
						Usage: "create new addresses",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "network",
								Aliases: []string{"n"},
								Value:   "BTC - Bitcoin",
								Usage:   "specify the network",
							},
							&cli.StringFlag{
								Name:    "mnemonic",
								Aliases: []string{"m"},
								Usage:   "specify the mnemonic, or generate a random one if it is empty",
							},
							&cli.StringFlag{
								Name:    "path",
								Aliases: []string{"p"},
								Value:   "bip44",
								Usage:   "specify bip32 extended key derivation path (options: bip44, bip49, bip84, m/*)",
							},
							&cli.Uint64Flag{
								Name:    "start",
								Aliases: []string{"s"},
								Value:   0,
								Usage:   "first index",
							},
							&cli.Uint64Flag{
								Name:    "end",
								Aliases: []string{"e"},
								Value:   0,
								Usage:   "last index",
							},
							&cli.BoolFlag{
								Name:    "harden",
								Aliases: []string{"hard"},
								Usage:   "use hardened addresses",
							},
							&cli.StringFlag{
								Name:    "segwit",
								Aliases: []string{"sw"},
								Value:   "",
								Usage:   "use segwit addresses (options: p2wpkh, p2wpkh_p2sh, p2wsh, p2wsh_p2sh)",
							},
							&cli.StringFlag{
								Name:    "bch",
								Aliases: []string{""},
								Value:   "",
								Usage:   "use legacy or BitPay style addresses for bitcoin cash, default is CashAddr (options: bitpay, legacy)",
							},
						},
						Action: CreateNewAddressesAction,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateNewMnemonicAction(c *cli.Context) error {
	s, err := mnemonic.Generate(c.Int("words"))
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}

func CreateNewAddressesAction(c *cli.Context) error {
	nw, err := networks.New(c.String("network"))
	if err != nil {
		return err
	}

	start := uint32(c.Uint64("start"))
	end := uint32(c.Uint64("end"))
	if start > end {
		return errors.New("start should equal or less than end")
	}
	if c.Bool("harden") {
		start += bip32.FirstHardenedChild
		end += bip32.FirstHardenedChild
	}

	mne := c.String("mnemonic")
	if mne == "" {
		mne, err = mnemonic.Generate(12)
		if err != nil {
			return err
		}
		fmt.Printf("Generated new mnemonic: %s\n\n", mne)
	}
	w := wallet.Wallet{
		Mnemonic: mne,
	}

	path := c.String("path")
	switch path {
	case "bip44":
		path = nw.Bip44Path()
		break
	case "bip49":
		path = nw.Bip49Path()
		break
	case "bip84":
		path = nw.Bip84Path()
		break
	}

	segwit := c.String("segwit")
	segwitType := wallet.SEGWIT_TYPE_DISABLE
	switch segwit {
	case "":
		break
	case "p2wpkh":
		segwitType = wallet.SEGWIT_TYPE_P2WPKH
		break
	case "p2wpkh_p2sh":
		segwitType = wallet.SEGWIT_TYPE_P2WPKH_P2SH
		break
	case "p2wsh":
		segwitType = wallet.SEGWIT_TYPE_P2WSH
		break
	case "p2wsh_p2sh":
		segwitType = wallet.SEGWIT_TYPE_P2WSH_P2SH
		break
	default:
		return errors.New("segwit available options: p2wpkh, p2wpkh_p2sh, p2wsh, p2wsh_p2sh")
	}

	_bch := c.String("bch")
	bchType := bch.ADDR_STYLE_CASH
	switch _bch {
	case "":
		break
	case "bitpay":
		bchType = bch.ADDR_STYLE_BITPAY
		break
	case "legacy":
		bchType = bch.ADDR_STYLE_LEGACY
		break
	default:
		return errors.New("bch available options: bitpay, legacy")
	}

	data := make([][]string, 0, end-start+1)

	for i := start; i <= end; i++ {
		acc, err := w.NewAddress(nw, path, i, segwitType)
		if err != nil {
			return err
		}

		if nw.IsBitcoinCash() && bchType != bch.ADDR_STYLE_CASH {
			acc.ToBitcoinCash(nil, bchType)
		}

		data = append(data, []string{acc.Path, acc.Address, acc.PublicKeyText, acc.PrivateKeyText})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Path", "Address", "Public Key", "Private Key"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	return nil
}
