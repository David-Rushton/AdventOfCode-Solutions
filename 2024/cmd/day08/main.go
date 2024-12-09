package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/aoc"
	"github.com/David-Rushton/AdventOfCode-Solutions/tree/main/2024/internal/iostr"
)

type diskFile struct {
	id         int
	usedBlocks int
	freeBlocks int
}

func (df *diskFile) toBlocks() string {
	return fmt.Sprintf(
		"%s%s",
		strings.Repeat(strconv.FormatInt(int64(df.id), 10), df.usedBlocks),
		strings.Repeat(".", df.freeBlocks))
}

func main() {
	fmt.Println("--- Day 9: Disk Fragmenter ---")
	fmt.Println()

	diskMap := parse(aoc.Input[0])
	blocks := getBlocks(diskMap)
	compacted := compact(blocks, aoc.Star == aoc.StarTwo)
	checksum := getChecksum(compacted)

	iostr.Verbosef("  Blocks:\t%v", blocks)
	iostr.Verbosef("  Compacted:\t%v", compacted)
	iostr.Verbosef("  Checksum:\t%d", checksum)

	fmt.Println()
	fmt.Printf("Result: %d\n", checksum)
}

func getChecksum(blocks []*diskFile) int64 {
	var result int64

	for i, block := range blocks {
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
	panic("not implemented")
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
