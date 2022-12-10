import os
import re
import sys
from collections import namedtuple

def main(path: str):
    TOP = 0
    LEFT = 1
    moves = open(path, 'r').read().splitlines()
    knotLocations = [(0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0), (0, 0)]
    tailVisited = {(0, 0): 0}
    moveCounter = 1

    for move in moves:
        elements = move.split(' ')
        direction = elements[0]
        quantity = int(elements[1])

        while (quantity > 0):

            if direction == 'U':
                knotLocations[0] = (knotLocations[0][TOP] + 1, knotLocations[0][LEFT])
            if direction == 'D':
                knotLocations[0] = (knotLocations[0][TOP] - 1, knotLocations[0][LEFT])
            if direction == 'L':
                knotLocations[0] = (knotLocations[0][TOP], knotLocations[0][LEFT] + 1)
            if direction == 'R':
                knotLocations[0] = (knotLocations[0][TOP], knotLocations[0][LEFT] - 1)

            for knot in range(1, 10):
                topGap = knotLocations[knot - 1][TOP] - knotLocations[knot][TOP]
                leftGap = knotLocations[knot - 1][LEFT] - knotLocations[knot][LEFT]
                if (topGap < -1 or topGap > 1) or (leftGap < -1 or leftGap > 1):
                    if knotLocations[knot - 1][TOP] > knotLocations[knot][TOP]:
                        knotLocations[knot] = (knotLocations[knot][TOP] + 1, knotLocations[knot][LEFT])
                    if knotLocations[knot - 1][TOP] < knotLocations[knot][TOP]:
                        knotLocations[knot] = (knotLocations[knot][TOP] - 1, knotLocations[knot][LEFT])
                    if knotLocations[knot - 1][LEFT] > knotLocations[knot][LEFT]:
                        knotLocations[knot] = (knotLocations[knot][TOP], knotLocations[knot][LEFT] + 1)
                    if knotLocations[knot - 1][LEFT] < knotLocations[knot][LEFT]:
                        knotLocations[knot] = (knotLocations[knot][TOP], knotLocations[knot][LEFT] - 1)

            tailVisited[knotLocations[9]] = 0

            quantity -= 1

        print(f'\tmove #{moveCounter} | move {move} | head {knotLocations[0]} | tail {knotLocations[8]}')
        moveCounter += 1

    plotVisited(tailVisited)

    print(f'visited: {len(tailVisited)}')

def plotVisited(tailVisited):
    for top in range(-20, -60, -1):
        for left in range(-20, -225, -1):
            if (top, left) in tailVisited:
                print('#', end = '')
            else:
                print('.', end = '')
        print()

if __name__ == '__main__':
    isTest = sys.argv[1] in ['true', 'test', 'yes', 'on']
    path = 'day09.input.txt.test' if isTest else 'day09.input.txt'
    main(path)
