package eos_test

import (
	"encoding/hex"
	"testing"

	"github.com/keng42/xwallet/pkg/chains/eos"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	buf, err := hex.DecodeString("032415ae4dfc06792ed6503f5465d854635785d579d2adbb430644abc859af25e6")
	require.Nil(t, err)
	pub := eos.ConvertPub(buf)
	require.Equal(t, "EOS778GKtDRLPhsox8Qp5HV4beng89EjAhWGTpC9r6dfk6QKxFHCq", pub)

	buf, err = hex.DecodeString("1a3eddf858ecf2925d9a50f7bac532b8428a03183393540c875ba2d088e07e25")
	require.Nil(t, err)
	priv := eos.ConvertPriv(buf)
	require.Equal(t, "5J1qzLswg64EFqnPEEYccX9vHHYqRbCTr6oQ71f6k3ct6tbzWkh", priv)
}
