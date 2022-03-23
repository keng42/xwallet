package fio_test

import (
	"encoding/hex"
	"testing"

	"github.com/keng42/xwallet/pkg/chains/fio"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	buf, err := hex.DecodeString("02037520911d1ef0711e71ec7fd2e0507fe2a89a00de1165674ba2e3cff05cca6e")
	require.Nil(t, err)
	pub := fio.ConvertPub(buf)
	require.Equal(t, "FIO4v1fuNTqKGX25dR9migES31aQyfMk2P9pXPGdMe5XKM3CRyQkK", pub)

	buf, err = hex.DecodeString("6c6155da9262db7e00ecc9f168179367682a03c6a5d0dd220010bfc3a8d5704a")
	require.Nil(t, err)
	priv := fio.ConvertPriv(buf)
	require.Equal(t, "5Je21TRjhBLRvyh1kyN1mMawk4tjxPrAiFub7JhZhZp2KfoS6Qo", priv)
}
