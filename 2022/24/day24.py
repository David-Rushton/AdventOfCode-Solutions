from blizzard_map import CachedBlizzardMap
from data_types import *
from route_planner import RoutePlanner
import sys


def main(is_star_two: bool, path: str):
    valley_map, blizzard_map = parse_initial_state(path)
    route_planner = RoutePlanner()

    print('\n== Blizzard Basin ==\n')

    best_time = 0
    total_best_time = 0
    for iteration in range(3 if is_star_two else 1):
        if iteration > 0:
            swap_entry_and_exit(valley_map)

        explorer = valley_map.entry
        best_time = route_planner.get_best_time_to_exit(explorer, valley_map, blizzard_map, start_time=best_time)
        total_best_time += best_time

        print(f'- Best time {iteration + 1}: {best_time}')

    print(f'-Total best time: {total_best_time}')
    print()

def swap_entry_and_exit(valley_map: ValleyMap) -> None:
    temp = valley_map.entry
    valley_map.entry = valley_map.exit
    valley_map.exit = temp

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
    is_star_two = True if sys.argv[2] == 'star2' else False
    path = 'input.test.txt' if is_test else 'input.txt'
    main(is_star_two, path)
