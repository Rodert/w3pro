package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3protypes"
)

// Balance requests the balance of the given common.Address addr at the given
// blockNumber. If blockNumber is nil, the balance at the latest known block is
// requested.
func Balance(addr common.Address, blockNumber *big.Int) w3protypes.CallerFactory[big.Int] {
	return module.NewFactory(
		"eth_getBalance",
		[]any{addr, module.BlockNumberArg(blockNumber)},
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
