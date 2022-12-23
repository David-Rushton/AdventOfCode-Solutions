from dataclasses import dataclass
import os
import sys
import time

@dataclass
class Point:
    x: int
    y: int

@dataclass
class CandidateMove:
    elf: Point
    to: Point

ELF = '#'
EMPTY_TILE = '.'
ROUNDS = 10

NORTH_OFFSET = Point(x=0, y=-1)
NORTH_EAST_OFFSET = Point(x=1, y=-1)
EAST_OFFSET = Point(x=1, y=0)
SOUTH_EAST_OFFSET = Point(x=1, y=1)
SOUTH_OFFSET = Point(x=0, y=1)
SOUTH_WEST_OFFSET = Point(x=-1, y=1)
WEST_OFFSET = Point(x=-1, y=0)
NORTH_WEST_OFFSET = Point(x=-1, y=-1)

NORTH_ISH_OFFSETS = [NORTH_OFFSET, NORTH_EAST_OFFSET, NORTH_WEST_OFFSET]
SOUTH_ISH_OFFSETS = [SOUTH_OFFSET, SOUTH_EAST_OFFSET, SOUTH_WEST_OFFSET]
WEST_ISH_OFFSETS = [WEST_OFFSET, NORTH_WEST_OFFSET, SOUTH_WEST_OFFSET]
EAST_ISH_OFFSETS = [EAST_OFFSET, NORTH_EAST_OFFSET, SOUTH_EAST_OFFSET]


def main(path: str) -> None:
    elves = list(parse_elves(path))
    min_x = min_y = 0
    max_x = max_y = 15
    offset_groups = [NORTH_ISH_OFFSETS, SOUTH_ISH_OFFSETS, WEST_ISH_OFFSETS, EAST_ISH_OFFSETS]
    round = 0

    os.system('cls')
    print_map(elves, round, ['Elf starting positions'], min_x, max_x, min_y, max_y)

    # while round < ROUNDS:
    while True:
        round += 1
        elves_moved = 0
        events = []

        # get candidated locations
        candidate_moves, candidate_points = get_candidates(elves, offset_groups)

        # move
        for candidate_move in candidate_moves:
            if candidate_points.count(candidate_move.to) == 1:
                events.append(f'Elf moving from {candidate_move.elf.x}x{candidate_move.elf.y} -> {candidate_move.to.x}x{candidate_move.to.y}')

                # don't replace elf
                # update; so the change is respected in the elves collection.
                candidate_move.elf.x = candidate_move.to.x
                candidate_move.elf.y = candidate_move.to.y

                if candidate_move.to.x < min_x:
                    min_x = candidate_move.to.x

                if candidate_move.to.x > max_x:
                    max_x = candidate_move.to.x

                if candidate_move.to.y < min_y:
                    min_y = candidate_move.to.y

                if candidate_move.to.y > max_y:
                    max_y = candidate_move.to.y

                elves_moved += 1
            else:
                events.append(f'Elf not moving from {candidate_move.to.x}x{candidate_move.to.y}')

        if elves_moved == 0:
            break

        offset_groups.append(offset_groups.pop(0))

        print_map(elves, round, events, min_x, max_x, min_y, max_y)

    min_x, min_y, max_x, max_y = get_enclosing_retangle(elves)
    print_map(elves, round, events, min_x, max_x, min_y, max_y)

def get_enclosing_retangle(elves: list[Point]) -> tuple[int, int, int, int]:
    min_x = min_y = max_x = max_y = 0
    for elf in elves:
        if elf.x < min_x:
            min_x = elf.x

        if elf.x > max_x:
            max_x = elf.x

        if elf.y < min_y:
            min_y = elf.y

        if elf.y > max_y:
            max_y = elf.y

    return min_x, min_y, max_x, max_y

def get_candidates(elves: list[Point], offset_groups = list[list[Point]]) -> tuple[list[CandidateMove], dict[Point, int]]:
    candidate_moves = []
    candidate_points = []

    for elf in elves:
        if has_empty_neighbours(elves, elf):
            continue

        for offsets in offset_groups:
            if has_empty_offsets(elves, elf, offsets):
                candidate_point = get_offset(elf, offsets[0])
                candidate_points.append(candidate_point)
                candidate_moves.append(CandidateMove(elf=elf, to=candidate_point))
                break

    return candidate_moves, candidate_points

def has_empty_offsets(elves: list[Point], starting_point: Point, offsets: list[Point]) -> bool:
    for offset in offsets:
        if get_offset(starting_point, offset) in elves:
            return False

    return True

def has_empty_neighbours(elves: list[Point], starting_point: Point) -> bool:
    for y_offset in (-1, 0, 1):
        for x_offset in (-1, 0, 1):
            if not (y_offset == 0 and x_offset == 0):
                if get_offset(starting_point, Point(x_offset, y_offset)) in elves:
                    return False

    return True

def get_offset(starting_point: Point, offset: Point) -> Point:
    return Point(x=starting_point.x + offset.x, y=starting_point.y + offset.y)

def print_map(elves: list[Point], round: int, events: list[str], min_x: int, max_x: int, min_y: int, max_y: int):
    os.system('cls')
    print('\033[1;1H', end='')
    print('\n== Unstable Diffusion ==')
    print(f'- Round: {round}')
    print(f'- Bounds: {min_x}x{min_y} -> {max_x}x{max_y}')
    print(f'- Empty tiles: {(max_x - min_x + 1) * (max_y - min_y + 1) - len(elves)}\n')

    for y in range(min_y, max_y + 1):
        row = ''
        for x in range(min_x, max_x + 1):
            if Point(x, y) in elves:
                row += ELF
            else:
                row += EMPTY_TILE
        print(row.ljust(50))
    print(''.ljust(50))

    # for event in events:
    #     print(f'- {event}')

    # if True:
    #     # print('\nPress any key to continue\n')
    #     input()

    time.sleep(.5)

def parse_elves(path: str) -> list[Point]:
    y = -1
    for line in open(path, 'r').read().splitlines():
        y += 1
        x = -1
        for char in line:
            x += 1
            if char == ELF:
                yield Point(x, y)


if __name__ == '__main__':
    path = 'input.txt' if len(sys.argv) == 1 else f'input.{sys.argv[1]}.txt'
    main(path)
