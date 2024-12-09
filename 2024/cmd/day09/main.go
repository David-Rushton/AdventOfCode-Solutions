package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
)

type diskFile struct {
	id         int
	usedBlocks int
	freeBlocks int
}

func main() {
	fmt.Println("--- Day 9: Disk Fragmenter ---")
	fmt.Println()

	diskMap := parse(aoc.Input[0])
	blocks := getBlocks(diskMap)
	compacted := compact(blocks, aoc.Star == aoc.StarTwo)
	checksum := getChecksum(compacted)

	fmt.Println()
	fmt.Printf("Result: %d\n", checksum)
}

func getChecksum(blocks []*diskFile) int64 {
	var result int64

	for i, block := range blocks {
		if block == nil {
			continue
		}
		result += int64(i) * int64(block.id)
	}

	return result
}

func compact(blocks []*diskFile, moveWholeFiles bool) []*diskFile {
	if moveWholeFiles {
		return compactFiles(blocks)
	} else {
		return compactBlocks(blocks)
	}
}

func compactFiles(blocks []*diskFile) []*diskFile {
	result := append([]*diskFile{}, blocks...)

	type freeMemory struct {
		at  int
		len int
	}

	freeMemoryMap := []*freeMemory{}
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == nil {
			j := i + 1
			for ; j < len(blocks) && blocks[j] == nil; j++ {
			}

			freeMemoryMap = append(freeMemoryMap, &freeMemory{at: i, len: j - i})
			i = j
		}
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		for ; blocks[i] == nil; i-- {
		}

		requiredSpace := blocks[i].usedBlocks
		for x := 0; x < len(freeMemoryMap); x++ {
			if freeMemoryMap[x].len >= requiredSpace && i > freeMemoryMap[x].at {
				for y := 0; y < requiredSpace; y++ {
					result[freeMemoryMap[x].at+y] = result[i]
					result[i] = nil
					i--
				}
				i++

				freeMemoryMap[x].at += requiredSpace
				freeMemoryMap[x].len -= requiredSpace

				break
			}
		}
	}

	return result
}

func compactBlocks(blocks []*diskFile) []*diskFile {
	var result []*diskFile

	i := 0
	j := len(blocks)
	for ; i < j; i++ {
		if blocks[i] == nil {
			j--
			for ; j > i; j-- {
				if blocks[j] != nil {
					result = append(result, blocks[j])
					break
				}
			}
		} else {
			result = append(result, blocks[i])
		}
	}

	return result
}

func getBlocks(diskMap []diskFile) []*diskFile {
	var result []*diskFile

	for _, diskFile := range diskMap {
		for i := 0; i < diskFile.usedBlocks; i++ {
			result = append(result, &diskFile)
		}

		for i := 0; i < diskFile.freeBlocks; i++ {
			result = append(result, nil)
		}
	}

	return result
}

func parse(input string) []diskFile {
	result := []diskFile{}

	var idSeed int
	var index int
	for {
		var freeBlocks int
		if index+1 < len(input) {
			freeBlocks = toInt(string(input[index+1]))
		}

		result = append(result, diskFile{
			id:         idSeed,
			usedBlocks: toInt(string(input[index])),
			freeBlocks: freeBlocks})

		idSeed++

		index += 2
		if index >= len(input) {
			break
		}
	}

	return result
}

func toInt(s string) int {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Cannot convert %s to int.", s)
	}
	return int(num)
}
