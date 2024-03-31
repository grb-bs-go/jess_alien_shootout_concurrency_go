package data

import "sync"

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