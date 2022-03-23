package nano_test

import (
	"encoding/hex"
	"testing"

	"github.com/keng42/xwallet/pkg/chains/nano"
	"github.com/stretchr/testify/require"
	"github.com/tyler-smith/go-bip32"
)

var mne = "term mushroom resemble heavy calm tribe leader aim coyote polar during neglect"

func TestA(t *testing.T) {
	addr, pub, priv, err := nano.NewAddress(mne, "", "m/44'/165'", bip32.FirstHardenedChild+0)
	require.Nil(t, err)
	require.Equal(t, "nano_3kzubs5cwcc6feax6sd6rro3xnk6bu5wq74hycx3s94zx1bki9uz6dxfibqx", addr)
	require.Equal(t, "cbfb4e46ae29446b11d26564c62a1ed2444ec7cb944ff2ba1c9c5fe813281f7f", hex.EncodeToString(pub))
	require.Equal(t, "7d1197110dd1012b3c18151e5ace8d59e21b051c98bcad28ee9594330c967235", hex.EncodeToString(priv))
}
