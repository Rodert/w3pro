package txpool_test

import (
	"testing"

	"github.com/rodert/w3pro/module/txpool"
	"github.com/rodert/w3pro/rpctest"
)

func TestStatus(t *testing.T) {
	tests := []rpctest.TestCase[txpool.StatusResponse]{
		{
			Golden:  "status",
			Call:    txpool.Status(),
			WantRet: txpool.StatusResponse{Pending: 10, Queued: 7},
		},
	}

	rpctest.RunTestCases(t, tests)
}
