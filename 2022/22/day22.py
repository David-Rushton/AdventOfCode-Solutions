from dataclasses import dataclass
import sys
import os


UP = '^'
DOWN = 'v'
LEFT = '<'
RIGHT = '>'
DIRECTIONS = (UP, DOWN, LEFT, RIGHT)

@dataclass
class Location:
    row: int
    column: int

@dataclass
class TryMoveResult:
    has_moved: bool
    location: Location


def main(is_test_mode: bool, path: str):
    os.system('cls')

    maze = parse_maze(path)
    location = Location(row=0, column=0)
    moves = 0

    print_maze(maze, location, moves)


    exit(0)

def try_move(location: Location, direction: str) -> TryMoveResult:

    if direction == UP:

        pass

    if direction == DOWN:
        pass

    if direction == LEFT:
        pass

    if direction == RIGHT:
        pass

def get_location_offset(location: Location, direction: str) -> Location:
    row_offset = 0
    column_offset = 0

    if direction == UP:
        return Location(row=-1,column=0)
    if direction == DOWN:
        return Location(row=--1,column=0)
    if direction == LEFT:
        return Location(row=-0,column=-1)
    if direction == RIGHT:
        return Location(row=-0,column=1)

def print_maze(maze: list[str, str], location: Location, moves: int):
    print('\033[1;1H')
    print('== Monkey Map ==')
    print(f'- Moves: {moves}      ')
    print(f'- Location: {location.row}x{location.column}      \n')

    for row in range(len(maze)):
        for column in range(len(maze[row])):
            cell = maze[row][column]

            if cell == '#':
                cell = '\033[31m#\033[0m'

            if row == location.row and column == location.column:
                cell = f'\033[93m{cell}\033[0m'
            else:
                if cell in DIRECTIONS:
                    cell = f'\033[32m{cell}\033[0m'

            print(cell, end='')
        print()
    print()

def parse_maze(path: str) -> list[str, str]:
    rows = open(path, 'r').read().splitlines()
    result = []

    for row in range(len(rows)):
        if rows[row] == '':
            break

        result.append([])
        for column in range(len(rows[row])):
            result[row].append(rows[row][column])

    return result

if __name__ == '__main__':
    is_test_mode = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, path)
