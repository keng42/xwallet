package xrp_test

import (
	"testing"

	"github.com/keng42/xwallet/pkg/chains/xrp"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	addr, err := xrp.ConvertAddr("1BZ7ChFWTYG8a3Q1YbnpKyXvECTaxcmwQU")
	require.Nil(t, err)
	require.Equal(t, "rBZfU6EWTYG32sQrYb8FKyXvNUT2xcmAQ7", addr)

	priv, err := xrp.ConvertPriv("L5APy9qHALKJ9vLw1bUBWRE83q83JY78aKjhrdoZKioCutXg3tRS")
	require.Nil(t, err)
	require.Equal(t, "ecfe8a07fadd51b943c20b60b44083005d1627a68e2da3a01390b4d222ab55be", priv)
}
