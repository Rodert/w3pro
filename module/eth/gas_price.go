package eth

import (
	"math/big"

	"github.com/lmittmann/w3/w3types"
	"github.com/rodert/w3pro/internal/module"
)

// GasPrice requests the current gas price in wei.
func GasPrice() w3types.CallerFactory[big.Int] {
	return module.NewFactory(
		"eth_gasPrice",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}