package eth

import (
	"math/big"

	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3protypes"
)

// BlockNumber requests the number of the most recent block.
func BlockNumber() w3protypes.CallerFactory[big.Int] {
	return module.NewFactory(
		"eth_blockNumber",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
