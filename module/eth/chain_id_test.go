package eth_test

import (
	"testing"

	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/rpctest"
)

func TestChainID(t *testing.T) {
	tests := []rpctest.TestCase[uint64]{
		{
			Golden:  "chain_id",
			Call:    eth.ChainID(),
			WantRet: 1,
		},
	}

	rpctest.RunTestCases(t, tests)
}
