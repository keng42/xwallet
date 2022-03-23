package xwc_test

import (
	"encoding/hex"
	"testing"

	"github.com/keng42/xwallet/pkg/chains/xwc"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	buf, err := hex.DecodeString("032b8a52103322436ed855553d5d3d0bc6b326ea5ea52db8c71c6e03f2b3715fed")
	require.Nil(t, err)

	addr := xwc.ConvertAddr(buf)
	require.Equal(t, "XWCNLwEkdGF2LbsLoxPXpWX2WT2ZxgoeRKqHf", addr)

	pub := xwc.ConvertPub(buf)
	require.Equal(t, "XWC7AQi5AuHg5RUuSCdaVYCLRFNfu7pQYGHHYAhLftnKLvA2Tdec3", pub)

	buf, err = hex.DecodeString("1f023b7090eb9b937eac66f694af2fd65b2aae70e9aba648976a0ac67d488106")
	require.Nil(t, err)

	priv := xwc.ConvertPriv(buf)
	require.Equal(t, "5J3wf6XCeWY5Rz3HWbQECf1XFZNKpERDLvcv5syTWmw43tCDmPV", priv)
}
