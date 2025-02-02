package main

type Player struct {
	hitPoints   int
	damage      int
	armour      int
	manaTotal   int
	manaSpent   int
	castHistory string
}

func newPlayer(hitPoints, manaTotal int) Player {
	return Player{
		hitPoints:   hitPoints,
		manaTotal:   manaTotal,
		castHistory: "",
	}
}
