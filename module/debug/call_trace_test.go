package debug_test

import (
	"testing"

	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/debug"
	"github.com/rodert/w3pro/rpctest"
	"github.com/rodert/w3pro/w3types"
)

func TestCallTraceTx(t *testing.T) {
	tests := []rpctest.TestCase[debug.CallTrace]{
		{
			Golden: "traceCall_callTracer",
			Call: debug.CallTraceCall(&w3types.Message{
				From:  w3pro.A("0x000000000000000000000000000000000000c0Fe"),
				To:    w3pro.APtr("0x000000000000000000000000000000000000dEaD"),
				Value: w3pro.I("1 ether"),
			}, nil, w3types.State{
				w3pro.A("0x000000000000000000000000000000000000c0Fe"): {Balance: w3pro.I("1 ether")},
			}),
			WantRet: debug.CallTrace{
				From:  w3pro.A("0x000000000000000000000000000000000000c0Fe"),
				To:    w3pro.A("0x000000000000000000000000000000000000dEaD"),
				Type:  "CALL",
				Gas:   49979000,
				Value: w3pro.I("1 ether"),
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}
