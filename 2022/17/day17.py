from dataclasses import dataclass
from functools import partial
import sys
import os
import time

GRID_WIDTH = 7
# SETTLED_ROCK_TARGET = 2022
SETTLED_ROCK_TARGET = 1000000000000
FRAME_DELAY_IN_MS = 0
SAVE_LAST_X_SETTLED_POINTS = (7 * 25)
LO_SAMPLE_RATE = 500
HIGH_SAMPLE_RATE = 10000

@dataclass
class Point:
    x: int
    y: int


def main(path: str):
    sequence = open(path, 'r').read().strip()
    next_airflow = get_airflow(sequence)
    next_rock = get_rocks()
    move = 0
    settled_rocks = []
    count_of_settled_rocks = 0
    max_y = -1
    # sample = open('sample.csv', 'x')
    # sample.write('move,max_y,settled_rocks\n')

    os.system('cls')
    print_grid(move, count_of_settled_rocks, [], settled_rocks, max_y, force_redraw=True)

    while True:
        rock = next_rock(max_y + 4)
        print_grid(move, count_of_settled_rocks, rock, settled_rocks, max_y)

        while True:

            # try move across
            x_offset = -1 if next_airflow() == '<' else 1
            offset_rock = list(map(partial(offset_point, x_offset=x_offset, y_offset=0), rock))
            if can_move(offset_rock, settled_rocks):
                rock = offset_rock

            print_grid(move, count_of_settled_rocks, rock, settled_rocks, max_y)

            # try move down
            offset_rock = list(map(partial(offset_point, x_offset=0, y_offset=-1), rock))
            if can_move(offset_rock, settled_rocks):
                rock = offset_rock
            else:
                for point in rock:
                    if point.y > max_y:
                        max_y = point.y
                    settled_rocks.append(point)

                    while len(settled_rocks) > SAVE_LAST_X_SETTLED_POINTS:
                        settled_rocks.pop(0)

                count_of_settled_rocks += 1

                rock = []

            move += 1
            # if move % (10091) == 0:
            # # if move % (10000) == 0:
            #     print_grid(move, count_of_settled_rocks, rock, settled_rocks, max_y, force_redraw=True)
            #     sample.write(f'{move},{max_y},{count_of_settled_rocks}\n')
            #     # x = input()


            if count_of_settled_rocks == 1600:
                print_grid(move, count_of_settled_rocks, rock, settled_rocks, max_y, force_redraw=True)
                x = input()


            print_grid(move, count_of_settled_rocks, rock, settled_rocks, max_y)

            if rock == []:
                break
        if count_of_settled_rocks == SETTLED_ROCK_TARGET:
            print_grid(move, count_of_settled_rocks, rock, settled_rocks, max_y, force_redraw=True)
            break

def can_move(rock: list[Point], settled_rocks: list[Point]):
    for point in rock:
        if point.x < 0 or point.x >= GRID_WIDTH:
            return False

        if point.y == -1:
            return False

        if point in settled_rocks:
            return False

    return True

def print_grid(move: int, count_of_settled_rocks: int, falling_rock: list[Point], settled_rocks: list[Point], max_y: int, force_redraw=False):

    is_verbose_mode = SETTLED_ROCK_TARGET <= 20
    sample_rate = LO_SAMPLE_RATE if SETTLED_ROCK_TARGET < 3000 else HIGH_SAMPLE_RATE
    if force_redraw or is_verbose_mode or move % sample_rate == 0:

        print('\033[1;1H')
        print('== Map ==')
        print(f'- Move: {move}          ')
        print(f'- Max Y: {max_y + 1}          ')
        print(f'- Settled Rocks: {count_of_settled_rocks}          \n')

        for y in range(max_y + 7, max_y - 25, -1):
            if y < -1:
                continue

            print(str(y).ljust(6), end='')

            for x in range(-1, GRID_WIDTH + 1):
                point = Point(x, y)
                cell = '.'

                if x in (-1, GRID_WIDTH):
                    cell = '|'

                if y == -1:
                    cell = '-'

                if (x == -1 and y == -1) or (x == GRID_WIDTH and y == -1):
                    cell = '+'

                if point in falling_rock:
                    cell = '@'

                if point in settled_rocks:
                    cell = '#'

                print(cell, end='')
            print()

        time.sleep(FRAME_DELAY_IN_MS)

def get_airflow(sequence: str):
    next = -1

    def get_next():
        nonlocal next
        nonlocal sequence
        next = 0 if next == len(sequence) - 1 else next + 1
        next_item = sequence[next]

        if next_item not in ['<', '>']:
            raise Exception("Unsupported airflow direction: {next_item}")

        return next_item
    return get_next

def get_rocks():

    sequence = [
        # -
        [Point(2, 0), Point(3, 0), Point(4, 0), Point(5, 0)],
        # +
        [Point(3, 2), Point(2, 1), Point(3, 1), Point(4, 1), Point(3, 0)],
        # J
        [Point(4, 2), Point(4, 1), Point(4, 0), Point(3, 0), Point(2, 0)],
        # |
        [Point(2, 0), Point(2, 1), Point(2, 2), Point(2, 3)],
        # []
        [Point(2, 0), Point(3, 0), Point(2, 1), Point(3, 1)]
    ]
    next = -1
    iteration = 0

    def get_next(y_offset):
        nonlocal next
        nonlocal iteration
        iteration += 1
        next = 0 if next == len(sequence) - 1 else next + 1
        next_item = sequence[next]

        if iteration % 5 == 1:
            if not next_item == sequence[0]:
                raise Exception(f'Returned incorrect shape: {iteration} {iteration % 5}')

        next_item = list(map(partial(offset_point, x_offset=0, y_offset=y_offset), sequence[next]))
        return next_item
    return get_next

def offset_point(point: Point, x_offset: int, y_offset: int):
    return Point(point.x + x_offset, point.y + y_offset)

if __name__ == '__main__':
    path = 'input.txt.test' if sys.argv[1] == 'test' else 'input.txt'
    main(path)

