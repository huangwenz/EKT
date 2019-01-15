package downloader

import (
	"bytes"
	"encoding/hex"
	"github.com/EducationEKT/EKT/crypto"
	"github.com/EducationEKT/EKT/param"
)

func Synchronise(key []byte) []byte {
	for _, node := range param.MainChainDelegateNode {
		value, err := node.GetDBValue(hex.EncodeToString(key))
		if err == nil && bytes.Equal(crypto.Sha3_256(value), key) {
			return value
		}
	}
	return nil
}
