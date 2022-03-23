package bch_test

import (
	"encoding/hex"
	"testing"

	"github.com/keng42/xwallet/pkg/chains/bch"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	// bch
	buf, err := hex.DecodeString("022bb58fb4028ff449e9af7c1e9da8c16577b0d54bcdd32f90e8684fc8cc8757e0")
	require.Nil(t, err)

	cAddr := "bitcoincash:qz0r7yu7ns40xm0yda8dlftc445a5km57gnqthn900"
	bAddr := "CWtctzVwTbBJPCnyrV9W35JXPSdBJPJ7Gd"

	_cAddr, err := bch.ConvertToCashAddr(buf, nil)
	require.Nil(t, err)
	require.Equal(t, cAddr, _cAddr, "bch cash addr not match")

	_bAddr, err := bch.ConvertToBitpayAddr(buf, nil)
	require.Nil(t, err)
	require.Equal(t, bAddr, _bAddr, "bch bitpay-style addr not match")

	// slp
	buf, err = hex.DecodeString("03efd1266b22fdecfc920d39d368a94cb76964bcebbfe8cc1d2a52c3415c91518c")
	require.Nil(t, err)

	cAddr = "simpleledger:qrkj8naglzd6vhpy2c7k8x55aplvzydmp5ha3jgwr7"
	bAddr = "Ce5mci56C4MpYA6k11pv19Ez1pfshuSn86"

	_cAddr, err = bch.ConvertToSlpCashAddr(buf, nil)
	require.Nil(t, err)
	require.Equal(t, cAddr, _cAddr, "slp cash addr not match")

	_bAddr, err = bch.ConvertToBitpayAddr(buf, nil)
	require.Nil(t, err)
	require.Equal(t, bAddr, _bAddr, "slp bitpay-style addr not match")
}
