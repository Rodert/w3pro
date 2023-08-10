package eth

import (
	"math/big"

	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3protypes"
)

// GasPrice requests the current gas price in wei.
func GasPrice() w3protypes.CallerFactory[big.Int] {
	return module.NewFactory(
		"eth_gasPrice",
		nil,
		module.WithRetWrapper(module.HexBigRetWrapper),
	)
}
