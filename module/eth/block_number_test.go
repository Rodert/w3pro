package eth_test

import (
	"math/big"
	"testing"

	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/rpctest"
)

func TestBlockNumber(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "block_number",
			Call:    eth.BlockNumber(),
			WantRet: *w3pro.I("0xc0fe"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
