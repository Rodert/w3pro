package eth

import (
	"github.com/lmittmann/w3/w3types"
	"github.com/rodert/w3pro/internal/module"
)

// ChainID requests the chains ID.
func ChainID() w3types.CallerFactory[uint64] {
	return module.NewFactory(
		"eth_chainId",
		nil,
		module.WithRetWrapper(module.HexUint64RetWrapper),
	)
}
