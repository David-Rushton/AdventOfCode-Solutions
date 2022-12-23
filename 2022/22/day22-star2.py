from ast import Tuple
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

FACE_CHANGES = []

@dataclass
class Location:
    row: int
    column: int

@dataclass
class TryMoveResult:
    has_moved: bool
    location: Location
    direction: str


def main(is_test_mode: bool, path: str):
    os.system('cls')

    maze = parse_maze(path)
    instructions = list(parse_instructions(path))
    location = get_starting_location(maze)
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
                    direction = move_result.direction
                    maze[location.row][location.column] = direction
                    is_dirty = True
                else:
                    instruction = 0

                if is_dirty:
                    moves += 1
                    print_maze(maze, location, instruction, direction, moves, is_test_mode=is_test_mode)
                    time.sleep(sleep_in_secs)

    print_maze(maze, location, instruction, direction, moves, is_test_mode=True)
    print_password(location, direction)
    exit(0)

def get_starting_location(maze: list[str]):
    row = 0
    for column in range(len(maze[row])):
        if maze[row][column] == PATH:
            return Location(row, column - 1)

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
    candidate_location, new_direction = offset_location(candidate_location, location_offset, direction, maze)

    if candidate_location.row < len(maze) and candidate_location.column < len(maze[candidate_location.row]):
        if maze[candidate_location.row][candidate_location.column] in DIRECTIONS:
            return TryMoveResult(has_moved=True, location=candidate_location, direction=new_direction)

        if maze[candidate_location.row][candidate_location.column] == PATH:
            return TryMoveResult(has_moved=True, location=candidate_location, direction=new_direction)

        if maze[candidate_location.row][candidate_location.column] == WALL:
            return TryMoveResult(has_moved=False, location=starting_location, direction=new_direction)

    raise Exception('Unable to find next move from {starting_location.row}x{starting_location.column}')

def copy_location(original_location: Location) -> Location:
    return Location(original_location.row, original_location.column)

def offset_location(location: Location, offset: Location, direction: str, maze: list[str, str]) -> tuple[Location, str]:
    new_location = Location(location.row + offset.row, location.column + offset.column)
    new_direction = direction

    # The map is a 2D grid, that represents a 3D.
    # If we are no longer on the map (an empty cell (' ') or out of bounds) move to another face.
    is_out_of_bounds = new_location.row < 0 or new_location.row >= len(maze) or new_location.column < 0 or new_location.column >= len(maze[new_location.row])
    if is_out_of_bounds:
        new_location, new_direction = change_face(location, direction)
    else:
        if maze[new_location.row][new_location.column] == VOID:
            new_location, new_direction = change_face(location, direction)

    return new_location, new_direction

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

def print_maze(maze: list[str, str], location: Location, instruction: any, direction: str, moves: int, is_test_mode=True):
    print('\033[1;1H', end='')
    print('== Monkey Map ==')
    print(f'- Moves: {moves}      ')
    print(f'- Location: {location.row + 1}x{location.column + 1}      ')
    print(f'- Instruction: {instruction}   ')
    print(f'- Direction: {direction}\n')

    min_row = 0
    max_row = 30

    if len(maze) > max_row:
        min_row = max(location.row - 15, 0)
        max_row = min_row + 30


    for row in range(len(maze)):
        if row < min_row or row > max_row:
            continue

        column_value = str(row).ljust(4)
        for column in range(len(maze[row])):
            cell = maze[row][column]

            if cell == WALL:
                cell = f'\033[31m{WALL}\033[0m'

            if row == location.row and column == location.column:
                cell = f'\033[93m{cell}\033[0m'
            else:
                if cell in DIRECTIONS:
                    cell = f'\033[32m{cell}\033[0m'

            column_value += cell
        print(column_value.ljust(200))
    print()

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

def change_face(location: Location, direction: str) -> tuple[Location, str]:
    """
        This a hack.
        To writing code to fold a flat map into a 3D cube I made a physical model.
        The hardcoded values below were taken from this.
        [3D debugging cube](.\debugging-cube.jpg)
    """

    # Up from face 1 == right on face 6
    if direction == UP and location.row == 0 and location.column >= 50 and location.column <= 99:
        new_location = Location(row=location.column + 100, column=0)
        FACE_CHANGES.append(f'Up from face 1 == right on face 6: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, RIGHT

    # Up from face 2 == up on face 6
    if direction == UP and location.row == 0 and location.column >= 100 and location.column <= 149:
        new_location = Location(row=199, column=location.column - 100)
        FACE_CHANGES.append(f'Up from face 2 == up on face 6: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, UP

    # Right from face 2 == left on face 4
    if direction == RIGHT and location.column == 149 and location.row >= 0 and location.row <= 49:
        new_row = abs(location.row - 149)
        new_location = Location(row=new_row, column=99)
        FACE_CHANGES.append(f'Right from face 2 == left on face 4: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, LEFT

    # Down from face 2 == left on face 3
    if direction == DOWN and location.row == 49 and location.column >= 100 and location.column <= 149:
        new_row = location.column - 50
        new_location = Location(row=new_row, column=99)
        FACE_CHANGES.append(f'Down from face 2 == left on face 3: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, LEFT

    # Right from face 3 == up on face 2
    if direction == RIGHT and location.column == 99 and location.row >= 50 and location.row <= 99:
        new_column = location.row + 50
        new_location = Location(row=49, column=new_column)
        FACE_CHANGES.append(f'Right from face 3 == up on face 2: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, UP

    # Right from face 4 == left on face 2
    if direction == RIGHT and location.column == 99 and location.row >= 100 and location.row <= 149:
        new_row = abs(location.row - 149)
        new_location = Location(row=new_row, column=149)
        FACE_CHANGES.append(f'Right from face 4 == left on face 2: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, LEFT

    # Down from face 4 == left on face 6
    if direction == DOWN and location.row == 149 and location.column >= 50 and location.column <= 99:
        new_location = Location(row=location.column + 100, column=49)
        FACE_CHANGES.append(f'Down from face 4 == left on face 6: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, LEFT

    # Right from face 6 == up on face 4
    if direction == RIGHT and location.column == 49 and location.row >= 150 and location.row <= 199:
        new_location = Location(row=149, column=location.row - 100)
        FACE_CHANGES.append(f'Right from face 6 == up on face 4: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, UP

    # Down from face 6 == down on face 2
    if direction == DOWN and location.row == 199 and location.column >= 0 and location.column <= 49:
        new_location = Location(row=0, column=location.column + 100)
        FACE_CHANGES.append(f'Down from face 6 == down on face 2: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, DOWN

    # Left from face 6 == down on face 1
    if direction == LEFT and location.column == 0 and location.row >= 150 and location.row <= 199:
        new_location = Location(row=0, column=location.row - 100)
        FACE_CHANGES.append(f'Left from face 6 == down on face 1: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, DOWN

    # Left from face 5 == right on face 1
    if direction == LEFT and location.column == 0 and location.row >= 100 and location.row <= 149:
        new_row = abs(location.row - 149)
        new_location = Location(row=new_row, column=50)
        FACE_CHANGES.append(f'Left from face 5 == right on face 1: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, RIGHT

    # Up from face 5 == right on face 3
    if direction == UP and location.row == 100 and location.column >= 0 and location.column <= 49:
        new_location = Location(row=location.column + 50, column=50)
        FACE_CHANGES.append(f'Up from face 5 == right on face 3: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, RIGHT

    # Left from face 3 == down on face 5
    if direction == LEFT and location.column == 50 and location.row >= 50 and location.row <= 99:
        new_location = Location(row=100, column=location.row - 50)
        FACE_CHANGES.append(f'Left from face 3 == down on face 5: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, DOWN

    # Left from face 1 == right on face 5
    if direction == LEFT and location.column == 50 and location.row >= 0 and location.row <= 49:
        new_row = abs(location.row - 149)
        new_location = Location(row=new_row, column=0)
        FACE_CHANGES.append(f'Left from face 1 == right on face 5: {location.row}x{location.column} -> {new_location.row}x{new_location.column}')
        return new_location, RIGHT

    raise Exception(f'Unexpected fell off the cube.  Last location: {location.row}x{location.column}.  Facing: {direction}.')

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
