package grouping

import (
	"time"

	"github.com/PRAgarawal/eet/eet"
	"github.com/go-kit/kit/log"
)

const NotificationInterval time.Duration = 7 * 24 * time.Hour

const DayTick = int(time.Friday)
const CollectionHourTick = 10
const LunchHourTick = 11
const MinuteTick = 30
const SecondTick = 00

type Grouper struct {
	repo   eet.Repository
	logger log.Logger
}

func NewGrouper(repo eet.Repository, logger log.Logger) *Grouper {
	return &Grouper{
		repo:   repo,
		logger: logger,
	}
}

func (g *Grouper) Start() {
	go g.collectionNotifyRoutine()
	go g.lunchNotifyRoutine()
}

func (g *Grouper) lunchNotifyRoutine() {
	ticker := g.newTicker(LunchHourTick)
	for {
		<-ticker.C
		g.logger.Log("message", "notifying lunch groups")
		ticker = g.newTicker(LunchHourTick)
	}
}

func (g *Grouper) collectionNotifyRoutine() {
	ticker := g.newTicker(CollectionHourTick)
	for {
		<-ticker.C
		g.logger.Log("message", "notifying lunch groups")
		ticker = g.newTicker(CollectionHourTick)
	}
}

func (g *Grouper) newTicker(hourTick int) *time.Ticker {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(), DayTick, hourTick, MinuteTick, SecondTick, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(NotificationInterval)
	}

	g.logger.Log("method", "newTicker", "nextTick", nextTick)
	diff := nextTick.Sub(time.Now())
	return time.NewTicker(diff)
}
