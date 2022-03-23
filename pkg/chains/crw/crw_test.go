package crw_test

import (
	"testing"

	"github.com/keng42/xwallet/pkg/chains/crw"
	"github.com/stretchr/testify/require"
)

func TestA(t *testing.T) {
	addr, err := crw.ConvertAddr("1QAWf6rh8DEc2W8pW4gqxHC5NT6kvMmMnE")
	require.Nil(t, err)
	require.Equal(t, "CRWb8Vyf89EovPYhVSieGdRgUHg9iRewg2Mf", addr)
}
