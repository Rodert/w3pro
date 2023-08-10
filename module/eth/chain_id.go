package eth

import (
	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3protypes"
)

// ChainID requests the chains ID.
func ChainID() w3protypes.CallerFactory[uint64] {
	return module.NewFactory(
		"eth_chainId",
		nil,
		module.WithRetWrapper(module.HexUint64RetWrapper),
	)
}
