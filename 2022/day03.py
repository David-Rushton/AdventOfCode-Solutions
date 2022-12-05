import sys
import os

def main():
    rucksacks = getInput()
    totalScore = 0
    common = {}
    for rucksack in rucksacks:
        rucksackLen = len(rucksack)
        if rucksackLen > 0:
            if rucksackLen % 2 != 0:
                raise Exception("Expected even number of items")

            halfway = int(rucksackLen / 2)
            compartment1 = []
            compartment2 = []

            for i in range(halfway):
                compartment1.append(rucksack[i])
                compartment2.append(rucksack[halfway + i])

            for item in compartment1:
                if item in compartment2 and item not in common:

                    common.append(item)

            commonScore = sum(map(convertToScore, common))
            totalScore += commonScore

            print(f'\t{rucksack} = {commonScore}')

    print(f'Total score: {totalScore}')

def convertToScore(item):
    if ord(item) > 97:
        return ord(item) - 96;

    return ord(item) - 38;

def getInput():
    for arg in sys.argv[1:]:
        input = ''
        if os.path.exists(arg):
            for line in open(arg, 'r').readlines():
                yield line.replace('\n' ,'')
        else:
            yield arg
    yield ''

if __name__ == '__main__':
    main()
