package modelimpl

import (
	"github.com/iamdejan/ghost-racer-game-server/pkg/model"
	"sync"
)

type roomEventFeed struct {
	events []model.RoomEvent
	eventsMutex sync.RWMutex

	waitChannel chan struct{}
}

func newRoomEventFeed() *roomEventFeed {
	return &roomEventFeed{
		waitChannel: make(chan struct{}),
	}
}

func (feed *roomEventFeed) put(event model.RoomEvent) {
	feed.eventsMutex.Lock()
	defer feed.eventsMutex.Unlock()

	feed.events = append(feed.events, event)
}

func (feed *roomEventFeed) List(lastID int) ([]model.RoomEvent, int, chan struct{}) {
	feed.eventsMutex.Lock()
	defer feed.eventsMutex.Unlock()

	return feed.events[lastID + 1:], len(feed.events) - 1, feed.waitChannel
}