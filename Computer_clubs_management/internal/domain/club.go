package domain

import (
	"fmt"
	"sort"
	"time"
)

const (
	CLIENTARRIVED = 1
	CLIENTSIT     = 2
	CLIENTWAIT    = 3
	CLIENTLEFT    = 4
	OUTCLIENTLEFT = 11
	OUTCLIENTSAT  = 12
	ERROREVENT    = 13
)

const (
	NotOpenYet       = "NotOpenYet"
	YouShallNotPass  = "YouShallNotPass"
	PlaceIsBusy      = "PlaceIsBusy"
	ClientUnknown    = "ClientUnknown"
	ICanWaitNoLonger = "ICanWaitNoLonger!"
)

type Config struct {
	numTables int
	tables    map[int]*Table
	openTime  time.Time
	closeTime time.Time
	price     int
	client    map[string]int
	queue     []string
}

func NewClub(numTables int, oT time.Time, cT time.Time, p int) *Config {
	return &Config{
		numTables: numTables,
		tables:    make(map[int]*Table, numTables),
		openTime:  oT,
		closeTime: cT,
		price:     p,
		client:    make(map[string]int),
		queue:     make([]string, 0, numTables),
	}
}

func (cfg *Config) GetNumOfTables() int {
	return cfg.numTables
}

func (cfg *Config) PrintOpenTime() {
	fmt.Println(cfg.openTime.Format("15:04"))
}

func (cfg *Config) HandleClientArrived(event Event) {
	fmt.Println(event.Time.Format("15:04"), event.ID, event.Client)
	if event.Time.Before(cfg.openTime) || event.Time.After(cfg.closeTime) {
		fmt.Println(event.Time.Format("15:04"), ERROREVENT, NotOpenYet)
		return
	}

	if _, ok := cfg.client[event.Client]; ok {
		fmt.Println(event.Time.Format("15:04"), ERROREVENT, YouShallNotPass)
		return
	}
	cfg.client[event.Client] = withoutTable

}

func (cfg *Config) HandleClientSit(event Event) {
	fmt.Println(event.Time.Format("15:04"), event.ID, event.Client, event.Table)
	if _, ok := cfg.client[event.Client]; !ok {
		fmt.Println(event.Time.Format("15:04"), ERROREVENT, ClientUnknown)
		return
	}
	if _, ok := cfg.tables[event.Table]; ok {
		if cfg.tables[event.Table].busy {
			fmt.Println(event.Time.Format("15:04"), ERROREVENT, PlaceIsBusy)
			return
		}
	} else {
		cfg.tables[event.Table] = &Table{}
	}
	if cfg.client[event.Client] != withoutTable {
		cfg.tables[cfg.client[event.Client]].revenueCalculation(event.Time, cfg.price)

		//duration := event.Time.Sub(cfg.tables[cfg.client[event.Client]].start)
		cfg.tables[cfg.client[event.Client]].busy = false
		//currentTime := cfg.tables[cfg.client[event.Client]].time
		//cfg.tables[cfg.client[event.Client]].time = currentTime.Add(duration)
		//cfg.tables[cfg.client[event.Client]].revenue += int(math.Ceil(float64(duration)/float64(time.Hour))) * cfg.price
	}

	cfg.tables[event.Table].busy = true
	cfg.tables[event.Table].start = event.Time
	cfg.client[event.Client] = event.Table

}

func (cfg *Config) HandleClientWaiting(event Event) {
	fmt.Println(event.Time.Format("15:04"), event.ID, event.Client)
	if len(cfg.tables) != cfg.numTables {
		fmt.Println(event.Time.Format("15:04"), ERROREVENT, ICanWaitNoLonger)
		return
	}
	for _, table := range cfg.tables {
		if !table.busy {
			fmt.Println(event.Time.Format("15:04"), ERROREVENT, ICanWaitNoLonger)
			return
		}
	}
	//if the client is already sitting, just ignore
	if cfg.client[event.Client] != withoutTable {
		return
	}

	for _, cn := range cfg.queue {
		if cn == event.Client {
			return
		}
	}

	if len(cfg.queue) == cfg.numTables {
		fmt.Println(event.Time.Format("15:04"), OUTCLIENTLEFT, event.Client)
		delete(cfg.client, event.Client)
		return
	}

	cfg.queue = append(cfg.queue, event.Client)

}

func (cfg *Config) HandleClientLeft(event Event) {
	fmt.Println(event.Time.Format("15:04"), event.ID, event.Client)
	if _, ok := cfg.client[event.Client]; !ok {
		fmt.Println(event.Time.Format("15:04"), ERROREVENT, ClientUnknown)
		return
	}
	if cfg.client[event.Client] != withoutTable {
		cfg.tables[cfg.client[event.Client]].revenueCalculation(event.Time, cfg.price)
		//duration := event.Time.Sub(cfg.tables[cfg.client[event.Client]].start)
		cfg.tables[cfg.client[event.Client]].busy = false
		//currentTime := cfg.tables[cfg.client[event.Client]].time
		//cfg.tables[cfg.client[event.Client]].time = currentTime.Add(duration)
		//cfg.tables[cfg.client[event.Client]].revenue += int(math.Ceil(float64(duration)/float64(time.Hour))) * cfg.price
	} else {
		for idx, client := range cfg.queue {
			if client == event.Client {
				cfg.queue = append(cfg.queue[:idx], cfg.queue[idx+1:]...)
			}
		}
	}
	freeTable := cfg.client[event.Client]
	delete(cfg.client, event.Client)
	if len(cfg.queue) != 0 && freeTable != withoutTable {
		client := cfg.queue[0]
		cfg.queue = cfg.queue[1:]
		cfg.HandleClientSitFromQueue(client, freeTable, event.Time)
	}
}

func (cfg *Config) HandleClientSitFromQueue(client string, table int, eventTime time.Time) {
	fmt.Println(eventTime.Format("15:04"), OUTCLIENTSAT, client, table)
	cfg.tables[table].busy = true
	cfg.tables[table].start = eventTime
	cfg.client[client] = table
}

func (cfg *Config) HandleClosingClub() {
	names := make([]string, 0)
	tables := make([]int, 0)
	for n := range cfg.client {
		names = append(names, n)
	}
	for t := range cfg.tables {
		tables = append(tables, t)
	}
	sort.Strings(names)
	sort.Ints(tables)
	for _, name := range names {
		fmt.Println(cfg.closeTime.Format("15:04"), OUTCLIENTLEFT, name)
	}

	for client, tableId := range cfg.client {
		if tableId != withoutTable {
			cfg.tables[tableId].revenueCalculation(cfg.closeTime, cfg.price)
			cfg.tables[tableId].busy = false
		}
		cfg.queue = []string{}
		delete(cfg.client, client)
	}
	fmt.Println(cfg.closeTime.Format("15:04"))
	it := 1
	for _, tableInd := range tables {
		cfg.tables[tableInd].start = time.Time{}
		fmt.Println(tableInd, cfg.tables[tableInd].revenue, cfg.tables[tableInd].time.Format("15:04"))
		it++

	}

	for it <= cfg.numTables {
		fmt.Println(it, 0, time.Time{}.Format("00:00"))
		it++
	}
}
