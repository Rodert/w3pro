package web3

import (
	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3types"
)

// ClientVersion requests the endpoints client version.
func ClientVersion() w3types.CallerFactory[string] {
	return module.NewFactory[string](
		"web3_clientVersion",
		nil,
	)
}
