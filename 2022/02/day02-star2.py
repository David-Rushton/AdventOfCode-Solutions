import sys
import os

ROCK = 'A'
PAPER = 'B'
SCISSORS = 'C'

OUTCOME_WIN = 'Z'
OUTCOME_LOSE = 'X'
OUTCOME_DRAW = 'Y'

POINTS_ROCK = 1
POINTS_PAPER = 2
POINTS_SCISSORS = 3

POINTS_WIN = 6
POINTS_LOSE = 0
POINTS_DRAW = 3

POINTS_CODE_MAP = {
    ROCK: POINTS_ROCK,
    PAPER: POINTS_PAPER,
    SCISSORS: POINTS_SCISSORS
}

WINNING_COMBINATIONS = [
    (POINTS_ROCK, POINTS_PAPER),
    (POINTS_PAPER, POINTS_SCISSORS),
    (POINTS_SCISSORS,  POINTS_ROCK)
]


def main():

    totalScore = 0

    for game in getInput():
        if len(game) == 3:
            players = game.split(' ')
            opponent = players[0]
            player = getMoveForOutcome(opponent, players[1])
            gameScore = getScore(opponent, player)
            totalScore += getScore(opponent, player)
            print(f'\tGame {game}.  You scored {gameScore}')

    print(f'Total score {totalScore}')

def getMoveForOutcome(opponentMove, desiredOutcome):
    if desiredOutcome == OUTCOME_WIN:
        if opponentMove == ROCK:
            return PAPER
        elif opponentMove == PAPER:
            return SCISSORS
        elif opponentMove == SCISSORS:
            return ROCK

    if desiredOutcome == OUTCOME_LOSE:
        if opponentMove == ROCK:
            return SCISSORS
        elif opponentMove == PAPER:
            return ROCK
        elif opponentMove == SCISSORS:
            return PAPER

    if desiredOutcome == OUTCOME_DRAW:
        return opponentMove

def getScore(opponent, player):
    opponentScore = POINTS_CODE_MAP[opponent]
    playerScore = POINTS_CODE_MAP[player]

    score = playerScore

    if opponentScore == playerScore:
        return score + POINTS_DRAW

    if (opponentScore, playerScore) in WINNING_COMBINATIONS:
        return score + POINTS_WIN

    return score + POINTS_LOSE

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
