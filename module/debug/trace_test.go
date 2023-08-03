package debug_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/debug"
	"github.com/rodert/w3pro/rpctest"
	"github.com/rodert/w3pro/w3types"
)

func TestTraceTx(t *testing.T) {
	tests := []rpctest.TestCase[debug.Trace]{
		{
			Golden: "traceTx__1150000_0",
			Call:   debug.TraceTx(w3pro.H("0x38f299591902bfada359527fa6b9b597a959c41c6f72a3b484807fbf52dc8abe"), nil),
			WantRet: debug.Trace{
				Gas: 22224,
			},
		},
		{
			Golden: "traceTx__12244000_0",
			Call:   debug.TraceTx(w3pro.H("0xac503dd98281d4d52c2043e297a6e684d175339a7ebf831605fe593f01ce82c3"), &debug.TraceConfig{EnableStack: true, EnableMemory: true, EnableStorage: true, Limit: 3}),
			WantRet: debug.Trace{
				Gas: 46121,
				StructLogs: []*debug.StructLog{
					{Pc: 0, Op: vm.PUSH1, Gas: 228380, GasCost: 3, Depth: 1},
					{Pc: 2, Op: vm.PUSH1, Gas: 228377, GasCost: 3, Depth: 1, Stack: []uint256.Int{*uint256.NewInt(0x60)}},
					{Pc: 4, Op: vm.MSTORE, Gas: 228374, GasCost: 12, Depth: 1, Stack: []uint256.Int{*uint256.NewInt(0x60), *uint256.NewInt(0x40)},
						Memory: w3pro.B("0x" +
							"0000000000000000000000000000000000000000000000000000000000000000" +
							"0000000000000000000000000000000000000000000000000000000000000000" +
							"0000000000000000000000000000000000000000000000000000000000000000"),
					},
				},
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}

func TestTraceCall(t *testing.T) {
	tests := []rpctest.TestCase[debug.Trace]{
		{
			Golden: "traceCall",
			Call: debug.TraceCall(&w3types.Message{
				From:  w3pro.A("0x000000000000000000000000000000000000c0Fe"),
				To:    w3pro.APtr("0x000000000000000000000000000000000000dEaD"),
				Value: w3pro.I("1 ether"),
			}, nil, &debug.TraceConfig{Overrides: w3types.State{
				w3pro.A("0x000000000000000000000000000000000000c0Fe"): {Balance: w3pro.I("1 ether")},
			}}),
			WantRet: debug.Trace{
				Gas: 21000,
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}
