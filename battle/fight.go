package battle

import (
	"conc2/data"
	"conc2/logger"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const damageAmountSeed int = 7

func Fight(alienIndex int, wg *sync.WaitGroup, battleTurn *data.BattleTurn, aliens []*data.Alien) []*data.Alien {

	logger.Log.Println(len(aliens), "battling")
	fmt.Println(len(aliens), "battle commencing")

	for {
		survivorCount := len(aliens)
		if aliens[alienIndex].ShootTime == 0 || time.Now().Unix()-aliens[alienIndex].ShootTime > 1 {
			// Alien needs mutex lock to take a turn
			battleTurn.Lock()
			fmt.Println("Alien got lock ", alienIndex)
			if !aliens[alienIndex].Alive {
				// GoR Done when Alien is dead
				battleTurn.Unlock()
				wg.Done()
				return aliens
			} else {
				// Check if only survivor
				for _, v := range aliens {
					if !v.Alive {
						survivorCount--
					}
				}
				//fmt.Println("SURVIVER COUNT", survivorCount)
				if survivorCount == 1 {
					fmt.Println("FINAL BATTLE STATUS", *aliens[alienIndex])
					for i, v := range aliens {
						fmt.Println("Alien STATUS", i, *v)
					}
					fmt.Println("************************WINNER is ", *aliens[alienIndex])
					battleTurn.Unlock()
					wg.Done()
					return aliens
				}
			}

			var targetIndex = selectTargetIndex(aliens, alienIndex)
			fmt.Println("Alien", alienIndex, "targets enemy alien", targetIndex)
			// Shoot, add damage
			aliens[targetIndex].Damage = aliens[targetIndex].Damage + shotDamageAmount()
			aliens[targetIndex].Hits = aliens[targetIndex].Hits + 1
			if aliens[targetIndex].Damage >= aliens[targetIndex].HitPoints {
				aliens[targetIndex].Alive = false
				fmt.Println(aliens[targetIndex].Name, *aliens[targetIndex], "is dead!")
				logger.Log.Println(aliens[targetIndex].Name, *aliens[targetIndex], "is dead!")
			} else {
				// Update Shoot Time
				aliens[alienIndex].ShootTime = time.Now().Unix()
			}
			for i, v := range aliens {
				fmt.Println("Alien STATUS", i, *v)
			}
			battleTurn.Unlock()
		}
	} // close for loop
}

// Called when two or more aliens alive
func selectTargetIndex(aliens []*data.Alien, currentAlien int) int {
	targetIndex := 0
	length := len(aliens)

	// generate a random integer
	for {
		randomInt := rand.Intn(length)
		//fmt.Println("Random number generated", randomInt)
		// Remove invalid targets
		if randomInt == currentAlien || !aliens[randomInt].Alive {
			// cant shoot yourself or dead alien
			continue
		}
		targetIndex = randomInt
		break
	} // close for loop
	return targetIndex
}

func shotDamageAmount() int {
	return rand.Intn(damageAmountSeed)
}
