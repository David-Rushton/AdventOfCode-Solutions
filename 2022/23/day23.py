from dataclasses import dataclass
import os
import sys
import time

@dataclass
class CandidateMove:
    elf: tuple[int, int]
    to: tuple[int, int]

ELF = '#'
EMPTY_TILE = '.'
ROUNDS = 10

NORTH_OFFSET = (0, -1)
NORTH_EAST_OFFSET = (1, -1)
EAST_OFFSET = (1, 0)
SOUTH_EAST_OFFSET = (1, 1)
SOUTH_OFFSET = (0, 1)
SOUTH_WEST_OFFSET = (-1, 1)
WEST_OFFSET = (-1, 0)
NORTH_WEST_OFFSET = (-1, -1)

NORTH_ISH_OFFSETS = [NORTH_OFFSET, NORTH_EAST_OFFSET, NORTH_WEST_OFFSET]
SOUTH_ISH_OFFSETS = [SOUTH_OFFSET, SOUTH_EAST_OFFSET, SOUTH_WEST_OFFSET]
WEST_ISH_OFFSETS = [WEST_OFFSET, NORTH_WEST_OFFSET, SOUTH_WEST_OFFSET]
EAST_ISH_OFFSETS = [EAST_OFFSET, NORTH_EAST_OFFSET, SOUTH_EAST_OFFSET]


def main(path: str) -> None:
    elves = set(parse_elves(path))
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
            if candidate_points[(candidate_move.to[0], candidate_move.to[1])] == 1:
                events.append(f'Elf moving from {candidate_move.elf[0]}x{candidate_move.elf[1]} -> {candidate_move.to[0]}x{candidate_move.to[1]}')

                elves.remove(candidate_move.elf)
                elves.add(candidate_move.to)

                if candidate_move.to[0] < min_x:
                    min_x = candidate_move.to[0]

                if candidate_move.to[0] > max_x:
                    max_x = candidate_move.to[0]

                if candidate_move.to[1] < min_y:
                    min_y = candidate_move.to[1]

                if candidate_move.to[1] > max_y:
                    max_y = candidate_move.to[1]

                elves_moved += 1
            else:
                events.append(f'Elf not moving from {candidate_move.to[0]}x{candidate_move.to[1]}')

        if elves_moved == 0:
            break

        offset_groups.append(offset_groups.pop(0))

        print_map(elves, round, events, min_x, max_x, min_y, max_y)

    min_x, min_y, max_x, max_y = get_enclosing_retangle(elves)
    print_map(elves, round, events, min_x, max_x, min_y, max_y)

def get_enclosing_retangle(elves: set[tuple[int, int]]) -> tuple[int, int, int, int]:
    min_x = min_y = max_x = max_y = 0
    for elf in elves:
        if elf[0] < min_x:
            min_x = elf[0]

        if elf[0] > max_x:
            max_x = elf[0]

        if elf[1] < min_y:
            min_y = elf[1]

        if elf[1] > max_y:
            max_y = elf[1]

    return min_x, min_y, max_x, max_y

def get_candidates(elves: set[tuple[int, int]], offset_groups = list[list[tuple[int, int]]]) -> tuple[list[CandidateMove], set[tuple[int, int]]]:
    candidate_moves = []
    candidate_points = {}

    for elf in elves:
        neighbours = list(get_neighbours(elf, elves))

        if len(neighbours) == 0:
            continue

        for offsets in offset_groups:
            if has_empty_offsets(neighbours, elf, offsets):
                candidate_point = get_offset(elf, offsets[0])

                if candidate_point in candidate_points:
                    candidate_points[candidate_point] += 1
                else:
                    candidate_points[candidate_point] = 1

                candidate_moves.append(CandidateMove(elf=elf, to=candidate_point))
                break

    return candidate_moves, candidate_points

def get_neighbours(elf: tuple[int, int], elves: set[tuple[int, int]]):
    for y in (-1, 0, 1):
        for x in (-1, 0 , 1):
            if not (y == 0 and x == 0):
                if (elf[0] + x, elf[1] + y) in elves:
                    yield (elf[0] + x, elf[1] + y)

def has_empty_offsets(elves: set[tuple[int, int]], starting_point: tuple[int, int], offsets: list[tuple[int, int]]) -> bool:
    for offset in offsets:
        if get_offset(starting_point, offset) in elves:
            return False

    return True

def get_offset(starting_point: tuple[int, int], offset: tuple[int, int]) -> tuple[int, int]:
    return (starting_point[0] + offset[0], starting_point[1] + offset[1])

def print_map(elves: set[tuple[int, int]], round: int, events: list[str], min_x: int, max_x: int, min_y: int, max_y: int):
    os.system('cls')
    print('\033[1;1H', end='')
    print('\n== Unstable Diffusion ==')
    print(f'- Round: {round}')
    print(f'- Bounds: {min_x}x{min_y} -> {max_x}x{max_y}')
    print(f'- Empty tiles: {(max_x - min_x + 1) * (max_y - min_y + 1) - len(elves)}\n')

    for y in range(min_y, max_y + 1):
        row = ''
        for x in range(min_x, max_x + 1):
            if (x, y) in elves:
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

    # time.sleep(.5)

def parse_elves(path: str) -> set[tuple[int, int]]:
    y = -1
    for line in open(path, 'r').read().splitlines():
        y += 1
        x = -1
        for char in line:
            x += 1
            if char == ELF:
                yield (x, y)


if __name__ == '__main__':
    path = 'input.txt' if len(sys.argv) == 1 else f'input.{sys.argv[1]}.txt'
    main(path)
