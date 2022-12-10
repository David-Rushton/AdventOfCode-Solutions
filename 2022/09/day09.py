import os
import re
import sys
from collections import namedtuple

def main(path: str):
    TOP = 0
    LEFT = 1
    moves = open(path, 'r').read().splitlines()
    headLocation = (0, 0)
    tailLocation = (0, 0)
    tailVisited = {(0, 0): 0}
    moveCounter = 1

    for move in moves:
        elements = move.split(' ')
        direction = elements[0]
        quantity = int(elements[1])

        while (quantity > 0):

            if direction == 'U':
                headLocation = (headLocation[TOP] + 1, headLocation[LEFT])
            if direction == 'D':
                headLocation = (headLocation[TOP] - 1, headLocation[LEFT])
            if direction == 'L':
                headLocation = (headLocation[TOP], headLocation[LEFT] + 1)
            if direction == 'R':
                headLocation = (headLocation[TOP], headLocation[LEFT] - 1)

            topGap = headLocation[TOP] - tailLocation[TOP]
            leftGap = headLocation[LEFT] - tailLocation[LEFT]
            if (topGap < -1 or topGap > 1) or (leftGap < -1 or leftGap > 1):
                if headLocation[TOP] > tailLocation[TOP]:
                    tailLocation = (tailLocation[TOP] + 1, tailLocation[LEFT])
                if headLocation[TOP] < tailLocation[TOP]:
                    tailLocation = (tailLocation[TOP] - 1, tailLocation[LEFT])
                if headLocation[LEFT] > tailLocation[LEFT]:
                    tailLocation = (tailLocation[TOP], tailLocation[LEFT] + 1)
                if headLocation[LEFT] < tailLocation[LEFT]:
                    tailLocation = (tailLocation[TOP], tailLocation[LEFT] - 1)

                tailVisited[tailLocation] = 0

            quantity -= 1

        print(f'\tmove #{moveCounter} | move {move} | head {headLocation} | tail {tailLocation} | gap ({topGap}, {leftGap})')
        moveCounter += 1

    print(f'visited: {len(tailVisited)}')


if __name__ == '__main__':
    isTest = sys.argv[1] in ['true', 'test', 'yes', 'on']
    path = 'day09.input.txt.test' if isTest else 'day09.input.txt'
    main(path)
