import sys
import os

def main():
    rucksacks = getInput()
    totalScore = 0
    badges = []
    counter = 0
    rucksackContents = {
        0: [],
        1: [],
        2: []
    }

    for rucksack in rucksacks:
        if len(rucksack) > 0:

            groupCounter = counter % 3
            rucksackContents[groupCounter] = []

            for i in range(len(rucksack)):
                rucksackContents[groupCounter].append(rucksack[i])

            if groupCounter == 2:
                common = []
                for item in rucksackContents[0]:
                    if item in rucksackContents[1] and item in rucksackContents[2] and item not in common:
                        common.append(item)
                        badges.append(item)
                        print(f'\tline: {counter} || group: {groupCounter} || item: {item}')

            counter += 1

    commonScore = sum(map(convertToScore, badges))
    print(f'Total score: {commonScore} || Badges found: {len(badges)}')

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
