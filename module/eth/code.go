package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3protypes"
)

// Code requests the code of the given common.Address addr at the given
// blockNumber. If blockNumber is nil, the code at the latest known block is
// requested.
func Code(addr common.Address, blockNumber *big.Int) w3protypes.CallerFactory[[]byte] {
	return module.NewFactory(
		"eth_getCode",
		[]any{addr, module.BlockNumberArg(blockNumber)},
		module.WithRetWrapper(module.HexBytesRetWrapper),
	)
}
