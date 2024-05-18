package parsing

import (
	"club_control/internal/domain"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const CLIENTSIT = 2

func ParseEvent(line string, numTables int, seqCheck time.Time) (*domain.Event, string, error) {
	parts := strings.Split(line, " ")
	if len(parts) < 3 {
		return nil, line, errors.New("the line must contain minimum 3 objects")
	}
	time, err := domain.ValidateTime(parts[0])
	if err != nil {
		return nil, line, fmt.Errorf("invalid time: %s", err.Error())
	}
	if time.Before(seqCheck) && !seqCheck.IsZero() {
		return nil, line, errors.New("incorrect time: events must happen sequentially")
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, line, fmt.Errorf("invalid id: %s", err.Error())
	}
	client := parts[2]
	if !domain.ValidateName(client) {
		return nil, line, errors.New("invalid name")
	}
	table := 0

	if id == CLIENTSIT {
		if len(parts) != 4 {
			return nil, line, errors.New("the line must contain 4 objects")
		}
		table, err = strconv.Atoi(parts[3])
		if err != nil {
			return nil, line, fmt.Errorf("invalid table number: %s", err.Error())
		}
		if table < 1 || table > numTables {
			return nil, line, errors.New("incorrect table number")
		}
	}
	return (domain.NewEvent(&domain.EventOptions{Time: time, ID: id, Client: client, Table: table})), "", nil

}
