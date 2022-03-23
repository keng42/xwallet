package atom_test

import (
	"encoding/hex"
	"testing"

	"github.com/keng42/xwallet/pkg/chains/atom"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	buf, err := hex.DecodeString("0388116318efe4875bc10da70a63981ad6a72f1c52f622a3e44adb49ff5f27f487")
	require.Nil(t, err)

	addr, err := atom.ConvertAddr(buf, "cosmos")
	require.Nil(t, err)
	require.Equal(t, "cosmos187ls3x4ezwcm29z0azpqqun7edazvklkfhd00r", addr)

	pub, err := atom.ConvertPub(buf, "cosmos")
	require.Equal(t, "cosmospub1addwnpepqwypzcccaljgwk7ppkns5cucrtt2wtcu2tmz9glyftd5nl6lyl6gwtj92fs", pub)
}
