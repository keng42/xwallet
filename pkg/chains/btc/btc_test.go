package btc_test

import (
	"encoding/hex"
	"testing"

	"github.com/keng42/xwallet/pkg/chains/btc"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	var addr string
	var err error
	var buf []byte

	// bip44
	buf, err = hex.DecodeString("02fb207bd0e7a3dff762a91a70d54fa65674b74aa6bc1e302deffc7bce5163643c")
	require.Nil(t, err)

	addr, err = btc.P2PKH(buf, nil)
	require.Nil(t, err)
	require.Equal(t, "1NaFxCzxCJyYK6fP4kn936bSbpxvNTaBe6", addr)

	// bip49
	buf, err = hex.DecodeString("023f3de32671ef4382147adf73fd2ca303be8db4d43ed0e1f6048518826ca649ff")
	require.Nil(t, err)

	addr, err = btc.P2SH(buf, nil)
	require.Nil(t, err)
	require.Equal(t, "352roE6imx6zKNQbgCn1wDTgGk5Lpieiee", addr)

	// bip141
	buf, err = hex.DecodeString("02b4c07b1b5bd4f3cbf4a68d861a0f66b2a4160a1afebdbe4f9b8a4778fef70bcd")
	require.Nil(t, err)

	addr, err = btc.P2WPKH(buf, nil) // bip84
	require.Nil(t, err)
	require.Equal(t, "bc1q4p6f4dzdv4djy4yq37t6g4ptagch2wd5nmynjm", addr)

	addr, err = btc.P2WPKHInP2SH(buf, nil) // bip49
	require.Nil(t, err)
	require.Equal(t, "32DHZrBtdwCUD4DksJr7cEadvRmJGqybE9", addr)

	addr, err = btc.P2WSH(buf, nil)
	require.Nil(t, err)
	require.Equal(t, "bc1qj2lxtnxd9wmskjk6kryzfgxmpwng4awjy97qfsg5n9wp3g2dpr0qd58jmr", addr)

	addr, err = btc.P2WSHInP2SH(buf, nil)
	require.Nil(t, err)
	require.Equal(t, "3K5Ck2kX2N9nnDt9DduH3Bn9dzFMibVdkV", addr)
	require.Nil(t, nil)
}

func TestB(t *testing.T) {
	var addr string
	var err error
	var buf []byte

	// zcash
	np := btc.NewNetworkParams(nil, 7352, 7357)

	// bip44
	buf, err = hex.DecodeString("02c407864604c45fdee03d99b6700193e73da5f24568a78a2a78b287040cd0173a")
	require.Nil(t, err)

	addr, err = btc.P2PKH(buf, &np)
	require.Nil(t, err)
	require.Equal(t, "t1KmTgxFzFLarBnAFMMXDwYii6uBNvAppyZ", addr)
}
