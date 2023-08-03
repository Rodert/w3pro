package web3

import (
	"github.com/lmittmann/w3/w3types"
	"github.com/rodert/w3pro/internal/module"
)

// ClientVersion requests the endpoints client version.
func ClientVersion() w3types.CallerFactory[string] {
	return module.NewFactory[string](
		"web3_clientVersion",
		nil,
	)
}
