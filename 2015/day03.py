from collections import namedtuple
import sys
import os

def main():
    top = 0
    left = 0
    instructions = getInput()
    visitedLocations = {}

    for instruction in instructions:
        for current in instruction:
            if current == '^':
                    top += 1
            elif current == '>':
                    left -= 1
            elif current == 'v':
                    top -= 1
            elif current == '<':
                    left += 1

            key = (top, left)
            if key in visitedLocations:
                visitedLocations[key] += 1
                print(f'\tRepeat visit to {visitedLocations[key]} {top}-{left}')
            else:
                print(f'\tFirst visit to {key} {top}-{left}')
                visitedLocations[key] = 1

    print(f'Visited {len(visitedLocations)} houses')


def getInput():
    for arg in sys.argv[1:]:
        if os.path.exists(arg):
            for line in open(arg, 'r').readlines():
                if len(line) > 1:
                    yield line.replace('\n' ,'')
        else:
            yield arg

if __name__ == '__main__':
    main()
