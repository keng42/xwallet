package swtc_test

import (
	"testing"

	"github.com/keng42/xwallet/pkg/chains/swtc"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	addr, err := swtc.ConvertAddr("15fPsb15EM9iumsXNkDDXEMES88G2iZsM1")
	require.Nil(t, err)
	require.Equal(t, "jnCP1bjnNM95um1X4kDDXNMNS33Gp5Z1Mj", addr)

	priv, err := swtc.ConvertPriv("L2dCa9j2WggXMV46HBHEb7pFWLWGhXcnLRCGZLqRPbiej5oJKwFf")
	require.Nil(t, err)
	require.Equal(t, "a14555f94abb29fa12069eb2689d1ad39f97369274bf2aefa9a17556fc0bb7d7", priv)
}
