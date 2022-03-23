package rsk_test

import (
	"testing"

	"github.com/keng42/xwallet/pkg/chains/rsk"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	addr := rsk.ConvertAddr("23373b6586b0f83746fb657a121a46c5f19a1a7c", 30)
	require.Equal(t, "0x23373b6586B0f83746Fb657a121A46c5f19A1A7c", addr)
}
