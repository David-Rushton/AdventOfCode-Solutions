import sys
import os

def main():
    elves = []
    currentCalories = 0
    for calories in getInput():
        if calories == '':
            print(f'We found an elf carrying {currentCalories} calories')
            elves.append(currentCalories)
            currentCalories = 0
        else:
            currentCalories += int(calories)

    sortedElves = sorted(elves, reverse=True)
    print(f'\nWe found {len(elves)} elves.')
    print(f'\tElf #1 {sortedElves[0]}')
    print(f'\tElf #1 {sortedElves[1]}')
    print(f'\tElf #1 {sortedElves[2]}')
    print(f'The top three are carrying {sum(sortedElves[0:3])}')


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
