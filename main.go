package main

import (
	"conc2/battle"
	"conc2/commandline"
	"conc2/data"
	"conc2/logger"
	"fmt"
	"sync"
	"time"
)
var (
	battleTurn       *data.BattleTurn
	aliens           []*data.Alien
	alienCount       int
)

func main() {
	fmt.Println("Start",time.Now())
    battleTurn := new(data.BattleTurn)
	aliens = commandline.SpawnAliens()
	alienCount = len(aliens)
	logger.Log.Println("No Aliens", alienCount)

	// Start Alien GoRoutines
	wg := sync.WaitGroup{}
	for i := 0; i < alienCount; i++ {
		wg.Add(1)
		// each GoR is an alien!
		go battle.Fight(i, &wg, battleTurn, aliens)
	}
	wg.Wait()
}


