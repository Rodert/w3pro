/*
scan_blocks interates over blocks with their transactions from a given start block.

Usage:

	scan_blocks [flags]

Flags:

	-start uint
		Start block (default 10_000_000)
	-h, --help
		help for scan_blocks
*/
package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rodert/w3pro"
	"github.com/rodert/w3pro/module/eth"
	"github.com/rodert/w3pro/w3protypes"
)

var (
	// number of blocks to fetch in a single request
	bulkSize = 100

	// flags
	startBlock uint64
)

func main() {
	// parse flags
	flag.Uint64Var(&startBlock, "start", 10_000_000, "Start block")
	flag.Usage = func() {
		fmt.Println("scan_blocks interates over blocks with their transactions from a given start block.")
		flag.PrintDefaults()
	}
	flag.Parse()

	// connect to RPC endpoint
	client := w3pro.MustDial("https://rpc.ankr.com/eth")
	defer client.Close()

	// fetch blocks in bulk
	calls := make([]w3protypes.Caller, bulkSize)
	blocks := make([]types.Block, bulkSize)

	for i, txCount := 0, 0; ; i++ {
		j := i % bulkSize
		calls[j] = eth.BlockByNumber(new(big.Int).SetUint64(startBlock + uint64(i))).Returns(&blocks[j])

		if j == bulkSize-1 {
			if err := client.Call(calls...); err != nil {
				fmt.Printf("Call failed: %v\r", err)
				return
			}

			for _, block := range blocks {
				fmt.Printf("%+v\n\n", block)
				txCount += len(block.Transactions())
				processBlock(&block)
			}
			fmt.Printf("\rFetched %d blocks with a total of %d transactions", i+1, txCount)
		}
	}
}

func processBlock(b *types.Block) {
	// Do something with the block and its transactions...
}
