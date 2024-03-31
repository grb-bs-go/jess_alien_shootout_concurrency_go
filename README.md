# jess_alien_shootout_concurrency_go
Example Golang program demo'ing Go Routines, shared data and concurrent execution.
Scenario is based on an Alien 'Shootout' where a number of (command-line input) 20HP aliens shoot at each other, doing 1d7 damage until all other dead. An Alien can not shoot itself nor a dead alien. Each Alien must shoot at another randomly selected alien, then pause for 1 second before attempting another shot. 
In this implementation, each GoRoutine represents an alien in this implementation, concurrently processing same data set via a sync mutex (no channels in this impl).

Data Structure:
type Alien struct {
	Name      string
	Alive     bool
	HitPoints int
	Damage    int
	Hits      int
	ShootTime int64
}

type BattleTurn struct {
	sync.Mutex
}

Usage:
PS C:\aaaGo2ConcMkII> go run main.go
Usage: main.go <No of aliens>
exit status 1

Example Run:

PS C:\aaaGo2ConcMkII> go run main.go 7

Build GO PATH: buildlogger.log

Start 2024-03-31 16:33:12.3142959 +0000 GMT m=+0.006021401

Args [C:\Users\gavin\AppData\Local\Temp\go-build1903326856\b001\exe\main.exe 7]

7 battle commencing

Alien got lock  6

Alien 6 targets enemy alien 3

7 battle commencing

7 battle commencing

7 battle commencing

7 battle commencing

7 battle commencing

Alien STATUS 0 {Alien1 true 20 0 0 0}

Alien STATUS 1 {Alien2 true 20 0 0 0}

Alien STATUS 2 {Alien3 true 20 0 0 0}

Alien STATUS 3 {Alien4 true 20 4 1 0}

Alien STATUS 4 {Alien5 true 20 0 0 0}

Alien STATUS 5 {Alien6 true 20 0 0 0}

Alien STATUS 6 {Alien7 true 20 0 0 1711902792}

7 battle commencing

Alien got lock  4

Alien 4 targets enemy alien 0

Alien STATUS 0 {Alien1 true 20 1 1 0}

Alien STATUS 1 {Alien2 true 20 0 0 0}

Alien STATUS 2 {Alien3 true 20 0 0 0}

Alien STATUS 3 {Alien4 true 20 4 1 0}

Alien STATUS 4 {Alien5 true 20 0 0 1711902792}

Alien STATUS 5 {Alien6 true 20 0 0 0}

Alien STATUS 6 {Alien7 true 20 0 0 1711902792}

Alien got lock  0

Alien 0 targets enemy alien 1

Alien STATUS 0 {Alien1 true 20 1 1 1711902792}

Alien STATUS 1 {Alien2 true 20 4 1 0}

Alien STATUS 2 {Alien3 true 20 0 0 0}

Alien STATUS 3 {Alien4 true 20 4 1 0}

Alien STATUS 4 {Alien5 true 20 0 0 1711902792}

Alien STATUS 5 {Alien6 true 20 0 0 0}

Alien STATUS 6 {Alien7 true 20 0 0 1711902792}

Alien got lock  3

Alien 3 targets enemy alien 4

Alien STATUS 0 {Alien1 true 20 1 1 1711902792}

Alien STATUS 1 {Alien2 true 20 4 1 0}

Alien STATUS 2 {Alien3 true 20 0 0 0}

Alien STATUS 3 {Alien4 true 20 4 1 1711902792}

Alien STATUS 4 {Alien5 true 20 4 1 1711902792}

Alien STATUS 5 {Alien6 true 20 0 0 0}

Alien STATUS 6 {Alien7 true 20 0 0 1711902792}

Alien got lock  2

Alien 2 targets enemy alien 6

Alien STATUS 0 {Alien1 true 20 1 1 1711902792}

Alien STATUS 1 {Alien2 true 20 4 1 0}

Alien STATUS 2 {Alien3 true 20 0 0 1711902792}

Alien STATUS 3 {Alien4 true 20 4 1 1711902792}

Alien STATUS 4 {Alien5 true 20 4 1 1711902792}

Alien STATUS 5 {Alien6 true 20 0 0 0}

Alien STATUS 6 {Alien7 true 20 6 1 1711902792}

Alien got lock  5

Alien 5 targets enemy alien 4

Alien STATUS 0 {Alien1 true 20 1 1 1711902792}

Alien STATUS 1 {Alien2 true 20 4 1 0}

Alien STATUS 2 {Alien3 true 20 0 0 1711902792}

Alien STATUS 3 {Alien4 true 20 4 1 1711902792}

Alien STATUS 4 {Alien5 true 20 4 2 1711902792}

Alien STATUS 5 {Alien6 true 20 0 0 1711902792}

Alien STATUS 6 {Alien7 true 20 6 1 1711902792}

Alien got lock  1

Alien 1 targets enemy alien 3

Alien STATUS 0 {Alien1 true 20 1 1 1711902792}

Alien STATUS 1 {Alien2 true 20 4 1 1711902792}

Alien STATUS 2 {Alien3 true 20 0 0 1711902792}

Alien STATUS 3 {Alien4 true 20 4 2 1711902792}

Alien STATUS 4 {Alien5 true 20 4 2 1711902792}

Alien STATUS 5 {Alien6 true 20 0 0 1711902792}

Alien STATUS 6 {Alien7 true 20 6 1 1711902792}

Alien got lock  1

Alien 1 targets enemy alien 6

Alien STATUS 0 {Alien1 true 20 1 1 1711902792}

Alien STATUS 1 {Alien2 true 20 4 1 1711902794}

Alien STATUS 2 {Alien3 true 20 0 0 1711902792}

Alien STATUS 3 {Alien4 true 20 4 2 1711902792}

Alien STATUS 4 {Alien5 true 20 4 2 1711902792}

Alien STATUS 5 {Alien6 true 20 0 0 1711902792}

Alien STATUS 6 {Alien7 true 20 11 2 1711902792}

Alien got lock  4

Alien 4 targets enemy alien 2

Alien STATUS 0 {Alien1 true 20 1 1 1711902792}

Alien STATUS 1 {Alien2 true 20 4 1 1711902794}

Alien STATUS 2 {Alien3 true 20 6 1 1711902792}

Alien STATUS 3 {Alien4 true 20 4 2 1711902792}

Alien STATUS 4 {Alien5 true 20 4 2 1711902794}

Alien STATUS 5 {Alien6 true 20 0 0 1711902792}

Alien STATUS 6 {Alien7 true 20 11 2 1711902792}

Alien got lock  0

Alien 0 targets enemy alien 6

Alien STATUS 0 {Alien1 true 20 1 1 1711902794}

Alien STATUS 1 {Alien2 true 20 4 1 1711902794}

Alien STATUS 2 {Alien3 true 20 6 1 1711902792}

Alien STATUS 3 {Alien4 true 20 4 2 1711902792}

Alien STATUS 4 {Alien5 true 20 4 2 1711902794}

Alien STATUS 5 {Alien6 true 20 0 0 1711902792}

Alien STATUS 6 {Alien7 true 20 15 3 1711902792}

Alien got lock  5

Alien 5 targets enemy alien 1

Alien STATUS 0 {Alien1 true 20 1 1 1711902794}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 6 1 1711902792}

Alien STATUS 3 {Alien4 true 20 4 2 1711902792}

Alien STATUS 4 {Alien5 true 20 4 2 1711902794}

Alien STATUS 5 {Alien6 true 20 0 0 1711902794}

Alien STATUS 6 {Alien7 true 20 15 3 1711902792}

Alien got lock  3

Alien 3 targets enemy alien 2

Alien STATUS 0 {Alien1 true 20 1 1 1711902794}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 6 2 1711902792}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 true 20 4 2 1711902794}

Alien STATUS 5 {Alien6 true 20 0 0 1711902794}

Alien STATUS 6 {Alien7 true 20 15 3 1711902792}

Alien got lock  2

Alien 2 targets enemy alien 0

Alien STATUS 0 {Alien1 true 20 3 2 1711902794}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 6 2 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 true 20 4 2 1711902794}

Alien STATUS 5 {Alien6 true 20 0 0 1711902794}

Alien STATUS 6 {Alien7 true 20 15 3 1711902792}

Alien got lock  6

Alien 6 targets enemy alien 4

Alien STATUS 0 {Alien1 true 20 3 2 1711902794}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 6 2 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 true 20 6 3 1711902794}

Alien STATUS 5 {Alien6 true 20 0 0 1711902794}

Alien STATUS 6 {Alien7 true 20 15 3 1711902794}

Alien got lock  4

Alien 4 targets enemy alien 5

Alien STATUS 0 {Alien1 true 20 3 2 1711902794}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 6 2 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 true 20 6 3 1711902796}

Alien STATUS 5 {Alien6 true 20 1 1 1711902794}

Alien STATUS 6 {Alien7 true 20 15 3 1711902794}

Alien got lock  6

Alien 6 targets enemy alien 4

Alien STATUS 0 {Alien1 true 20 3 2 1711902794}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 6 2 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 true 20 11 4 1711902796}

Alien STATUS 5 {Alien6 true 20 1 1 1711902794}

Alien STATUS 6 {Alien7 true 20 15 3 1711902796}

Alien got lock  0

Alien 0 targets enemy alien 2

Alien STATUS 0 {Alien1 true 20 3 2 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 10 3 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 true 20 11 4 1711902796}

Alien STATUS 5 {Alien6 true 20 1 1 1711902794}

Alien STATUS 6 {Alien7 true 20 15 3 1711902796}

Alien got lock  5

Alien 5 targets enemy alien 4

Alien STATUS 0 {Alien1 true 20 3 2 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 10 3 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 true 20 17 5 1711902796}

Alien STATUS 5 {Alien6 true 20 1 1 1711902796}

Alien STATUS 6 {Alien7 true 20 15 3 1711902796}

Alien got lock  1

Alien 1 targets enemy alien 4

Alien5 {Alien5 false 20 20 6 1711902796} is dead!

Alien STATUS 0 {Alien1 true 20 3 2 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902794}

Alien STATUS 2 {Alien3 true 20 10 3 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 1 1 1711902796}

Alien STATUS 6 {Alien7 true 20 15 3 1711902796}

Alien got lock  1

Alien 1 targets enemy alien 5

Alien STATUS 0 {Alien1 true 20 3 2 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902796}

Alien STATUS 2 {Alien3 true 20 10 3 1711902794}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 3 2 1711902796}

Alien STATUS 6 {Alien7 true 20 15 3 1711902796}

Alien got lock  2

Alien 2 targets enemy alien 0

Alien STATUS 0 {Alien1 true 20 4 3 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902796}

Alien STATUS 2 {Alien3 true 20 10 3 1711902796}

Alien STATUS 3 {Alien4 true 20 4 2 1711902794}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 3 2 1711902796}

Alien STATUS 6 {Alien7 true 20 15 3 1711902796}

Alien got lock  3

Alien 3 targets enemy alien 5

Alien STATUS 0 {Alien1 true 20 4 3 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902796}

Alien STATUS 2 {Alien3 true 20 10 3 1711902796}

Alien STATUS 3 {Alien4 true 20 4 2 1711902796}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 3 3 1711902796}

Alien STATUS 6 {Alien7 true 20 15 3 1711902796}

Alien got lock  3

Alien 3 targets enemy alien 6

Alien STATUS 0 {Alien1 true 20 4 3 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902796}

Alien STATUS 2 {Alien3 true 20 10 3 1711902796}

Alien STATUS 3 {Alien4 true 20 4 2 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 3 3 1711902796}

Alien STATUS 6 {Alien7 true 20 15 4 1711902796}

Alien got lock  6

Alien 6 targets enemy alien 5

Alien STATUS 0 {Alien1 true 20 4 3 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902796}

Alien STATUS 2 {Alien3 true 20 10 3 1711902796}

Alien STATUS 3 {Alien4 true 20 4 2 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 9 4 1711902796}

Alien STATUS 6 {Alien7 true 20 15 4 1711902798}

Alien got lock  5

Alien 5 targets enemy alien 0

Alien STATUS 0 {Alien1 true 20 7 4 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902796}

Alien STATUS 2 {Alien3 true 20 10 3 1711902796}

Alien STATUS 3 {Alien4 true 20 4 2 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 9 4 1711902798}

Alien STATUS 6 {Alien7 true 20 15 4 1711902798}

Alien got lock  1

Alien 1 targets enemy alien 6

Alien STATUS 0 {Alien1 true 20 7 4 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902798}

Alien STATUS 2 {Alien3 true 20 10 3 1711902796}

Alien STATUS 3 {Alien4 true 20 4 2 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 9 4 1711902798}

Alien STATUS 6 {Alien7 true 20 19 5 1711902798}

Alien got lock  4

Alien got lock  2

Alien 2 targets enemy alien 3

Alien STATUS 0 {Alien1 true 20 7 4 1711902796}

Alien STATUS 1 {Alien2 true 20 9 2 1711902798}

Alien STATUS 2 {Alien3 true 20 10 3 1711902798}

Alien STATUS 3 {Alien4 true 20 9 3 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 9 4 1711902798}

Alien STATUS 6 {Alien7 true 20 19 5 1711902798}

Alien got lock  0

Alien 0 targets enemy alien 5

Alien STATUS 0 {Alien1 true 20 7 4 1711902798}

Alien STATUS 1 {Alien2 true 20 9 2 1711902798}

Alien STATUS 2 {Alien3 true 20 10 3 1711902798}

Alien STATUS 3 {Alien4 true 20 9 3 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902798}

Alien STATUS 6 {Alien7 true 20 19 5 1711902798}

Alien got lock  2

Alien 2 targets enemy alien 1

Alien STATUS 0 {Alien1 true 20 7 4 1711902798}

Alien STATUS 1 {Alien2 true 20 10 3 1711902798}

Alien STATUS 2 {Alien3 true 20 10 3 1711902800}

Alien STATUS 3 {Alien4 true 20 9 3 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902798}

Alien STATUS 6 {Alien7 true 20 19 5 1711902798}

Alien got lock  6

Alien 6 targets enemy alien 2

Alien STATUS 0 {Alien1 true 20 7 4 1711902798}

Alien STATUS 1 {Alien2 true 20 10 3 1711902798}

Alien STATUS 2 {Alien3 true 20 13 4 1711902800}

Alien STATUS 3 {Alien4 true 20 9 3 1711902798}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902798}

Alien STATUS 6 {Alien7 true 20 19 5 1711902800}

Alien got lock  3

Alien 3 targets enemy alien 1

Alien STATUS 0 {Alien1 true 20 7 4 1711902798}

Alien STATUS 1 {Alien2 true 20 11 4 1711902798}

Alien STATUS 2 {Alien3 true 20 13 4 1711902800}

Alien STATUS 3 {Alien4 true 20 9 3 1711902800}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902798}

Alien STATUS 6 {Alien7 true 20 19 5 1711902800}

Alien got lock  5

Alien 5 targets enemy alien 1

Alien STATUS 0 {Alien1 true 20 7 4 1711902798}

Alien STATUS 1 {Alien2 true 20 14 5 1711902798}

Alien STATUS 2 {Alien3 true 20 13 4 1711902800}

Alien STATUS 3 {Alien4 true 20 9 3 1711902800}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902800}

Alien STATUS 6 {Alien7 true 20 19 5 1711902800}

Alien got lock  0

Alien 0 targets enemy alien 2

Alien STATUS 0 {Alien1 true 20 7 4 1711902800}

Alien STATUS 1 {Alien2 true 20 14 5 1711902798}

Alien STATUS 2 {Alien3 true 20 14 5 1711902800}

Alien STATUS 3 {Alien4 true 20 9 3 1711902800}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902800}

Alien STATUS 6 {Alien7 true 20 19 5 1711902800}

Alien got lock  1

Alien 1 targets enemy alien 0

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 true 20 14 5 1711902800}

Alien STATUS 2 {Alien3 true 20 14 5 1711902800}

Alien STATUS 3 {Alien4 true 20 9 3 1711902800}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902800}

Alien STATUS 6 {Alien7 true 20 19 5 1711902800}

Alien got lock  5

Alien 5 targets enemy alien 3

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 true 20 14 5 1711902800}

Alien STATUS 2 {Alien3 true 20 14 5 1711902800}

Alien STATUS 3 {Alien4 true 20 11 4 1711902800}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902800}

Alien got lock  6

Alien 6 targets enemy alien 1

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 true 20 14 6 1711902800}

Alien STATUS 2 {Alien3 true 20 14 5 1711902800}

Alien STATUS 3 {Alien4 true 20 11 4 1711902800}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902802}

Alien got lock  3

Alien 3 targets enemy alien 1

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 true 20 19 7 1711902800}

Alien STATUS 2 {Alien3 true 20 14 5 1711902800}

Alien STATUS 3 {Alien4 true 20 11 4 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902802}

Alien got lock  0

Alien 0 targets enemy alien 1

Alien2 {Alien2 false 20 22 8 1711902800} is dead!

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 false 20 22 8 1711902800}

Alien STATUS 2 {Alien3 true 20 14 5 1711902800}

Alien STATUS 3 {Alien4 true 20 11 4 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 true 20 15 5 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902802}

Alien got lock  0

Alien 0 targets enemy alien 5

Alien6 {Alien6 false 20 20 6 1711902802} is dead!

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 false 20 22 8 1711902800}

Alien STATUS 2 {Alien3 true 20 14 5 1711902800}

Alien STATUS 3 {Alien4 true 20 11 4 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 false 20 20 6 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902802}

Alien got lock  1

Alien got lock  2

Alien 2 targets enemy alien 3

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 false 20 22 8 1711902800}

Alien STATUS 2 {Alien3 true 20 14 5 1711902802}

Alien STATUS 3 {Alien4 true 20 14 5 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 false 20 20 6 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902802}

Alien got lock  0

Alien 0 targets enemy alien 2

Alien3 {Alien3 false 20 20 6 1711902802} is dead!

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 false 20 22 8 1711902800}

Alien STATUS 2 {Alien3 false 20 20 6 1711902802}

Alien STATUS 3 {Alien4 true 20 14 5 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 false 20 20 6 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902802}

Alien got lock  0

Alien 0 targets enemy alien 3

Alien4 {Alien4 false 20 20 6 1711902802} is dead!

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 false 20 22 8 1711902800}

Alien STATUS 2 {Alien3 false 20 20 6 1711902802}

Alien STATUS 3 {Alien4 false 20 20 6 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 false 20 20 6 1711902802}

Alien STATUS 6 {Alien7 true 20 19 5 1711902802}

Alien got lock  0

Alien 0 targets enemy alien 6

Alien7 {Alien7 false 20 23 6 1711902802} is dead!

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 false 20 22 8 1711902800}

Alien STATUS 2 {Alien3 false 20 20 6 1711902802}

Alien STATUS 3 {Alien4 false 20 20 6 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 false 20 20 6 1711902802}

Alien STATUS 6 {Alien7 false 20 23 6 1711902802}

Alien got lock  0

FINAL BATTLE STATUS {Alien1 true 20 8 5 1711902800}

Alien STATUS 0 {Alien1 true 20 8 5 1711902800}

Alien STATUS 1 {Alien2 false 20 22 8 1711902800}

Alien STATUS 2 {Alien3 false 20 20 6 1711902802}

Alien STATUS 3 {Alien4 false 20 20 6 1711902802}

Alien STATUS 4 {Alien5 false 20 20 6 1711902796}

Alien STATUS 5 {Alien6 false 20 20 6 1711902802}

Alien STATUS 6 {Alien7 false 20 23 6 1711902802}

************************WINNER is  {Alien1 true 20 8 5 1711902800}
...
