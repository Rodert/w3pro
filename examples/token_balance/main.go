/*
获取指定账户的指定代币余额

token_balance prints the balance of an ERC20 token for a given account.

Usage:

	token_balance [flags]

Flags:

	-acc string
		Account address (default "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	-token string
		Token address (default "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
	-h, --help
		help for token_balance
*/
package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
)

var (
	// smart contract functions
	funcName      = w3pro.MustNewFunc("name()", "string")
	funcSymbol    = w3pro.MustNewFunc("symbol()", "string")
	funcDecimals  = w3pro.MustNewFunc("decimals()", "uint8")
	funcBalanceOf = w3pro.MustNewFunc("balanceOf(address)", "uint256")

	// flags
	addrAcc   common.Address
	addrToken common.Address
)

func main() {
	// parse flags
	flag.TextVar(&addrAcc, "acc", w3pro.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), "Account address")
	flag.TextVar(&addrToken, "token", w3pro.A("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), "Token address")
	flag.Usage = func() {
		fmt.Println("token_balance prints the balance of an ERC20 token for a given account.")
		flag.PrintDefaults()
	}
	flag.Parse()

	// connect to RPC endpoint
	client := w3pro.MustDial("https://rpc.ankr.com/eth")
	defer client.Close()

	// fetch token details and account balance
	var (
		name, symbol string
		decimals     uint8
		balance      big.Int
	)
	if err := client.Call(
		eth.CallFunc(funcName, addrToken).Returns(&name),
		eth.CallFunc(funcSymbol, addrToken).Returns(&symbol),
		eth.CallFunc(funcDecimals, addrToken).Returns(&decimals),
		eth.CallFunc(funcBalanceOf, addrToken, addrAcc).Returns(&balance),
	); err != nil {
		fmt.Printf("Call failed: %v\n", err)
		return
	}

	fmt.Printf("%s balance of %s: %s %s\n", name, addrAcc, w3pro.FromWei(&balance, decimals), symbol)
}
