package modelimpl

import (
	"github.com/iamdejan/ghost-racer-game-server/internal/utility"
	"testing"
)

func TestCircuit_ID(t *testing.T) {
	var ID uint64 = 1
	var c = circuit{
		circuitID: ID,
	}

	utility.AssertEquals(ID, c.ID(), "ID isn't equal!", t)
}