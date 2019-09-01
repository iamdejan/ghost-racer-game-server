package modelimpl

import (
	"fmt"
	"github.com/iamdejan/ghost-racer-game-server/internal/utility"
	"testing"
)

var p *player = &player{
	playerID:                  1,
	name:                      "Dejan",
	lapsCompleted:             10,
	checkpointsCompleted:      0.75,
	latestCheckpointTimestamp: 1567310552771932193,
}

func TestPlayer_PlayerID(t *testing.T) {
	utility.AssertEquals(p.PlayerID(), p.playerID, "Player ID should be " + fmt.Sprint(p.playerID) + "!", t)
}

func TestPlayer_Name(t *testing.T) {
	utility.AssertEquals(p.Name(), p.name, "Name should be " + p.name + "!", t)
}

func TestPlayer_LapsCompleted(t *testing.T) {
	utility.AssertEquals(p.LapsCompleted(), p.lapsCompleted, "Laps completed should be " + fmt.Sprint(p.lapsCompleted) + "!", t)
}

func TestPlayer_CheckpointsCompleted(t *testing.T) {
	utility.AssertEquals(p.CheckpointsCompleted(), p.checkpointsCompleted, "Checkpoints completed should be " + fmt.Sprint(p.checkpointsCompleted) + "!", t)
}

func TestPlayer_LatestCheckpointTimestamp(t *testing.T) {
	utility.AssertEquals(p.LatestCheckpointTimestamp(), p.latestCheckpointTimestamp, "Latest checkpoint timestamp should be " + fmt.Sprint(p.latestCheckpointTimestamp) + "!", t)
}