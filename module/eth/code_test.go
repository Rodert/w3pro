package eth_test

import (
	"testing"

	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/rpctest"
)

func TestCode(t *testing.T) {
	tests := []rpctest.TestCase[[]byte]{
		{
			Golden:  "get_code",
			Call:    eth.Code(w3pro.A("0x000000000000000000000000000000000000c0DE"), nil),
			WantRet: w3pro.B("0xdeadbeef"),
		},
	}

	rpctest.RunTestCases(t, tests)
}
