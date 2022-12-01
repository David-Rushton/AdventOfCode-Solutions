from collections import namedtuple
import sys
import os

def main():
    santaKey = (0, 0)
    roboKey = (0, 0)
    visited = {(0, 0): 1}
    santaTurn = True
    instructions = getInput()

    for instruction in instructions:
        for direction in instruction:

            print(f'Heading: {direction}')

            if santaTurn == True:
                print(f'\tSanta moving from {santaKey[0]}x{santaKey[1]}')
                santaKey = updateKey(santaKey, direction)
                key = santaKey

            if santaTurn == False:
                print(f'\tRobo moving from {roboKey[0]}x{roboKey[1]}')
                roboKey = updateKey(roboKey, direction)
                key = roboKey

            if key in visited:
                visited[key] += 1
                print(f'\t{key[0]}x{key[1]}: {visited[key]}')
            else:
                visited[key] = 1
                print(f'\t{key[0]}x{key[1]}: {visited[key]}')

            santaTurn = not santaTurn

    print(f'Visited: {len(visited)}')


def updateKey(key, direction):
    top = key[0]
    left = key[1]

    if direction == '^':
            top += 1
    elif direction == '>':
            left -= 1
    elif direction == 'v':
            top -= 1
    elif direction == '<':
            left += 1

    return (top, left)

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
