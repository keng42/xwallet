// Package trx provides utilities for TRX - Tron.
package trx

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
	"github.com/keng42/xwallet/pkg/utilities/hash"
)

// ConvertEthToTrx converts eth address to trx address
func ConvertAddr(s string) (addr string, err error) {
	if !common.IsHexAddress(s) {
		err = errors.New("invalid address")
		return
	}

	s = strings.Replace(s, "0x", "41", 1)

	buf, err := hex.DecodeString(s)
	if err != nil {
		return
	}

	checksum := hash.DoubleSha256(buf)
	checksum = checksum[:4]

	buf = append(buf, checksum...)
	addr = base58.Encode(buf)

	return
}

// UnconvertEthToTrx converts trx address to eth address
func UnconvertAddr(s string) (addr string, err error) {
	buf := base58.Decode(s)
	if len(buf) < 4 {
		err = errors.New("invalid base58 string")
		return
	}

	data := buf[:len(buf)-4]
	checksum := buf[len(buf)-4:]
	_checksum := hash.DoubleSha256(data)

	for i := 0; i < 4; i++ {
		if _checksum[i] != checksum[i] {
			err = errors.New("invalid base58 checksum")
			return
		}
	}

	addr = hex.EncodeToString(data)
	addr = strings.Replace(addr, "41", "0x", 1)
	addr = common.HexToAddress(addr).Hex()

	return
}
