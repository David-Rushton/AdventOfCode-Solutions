package main

type Boss struct {
	hitPoints int
	damage    int
}

func newBoss(hitPoints, damage int) Boss {
	return Boss{
		hitPoints: hitPoints,
		damage:    damage,
	}
}
