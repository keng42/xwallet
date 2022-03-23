package csc_test

import (
	"testing"

	"github.com/keng42/xwallet/pkg/chains/csc"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	addr, err := csc.ConvertAddr("1KsFfVwUNe96MUsHCtqpainvkiHZSruHQV")
	require.Nil(t, err)
	require.Equal(t, "cK1ECVA74e9aM71HUtqF258vk5HZSiuHQV", addr)

	priv, err := csc.ConvertPriv("L1vb44qPA6cMyJ7kbpZKjSj9pnqW4pYsW37E3sjj1nRmHrQshzdY")
	require.Nil(t, err)
	require.Equal(t, "8c60d7acc676539f58d0860ba36e84a934332cc2dab5fb26d6c50283b1f0487b", priv)
}
