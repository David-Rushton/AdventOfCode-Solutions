import sys
import os

A = X = ROCK = 1
B = Y = PAPER = 2
C = Z = SCISSORS = 3

ROCK_CODE = ('A', 'X')
PAPER_CODE = ('B', 'Y')
SCISSORS_CODE = ('C', 'Z')


WIN = 6
LOSE = 0
DRAW = 3

WINNING_COMBINATIONS = [
    (ROCK, PAPER),
    (PAPER, SCISSORS),
    (SCISSORS,  ROCK)
]


def main():

    totalScore = 0

    for game in getInput():
        if len(game) == 3:
            players = game.split(' ')
            opponent = lookupMoveValue(players[0])
            player = lookupMoveValue(players[1])
            gameScore = getScore(opponent, player)
            totalScore += getScore(opponent, player)
            print(f'\tGame {game}.  You scored {gameScore}')

    print(f'Total score {totalScore}')

def lookupMoveValue(moveCode):
    if moveCode in ROCK_CODE:
        return ROCK

    if moveCode in PAPER_CODE:
        return PAPER

    return SCISSORS

def getScore(opponent, player):
    score = player

    if opponent == player:
        return score + DRAW

    if (opponent, player) in WINNING_COMBINATIONS:
        return score + WIN

    return score + LOSE

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
