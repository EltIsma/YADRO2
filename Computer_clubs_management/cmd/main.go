package main

import (
	"bufio"
	"club_control/internal/parsing"
	"fmt"
	"log/slog"
	"os"
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


// func consumer(ctx context.Context, wg *sync.WaitGroup, eventCh <-chan Event, cfg *Config) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case event, ok := <-eventCh:
// 			if !ok {
// 				handleClosingClub(cfg)
// 				return
// 			}
// 			switch event.ID {
// 			case CLIENTARRIVED:
// 				handleClientArrived(cfg, event)
// 			case CLIENTSIT:
// 				handleClientSit(cfg, event)
// 			case CLIENTWAIT:
// 				handleClientWaiting(cfg, event)
// 			case CLIENTLEFT:
// 				handleClientLeft(cfg, event)
// 			}
// 		case <-ctx.Done():
// 			//handleClosingClub(cfg)
// 			return
// 		}
// 	}
// }

func main() {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		logger.Error("can't open file", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	config, badStr, err := parsing.ParseConfig(scanner)
	if err != nil {
		logger.Error("bad configuration", slog.String("error", err.Error()))
		fmt.Println(badStr)
		os.Exit(1)
	}

	config.PrintOpenTime()
	var seqCheck time.Time
	//eventCh := make(chan Event)
	// ctx, cancel := context.WithCancel(context.Background())
	// wg := sync.WaitGroup{}
	// wg.Add(1)

	// go func() {
	// 	defer func(){
	// 		wg.Done()
	//         close(eventCh)
	// 	}()
	// 	for scanner.Scan() {
	// 		event, line, err := parseEvent(scanner.Text(), config.numTables, seqCheck)
	// 		if err != nil {
	// 			fmt.Println(line)
	// 			logger.Error("invalid input data", slog.String("error", err.Error()))
	// 			cancel()
	// 			return
	// 		}
	// 		seqCheck = event.Time
	// 		eventCh <- *event
	// 	}
	// }()

	// wg.Add(1)
	// go consumer(ctx, &wg, eventCh, config)
	// wg.Wait()
	for scanner.Scan() {
		event, line, err := parsing.ParseEvent(scanner.Text(), config.GetNumOfTables(), seqCheck)
		if err != nil {
			fmt.Println(line)
			logger.Error("invalid input data", slog.String("error", err.Error()))
			os.Exit(1)
		}
		seqCheck = event.Time
		switch event.ID {
		case CLIENTARRIVED:
			config.HandleClientArrived(*event)
		case CLIENTSIT:
			config.HandleClientSit(*event)
		case CLIENTWAIT:
			config.HandleClientWaiting(*event)
		case CLIENTLEFT:
			config.HandleClientLeft(*event)
		}
	}
	config.HandleClosingClub()

}
