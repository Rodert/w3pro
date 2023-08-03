package eth

import (
	"github.com/Rodert/w3pro/internal/module"
	"github.com/lmittmann/w3/w3types"
)

// ChainID requests the chains ID.
func ChainID() w3types.CallerFactory[uint64] {
	return module.NewFactory(
		"eth_chainId",
		nil,
		module.WithRetWrapper(module.HexUint64RetWrapper),
	)
}
