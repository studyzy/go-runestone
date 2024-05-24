package main

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/btcsuite/btcd/wire"
)

func TestReplaceTx(t *testing.T) {
	data, _ := hex.DecodeString("0100000000010169fb8135e4b8206725a5481a0b850c5d565309d77171e9721ceddfb1739adfb60100000000f5ffffff030000000000000000096a5d0614c0a2331441e80300000000000022512097ff5e14536b0049b2c8ad5a72e6e59771d2f222f36bdae5064971aaf590668891e70e000000000022512097ff5e14536b0049b2c8ad5a72e6e59771d2f222f36bdae5064971aaf59066880140f2b0f3f9a5b8e111afdc812352c1e438e57a6a842064bb549d04316301bbce14578be3ece20fbdc3d407a49c65b2b02168ce2571c415c0e1d2640a0f56130ce200000000")
	tx := wire.NewMsgTx(wire.TxVersion)
	tx.Deserialize(bytes.NewReader(data))
}
