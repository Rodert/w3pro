package web3

import (
	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3protypes"
)

// ClientVersion requests the endpoints client version.
func ClientVersion() w3protypes.CallerFactory[string] {
	return module.NewFactory[string](
		"web3_clientVersion",
		nil,
	)
}
