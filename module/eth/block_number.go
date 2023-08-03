package eth

import (
	"math/big"

	"github.com/lmittmann/w3/w3types"
	"github.com/rodert/w3pro/internal/module"
)

// BlockNumber requests the number of the most recent block.
func BlockNumber() w3types.CallerFactory[big.Int] {
	return module.NewFactory(
		"eth_blockNumber",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
