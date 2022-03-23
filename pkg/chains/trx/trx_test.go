package trx_test

import (
	"testing"

	"github.com/keng42/xwallet/pkg/chains/trx"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	ethAddr := "0xe02F651FfD431153eCe9977E9b4155668b31e453"
	trxAddr := "TWQb96L4fWXE3NaPQBgyQ7TvsdfWsXEtrN"

	addr, err := trx.ConvertAddr(ethAddr)
	require.Nil(t, err)
	require.Equal(t, trxAddr, addr)

	addr, err = trx.UnconvertAddr(trxAddr)
	require.Nil(t, err)
	require.Equal(t, ethAddr, addr)
}
