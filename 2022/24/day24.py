from blizzard_map import CachedBlizzardMap
from data_types import *
from route_planner import RoutePlanner
import sys


def main(path: str):
    valley_map, blizzard_map = parse_initial_state(path)
    explorer = valley_map.entry

    print('\n== Blizzard Basin ==')

    best_time = RoutePlanner().get_best_time_to_exit(explorer, valley_map, blizzard_map)

    print(f'\n- Best time: {best_time}\n')

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
        entry=Location(1, 0),
        exit=Location(max_x - 1, max_y),
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
