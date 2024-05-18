package parsing

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"club_control/internal/domain"
)

func ParseConfig(bs *bufio.Scanner) (*domain.Config, string, error) {
	bs.Scan()
	numTables, err := strconv.Atoi(bs.Text())
	if err != nil || numTables <= 0 {
		return nil, bs.Text(), errors.New("invalid input: invalid number of tables")
	}
	var startTime, endTime time.Time
	bs.Scan()
	times := strings.Split(bs.Text(), " ")
	if len(times) != 2 {
		return nil, bs.Text(), errors.New("invalid input: line must contain open time and close time")
	}

	startTime, err = time.Parse("15:04", times[0])
	if err != nil {
		return nil, fmt.Sprintf("%s %s", times[0], times[1]), err
	}

	endTime, err = time.Parse("15:04", times[1])
	if err != nil {
		return nil, fmt.Sprintf("%s %s", times[0], times[1]), err
	}

	bs.Scan()
	price, err := strconv.Atoi(bs.Text())
	if err != nil || price <= 0 {
		return nil, bs.Text(), fmt.Errorf("invalid input: invalid price")
	}

	return domain.NewClub(numTables, startTime, endTime, price), "", nil
}
