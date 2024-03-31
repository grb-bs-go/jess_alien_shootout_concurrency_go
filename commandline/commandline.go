package commandline
import (
	"conc2/data"
	"fmt"
	"os"
	"strconv"
)
func SpawnAliens() []*data.Alien {
	argsWithProg := os.Args 
    if len(argsWithProg) != 2 {
		fmt.Println("Usage: main.go <No of aliens>")
		os.Exit(1)
	}
    fmt.Println("Args",argsWithProg)
	noAliens, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
			fmt.Println("ERROR: Please input an integer value", err)
	}	
	var aliens []*data.Alien = make([]*data.Alien, noAliens)

	for i := 0; i < noAliens; i++ {
		aliens[i] = &data.Alien{Name: fmt.Sprint("Alien", i+1),Alive: true,HitPoints: 20,Damage: 0,Hits: 0,ShootTime: 0}
	}
	return aliens
}