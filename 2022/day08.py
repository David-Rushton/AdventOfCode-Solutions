import os
import re
import sys
from collections import namedtuple

def main(isStarTwo: bool, path: str):
    trees = open(path, 'r').read().splitlines()
    columns = len(trees[0])
    rows = len(trees)
    treeMap = [[0 for x in range(rows)] for y in range(columns)]
    treeScore = [[0 for x in range(rows)] for y in range(columns)]

    # looking down
    for column in range(columns):
        highestTree = -1
        for row in range(rows):
            if int(trees[row][column]) > highestTree:
                treeMap[row][column] = 1
                highestTree = int(trees[row][column])

    # # looking up
    for column in range(-1, (columns + 1) * -1, -1):
        highestTree = -1
        for row in range(-1, (rows + 1) * -1, -1):
            if int(trees[row][column]) > highestTree:
                treeMap[row][column] = 1
                highestTree = int(trees[row][column])

    # looking left
    for row in range(rows):
        highestTree = -1
        for column in range(columns):
            if int(trees[row][column]) > highestTree:
                treeMap[row][column] = 1
                highestTree = int(trees[row][column])

    # looking right
    for row in range(rows):
        highestTree = -1
        for column in range(-1, (columns + 1) * -1, -1):
            if int(trees[row][column]) > highestTree:
                treeMap[row][column] = 1
                highestTree = int(trees[row][column])

    print(treeMap)

    score = 0
    for column in range(columns):
        for row in range(rows):
            score += treeMap[row][column]

    print(f'score: {score}')


    for row in range(rows):
        for column in range(columns):

            # looking up
            upScore = 0
            for rowOffset in range(row - 1, -1, -1):
                if trees[row][column] <= trees[rowOffset][column]:
                    upScore += 1
                    break
                else:
                    upScore += 1

            # looking down
            downScore = 0
            for rowOffset in range(row + 1, rows):
                if trees[row][column] <= trees[rowOffset][column]:
                    downScore += 1
                    break
                else:
                    downScore += 1

            # looking left
            leftScore = 0
            for columnOffset in range(column - 1, -1, -1):
                if trees[row][column] <= trees[row][columnOffset]:
                    leftScore += 1
                    break
                else:
                    leftScore += 1

            # looking right
            rightScore = 0
            for columnOffset in range(column + 1, columns):
                if trees[row][column] <= trees[row][columnOffset]:
                    rightScore += 1
                    break
                else:
                    rightScore += 1

            # upScore = upScore if upScore > 0 else 1
            # downScore = downScore if downScore > 0 else 1
            # leftScore = leftScore if leftScore > 0 else 1
            # rightScore = rightScore if rightScore > 0 else 1

            treeScore[row][column] = upScore * downScore * leftScore * rightScore

    print(treeScore)

    maxScore = 0
    for row in range(rows):
        for column in range(columns):
            if treeScore[row][column] > maxScore:
                maxScore = treeScore[row][column]

    print(f'Best score {maxScore}')






if __name__ == '__main__':
    isTest = sys.argv[1] in ['true', 'test', 'yes', 'on']
    isStarTwo = sys.argv[2] in ['star2', 'star-2', 'startwo', 'star-two']
    path = 'day08.input.txt.test' if isTest else 'day08.input.txt'
    main(isStarTwo, path)
