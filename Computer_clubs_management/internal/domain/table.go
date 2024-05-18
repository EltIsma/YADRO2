package domain

import (
	"math"
	"time"
)

type Table struct {
	busy    bool
	time    time.Time
	revenue int
	start   time.Time
}

func (t *Table) revenueCalculation(eventTime time.Time, price int) {
	duration := eventTime.Sub(t.start)
	currentTime := t.time
	t.time = currentTime.Add(duration)
	t.revenue += int(math.Ceil(float64(duration)/float64(time.Hour))) * price
}
