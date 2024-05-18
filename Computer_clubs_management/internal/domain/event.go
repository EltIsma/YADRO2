package domain

import (
	"regexp"
	"time"
)

const withoutTable = 0

type Event struct {
	Time   time.Time
	ID     int
	Client string
	Table  int
}

func NewEvent(options *EventOptions) *Event {
	if options.Table == 0 {
		options.Table = withoutTable
	}
	return &Event{
		Time:   options.Time,
		ID:     options.ID,
		Client: options.Client,
		Table:  options.Table,
	}
}

type EventOptions struct {
	Time   time.Time
	ID     int
	Client string
	Table  int
}

func ValidateName(name string) bool {
	re := regexp.MustCompile(`^[a-z0-9_-]+$`)
	return re.MatchString(name)
}

func ValidateTime(tl string) (time.Time, error) {
	t, err := time.Parse("15:04", tl)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
