package main

// 判断指定区块中，包含指定合约的调用方法。

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/ethclient"
)

var path = flag.String("rpc", "https://polygon-rpc.com", "Polygon Bor JSON RPC path") // 可以替换自己的 jsonrpc 链接
var blockId = flag.Int64("block", 31740481, "Fetch this block from Polygon")

// https://polygonscan.com/address/0x0000000000000000000000000000000000001010
const contractAddr = "0x0000000000000000000000000000000000001010"

func Check(err error) {
	if err != nil {
		log.Fatalf("Got error: %v", err)
	}
}

func main() {
	flag.Parse()

	client, err := ethclient.Dial(*path)
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Chain ID: %d Latest: %08d\n", chainID.Uint64(), header.Number)

	fdata, err := ioutil.ReadFile("mrc20.abi")
	Check(err)

	contractAbi, err := abi.JSON(bytes.NewReader(fdata))
	Check(err)

	for m := range contractAbi.Methods {
		fmt.Printf("Found method: %s\n", m)
	}
	fmt.Printf("All methods found\n\n")

	bno := new(big.Int).SetInt64(*blockId)
	block, err := client.BlockByNumber(context.Background(), bno)
	Check(err)

	fmt.Printf("Found %d transactions in block: %d\n\n", len(block.Transactions()), *blockId)
	var valid int
	for _, tx := range block.Transactions() {
		// Only pick the txns which are interfacing with the right contract.
		if tx.To() == nil || tx.To().String() != contractAddr {
			continue
		}
		input := tx.Data()
		if len(input) < 4 {
			fmt.Printf("Input data of size: %d\n", len(input))
			continue
		}
		valid++
		m, err := contractAbi.MethodById(input[:4]) // first 4 bytes = method id.
		Check(err)

		in := make(map[string]interface{})
		Check(m.Inputs.UnpackIntoMap(in, input[4:]))

		fmt.Printf("Txn Hash: %s Method: %s Inputs: %+v\n", tx.Hash().String(), m.Name, in)
	}
	fmt.Printf("\nDone parsing Block: %d. Found %d txns involving contract.\n", *blockId, valid)
}
