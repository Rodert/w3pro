package eth_test

import (
	"math/big"
	"testing"

	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/rpctest"
)

func TestGasPrice(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "gas_price",
			Call:    eth.GasPrice(),
			WantRet: *w3pro.I("0xc0fe"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
