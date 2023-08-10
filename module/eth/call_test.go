package eth_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/rpctest"
	"github.com/rodert/w3pro/w3protypes"
)

var (
	funcBalanceOf = w3pro.MustNewFunc("balanceOf(address)", "uint256")
)

func TestCall(t *testing.T) {
	tests := []rpctest.TestCase[[]byte]{
		{
			Golden: "call_func",
			Call: eth.Call(&w3protypes.Message{
				To:   w3pro.APtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Func: funcBalanceOf,
				Args: []any{w3pro.A("0x000000000000000000000000000000000000c0Fe")},
			}, nil, nil),
			WantRet: common.BigToHash(big.NewInt(0)).Bytes(),
		},
		{
			Golden: "call_func__overrides",
			Call: eth.Call(&w3protypes.Message{
				To:   w3pro.APtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Func: funcBalanceOf,
				Args: []any{w3pro.A("0x000000000000000000000000000000000000c0Fe")},
			}, nil, w3protypes.State{
				w3pro.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"): &w3protypes.Account{
					Storage: map[common.Hash]common.Hash{
						w3pro.H("0xf68b260b81af177c0bf1a03b5d62b15aea1b486f8df26c77f33aed7538cfeb2c"): w3pro.H("0x000000000000000000000000000000000000000000000000000000000000002a"),
					},
				},
			}),
			WantRet: common.BigToHash(big.NewInt(42)).Bytes(),
		},
	}

	rpctest.RunTestCases(t, tests)
}

func TestCallFunc(t *testing.T) {
	t.Parallel()

	srv := rpctest.NewFileServer(t, "testdata/call_func.golden")
	defer srv.Close()

	client := w3pro.MustDial(srv.URL())
	defer client.Close()

	var (
		balance     = new(big.Int)
		wantBalance = big.NewInt(0)
	)
	if err := client.Call(
		eth.CallFunc(funcBalanceOf, w3pro.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), w3pro.A("0x000000000000000000000000000000000000c0Fe")).Returns(balance),
	); err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	if wantBalance.Cmp(balance) != 0 {
		t.Fatalf("want %v, got %v", wantBalance, balance)
	}
}

func TestEstimateGas(t *testing.T) {
	tests := []rpctest.TestCase[uint64]{
		{
			Golden: "estimate_gas",
			Call: eth.EstimateGas(&w3protypes.Message{
				To:   w3pro.APtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Func: funcBalanceOf,
				Args: []any{w3pro.A("0x000000000000000000000000000000000000c0Fe")},
			}, nil),
			WantRet: 23750,
		},
	}

	rpctest.RunTestCases(t, tests)
}

func TestAccessList(t *testing.T) {
	tests := []rpctest.TestCase[eth.AccessListResponse]{
		{
			Golden: "create_access_list",
			Call: eth.AccessList(&w3protypes.Message{
				To:   w3pro.APtr("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
				Func: funcBalanceOf,
				Args: []any{w3pro.A("0x000000000000000000000000000000000000c0Fe")},
			}, nil),
			WantRet: eth.AccessListResponse{
				AccessList: types.AccessList{
					{
						Address: w3pro.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"),
						StorageKeys: []common.Hash{
							w3pro.H("0xf68b260b81af177c0bf1a03b5d62b15aea1b486f8df26c77f33aed7538cfeb2c"),
						},
					},
				},
				GasUsed: 26050,
			},
		},
	}

	rpctest.RunTestCases(t, tests)
}
