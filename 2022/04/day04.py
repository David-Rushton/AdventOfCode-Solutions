import sys
import os

def main():
    assignments = getInput()
    fullyOverlappingRangesDetected = 0
    partialOverlappingRangesDetected = 0

    for assignment in assignments:
        pairs = assignment.split(',')
        firstElf = pairs[0].split('-')
        secondElf = pairs[1].split('-')

        print(f'\t1: {firstElf[0]} -> {firstElf[1]} | 2: {secondElf[0]} -> {secondElf[1]}')

        if getIsFullyOverlapping(firstElf, secondElf):
            fullyOverlappingRangesDetected += 1
            print(f'\t\tFull overlap detected')

        if getIsPartiallyOverlapping(firstElf, secondElf):
            partialOverlappingRangesDetected += 1
            print(f'\t\tPartial overlap detected')

    print(f'\nDetected overlaps:\nFull: {fullyOverlappingRangesDetected}\nPartial: {partialOverlappingRangesDetected}')


def getIsFullyOverlapping(firstElf, secondElf):
    firstMin = int(firstElf[0])
    firstMax = int(firstElf[1])
    secondMin = int(secondElf[0])
    secondMax = int(secondElf[1])

    return (firstMin >= secondMin and firstMax <= secondMax) or (secondMin >= firstMin and secondMax <= firstMax)

def getIsPartiallyOverlapping(firstElf, secondElf):
    firstMin = int(firstElf[0])
    firstMax = int(firstElf[1])
    secondMin = int(secondElf[0])
    secondMax = int(secondElf[1])

    return (secondMin >= firstMin and secondMin <= firstMax) or (secondMax >= firstMin and secondMax <= firstMax) \
        or (firstMin >= secondMin and firstMin <= secondMax) or (firstMax >= secondMin and firstMax <= secondMax)

def getInput():
    for arg in sys.argv[1:]:
        if os.path.exists(arg):
            for line in open(arg, 'r').readlines():
                content = line.replace('\n' ,'')
                if len(content) > 0:
                    yield content
        else:
            yield arg

if __name__ == '__main__':
    main()
