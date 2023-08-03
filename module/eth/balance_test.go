package eth_test

import (
	"math/big"
	"testing"

	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/rpctest"
)

func TestBalance(t *testing.T) {
	tests := []rpctest.TestCase[big.Int]{
		{
			Golden:  "get_balance",
			Call:    eth.Balance(w3pro.A("0x000000000000000000000000000000000000c0Fe"), nil),
			WantRet: *w3pro.I("1 ether"),
		},
		{
			Golden:  "get_balance__at_block",
			Call:    eth.Balance(w3pro.A("0x000000000000000000000000000000000000c0Fe"), big.NewInt(255)),
			WantRet: *w3pro.I("0.1 ether"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
