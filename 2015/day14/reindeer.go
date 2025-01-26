package main

type reindeer struct {
	Name           string
	FlightSpeed    int
	FlightDuration int
	RestDuration   int

	isFlying  bool
	travelled int
	counter   int
	score     int
}

func (r *reindeer) Advance() {
	if r.counter == 0 {
		r.isFlying = !r.isFlying

		if r.isFlying {
			r.counter = r.FlightDuration
		} else {
			r.counter = r.RestDuration
		}
	}

	if r.isFlying {
		r.travelled += r.FlightSpeed
	}

	r.counter--
}

func (r *reindeer) GetDistanceTravelled() int {
	return r.travelled
}

func (r *reindeer) AddPoint() {
	r.score++
}

func (r *reindeer) GetPoints() int {
	return r.score
}
