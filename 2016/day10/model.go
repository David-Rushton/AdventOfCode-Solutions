package main

type input struct {
	botId int
	chip  chip
}

type instruction struct {
	fromBotId    int
	targetHighId int
	targetLowId  int
	isHightBot   bool
	isLowBot     bool
}

type chip int

type bot struct {
	chips []chip
}

type bin struct {
	botId int
	chip  chip
}

func (b *bot) addChip(value chip) {
	if len(b.chips) >= 2 {
		panic("Cannot add chip to bot.  Bots cannot contain more than 2 chips.")
	}

	b.chips = append(b.chips, value)
}

func (b *bot) takeChips() (high, low chip) {
	if len(b.chips) != 2 {
		panic("Cannot take chip from bot.  Bots must have 2 chips to proceed.")
	}

	if b.chips[0] > b.chips[1] {
		high = b.chips[0]
		low = b.chips[1]
	} else {
		low = b.chips[0]
		high = b.chips[1]
	}

	b.chips = []chip{}

	return high, low
}
