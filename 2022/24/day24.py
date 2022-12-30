from blizzard_map import CachedBlizzardMap
from dataclasses import dataclass
from enum import Enum
import sys

EXPLORER = 'E'
WALL = '#'
PATH = '.'
UP = '^'
DOWN = 'v'
LEFT = '>'
RIGHT = '<'
DIRECTIONS = [UP, DOWN, LEFT, RIGHT]

@dataclass(eq=True, frozen=True)
class Location:
    x: int
    y: int

@dataclass
class ValleyMap:
    start: Location
    end: Location
    max_x: int
    max_y: int
    locations: dict[Location, str]

def main(path: str):
    valley_map, blizzard_map = parse_initial_state(path)
    explorer = valley_map.start

    time = 0
    while time < 5:
        blizzards = blizzard_map.get_map(time)
        print_valley_map(valley_map, explorer, blizzards)
        input()
        print()
        time += 1

def print_valley_map(valley_map: ValleyMap, explorer: Location, blizzards: dict[Location, str]) -> None:
    for y in range(valley_map.max_y + 1):
        for x in range(valley_map.max_x + 1):
            current_location = Location(x, y)
            cell = valley_map.locations[Location(x, y)]

            if cell == WALL:
                cell = f'\033[32m{WALL}\033[0m'

            if current_location in blizzards:
                cell = f'\033[94m{blizzards[current_location]}\033[0m'

            # Always consider our explorer's location last.
            # To ensure it is printed when a cell contain more than 1 item.
            if current_location == explorer:
                cell = '\033[1;35mP\033[0m'

            print(cell, end='')
        print()

def parse_initial_state(path: str) -> tuple[ValleyMap, CachedBlizzardMap]:
    y = 0
    max_x = 0
    blizzards_map = CachedBlizzardMap()
    locations = {}

    for row in open(path, 'r').read().splitlines():

        if max_x == 0:
            max_x = len(row) - 1

        for x in range(len(row)):
            cell = row[x]

            if cell in DIRECTIONS:
                blizzards_map.add_blizzard(Location(x, y), cell)
                cell = PATH

            locations[Location(x, y)] = cell

        y += 1

    max_y = y - 1

    # Start cell is always 1, 0
    # End cell is always 1 to the right of the bottom right corner
    valley_map = ValleyMap(
        start=Location(1, 0),
        end=Location(max_x - 1, max_y),
        max_x=max_x,
        max_y=max_y,
        locations=locations
    )

    blizzards_map.set_boundaries(max_x, max_y)

    return (valley_map, blizzards_map)


if __name__ == '__main__':
    is_test = True if sys.argv[1] == 'test' else False
    path = 'input.test.txt' if is_test else 'input.txt'
    main(path)
