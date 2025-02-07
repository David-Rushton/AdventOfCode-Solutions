package main

import (
	"fmt"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2016/internal/aoc"
)

func main() {
	fmt.Println("--- Day 10: Balance Bots ---")
	fmt.Println()

	findHigh := chip(61)
	findLow := chip(17)
	if aoc.TestMode {
		findHigh = chip(5)
		findLow = chip(2)
	}

	inputs, instructions := parse(aoc.GetInput(10))
	targetBot, chipScore := process(inputs, instructions, findHigh, findLow)

	fmt.Println()
	fmt.Printf("Target Bot: %d\n", targetBot)
	fmt.Printf("Chip Score: %d\n", chipScore)
}

func process(inputs []input, instructions []instruction, findHigh, findLow chip) (targetBot, chipScore int) {
	targetBot = 0
	chipScore = 0

	// find the max bot/output id
	var maxId int
	for i := range inputs {
		if inputs[i].botId > maxId {
			maxId = inputs[i].botId
		}
	}

	for i := range instructions {
		if instructions[i].targetHighId > maxId {
			maxId = instructions[i].targetHighId
		}

		if instructions[i].targetLowId > maxId {
			maxId = instructions[i].targetLowId
		}
	}

	// declare bots and output outputs
	bots := make([]bot, maxId+1)
	outputs := make([]bin, maxId+1)

	// init bots
	for _, input := range inputs {
		fmt.Printf(" - Adding chip %d to bot %d\n", input.chip, input.botId)
		bots[input.botId].addChip(input.chip)
	}

	// process chips
	for len(instructions) > 0 {
		instruction := instructions[0]
		instructions = instructions[1:]

		// cannot process yet
		// add to end of queue
		if len(bots[instruction.fromBotId].chips) != 2 {
			instructions = append(instructions, instruction)
			continue
		}

		highId := instruction.targetHighId
		lowId := instruction.targetLowId
		highChip, lowChip := bots[instruction.fromBotId].takeChips()

		if highChip == findHigh && lowChip == findLow {
			fmt.Printf(" - 🏆🏆🏆 bot %v processed chips %v and %v🏆🏆🏆\n", instruction.fromBotId, highChip, lowChip)
			targetBot = instruction.fromBotId
		}

		if instruction.isHightBot {
			fmt.Printf(" - Adding high chip %d to bot %d\n", highChip, highId)
			bots[highId].addChip(highChip)
		} else {
			fmt.Printf(" - Adding high chip %d to output %d\n", highChip, highId)
			outputs[highId] = bin{instruction.fromBotId, highChip}
		}

		if instruction.isLowBot {
			fmt.Printf(" - Adding low chip %d to bot %d\n", lowChip, lowId)
			bots[lowId].addChip(lowChip)
		} else {
			fmt.Printf(" - Adding low chip %d to output %d\n", lowChip, lowId)
			outputs[lowId] = bin{instruction.fromBotId, lowChip}
		}
	}

	// Multiply first 3 bins
	chipScore = int(outputs[0].chip) * int(outputs[1].chip) * int(outputs[2].chip)

	return targetBot, chipScore
}

func parse(values []string) ([]input, []instruction) {
	inputs := []input{}
	instructions := []instruction{}

	for i := range values {
		elements := strings.Split(values[i], " ")

		if elements[0] == "value" {
			inputs = append(inputs, input{
				chip:  chip(aoc.ToInt(elements[1])),
				botId: aoc.ToInt(elements[5])})

			continue
		}

		if elements[0] == "bot" {
			instructions = append(instructions, instruction{
				fromBotId:    aoc.ToInt(elements[1]),
				targetHighId: aoc.ToInt(elements[11]),
				targetLowId:  aoc.ToInt(elements[6]),
				isHightBot:   elements[10] == "bot",
				isLowBot:     elements[5] == "bot"})

			continue
		}

		panic(fmt.Sprintf("Instruction not supported: %v", elements[0]))
	}

	return inputs, instructions
}
