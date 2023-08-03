package eth_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/rpctest"
)

func TestStorageAt(t *testing.T) {
	tests := []rpctest.TestCase[common.Hash]{
		{
			Golden:  "get_storage_at",
			Call:    eth.StorageAt(w3pro.A("0x000000000000000000000000000000000000c0DE"), w3pro.H("0x0000000000000000000000000000000000000000000000000000000000000001"), nil),
			WantRet: w3pro.H("0x0000000000000000000000000000000000000000000000000000000000000042"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
