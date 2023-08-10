package txpool

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rodert/w3pro/internal/module"
	"github.com/rodert/w3pro/w3protypes"
)

// Status requests the number of pending and queued transactions in the transaction pool.
func Status() w3protypes.CallerFactory[StatusResponse] {
	return module.NewFactory[StatusResponse](
		"txpool_status",
		nil,
	)
}

type StatusResponse struct {
	Pending uint
	Queued  uint
}

func (s *StatusResponse) UnmarshalJSON(data []byte) error {
	type statusResponse struct {
		Pending hexutil.Uint `json:"pending"`
		Queued  hexutil.Uint `json:"queued"`
	}

	var dec statusResponse
	if err := json.Unmarshal(data, &dec); err != nil {
		return err
	}
	s.Pending = uint(dec.Pending)
	s.Queued = uint(dec.Queued)
	return nil
}
