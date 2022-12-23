from dataclasses import dataclass
import sys
import os
import re
import time

VOID = ' '
WALL = '#'
PATH = '.'

UP = '^'
DOWN = 'v'
LEFT = '<'
RIGHT = '>'
DIRECTIONS = (UP, DOWN, LEFT, RIGHT)

ROTATE_LEFT = 'L'
ROTATE_RIGHT = 'R'
ROTATIONS = (ROTATE_LEFT, ROTATE_RIGHT)

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
    instructions = list(parse_instructions(path))
    location = Location(row=0, column=0)
    direction = '>'
    moves = 0
    sleep_in_secs = .5 if is_test_mode else 0

    for instruction in instructions:

        if instruction in ROTATIONS:
            direction = rotate_direction(direction, instruction)
            maze[location.row][location.column] = direction

        if isinstance(instruction, int):
            while instruction > 0:
                is_dirty = False
                instruction -= 1

                move_result = try_move(maze, location, direction)
                if move_result.has_moved:
                    location = move_result.location
                    maze[location.row][location.column] = direction
                    is_dirty = True
                else:
                    instruction = 0

                if is_dirty:
                    moves += 1
                    print_maze(maze, location, instruction, direction, moves, show_map=is_test_mode)
                    time.sleep(sleep_in_secs)

    print_maze(maze, location, instruction, direction, moves, show_map=True)
    print_password(location, direction)
    exit(0)

def rotate_direction(direction: str, instruction: str):
    if instruction == ROTATE_LEFT:
        if direction == UP:
            return LEFT
        if direction == DOWN:
            return RIGHT
        if direction == LEFT:
            return DOWN
        if direction == RIGHT:
            return UP

    if instruction == ROTATE_RIGHT:
        if direction == UP:
            return RIGHT
        if direction == DOWN:
            return LEFT
        if direction == LEFT:
            return UP
        if direction == RIGHT:
            return DOWN

    raise Exception(f'Rotation instruction not supported: {instruction}')

def try_move(maze: list[str, str], starting_location: Location, direction: str) -> TryMoveResult:
    location_offset = get_location_offset(direction)
    candidate_location = copy_location(starting_location)
    iterations = 0

    while True:
        candidate_location = offset_location(candidate_location, location_offset, maze)

        if candidate_location.row < len(maze) and candidate_location.column < len(maze[candidate_location.row]):
            if maze[candidate_location.row][candidate_location.column] in DIRECTIONS:
                return TryMoveResult(has_moved=True, location=candidate_location)

            if maze[candidate_location.row][candidate_location.column] == PATH:
                return TryMoveResult(has_moved=True, location=candidate_location)

            if maze[candidate_location.row][candidate_location.column] == WALL:
                return TryMoveResult(has_moved=False, location=starting_location)

        iterations += 1
        if iterations > 1000:
            break

    raise Exception('Unable to find candidate move in {iterations} moves')

def copy_location(original_location: Location) -> Location:
    return Location(original_location.row, original_location.column)

def offset_location(location: Location, offset: Location, maze: list[str, str]) -> Location:
    new_location = Location(location.row + offset.row, location.column + offset.column)

    if offset.row != 0 and new_location.row < 0:
        new_location.row = len(maze) - 1

    if offset.row != 0 and new_location.row >= len(maze):
        new_location.row = 0

    if offset.column != 0 and new_location.column < 0:
        new_location.column = len(maze[new_location.row]) - 1

    if offset.column != 0 and new_location.column >= len(maze[new_location.row]):
        new_location.column = 0

    return new_location

def get_location_offset(direction: str) -> Location:
    if direction == UP:
        return Location(row=-1, column=0)
    if direction == DOWN:
        return Location(row=1, column=0)
    if direction == LEFT:
        return Location(row=-0, column=-1)
    if direction == RIGHT:
        return Location(row=-0, column=1)

    raise Exception('Direction not supported: {direction}')

def print_maze(maze: list[str, str], location: Location, instruction: any, direction: str, moves: int, show_map=True):
    print('\033[1;1H', end='')
    print('== Monkey Map ==')
    print(f'- Moves: {moves}      ')
    print(f'- Location: {location.row + 1}x{location.column + 1}      ')
    print(f'- Instruction: {instruction}   ')
    print(f'- Direction: {direction}\n')



    min_row = 0
    max_row = len(maze)
    if max_row > 30:
        min_row = location.row - 20 if location.row > 20 else 0
        max_row = min_row + 40

    if True:
        for row in range(len(maze)):

            if row < min_row or row > max_row:
                continue;

            for column in range(len(maze[row])):
                cell = maze[row][column]

                if cell == WALL:
                    cell = f'\033[31m{WALL}\033[0m'

                if row == location.row and column == location.column:
                    cell = f'\033[93m{cell}\033[0m'
                else:
                    if cell in DIRECTIONS:
                        cell = f'\033[32m{cell}\033[0m'

                print(cell, end='')
            print()
        print()

    # if not is_test_mode:
    #     x = input()

def print_password(location: Location, direction: str):
    direction_scores = {
        RIGHT: 0,
        DOWN: 1,
        LEFT: 2,
        UP: 3
    }
    direction_score = direction_scores[direction]

    print('\n== Password ==')
    print(f'1000 * {location.row + 1} + 4 * {location.column + 1} + {direction_score}')
    print(f'= {1000 * (location.row + 1) + 4 * (location.column + 1) + direction_score}\n')

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

def parse_instructions(path: str) -> list:
    content = open(path, 'r').read()
    instructions = re.findall('[0-9LR]+', content)[0]
    buffer = ''

    for character in instructions:
        if character.isnumeric():
            buffer += character
        else:
            yield int(buffer)
            buffer = ''

            if character in ROTATIONS:
                yield character
            else:
                raise Exception(f'Unsupported rotation: {character}')

    if len(buffer) > 0:
        yield int(buffer)


if __name__ == '__main__':
    is_test_mode = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, path)
