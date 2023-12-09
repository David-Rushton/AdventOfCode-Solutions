from __future__ import annotations
from dataclasses import dataclass
from typing import Iterator
import sys


@dataclass
class Location:
    name: str
    left: Location
    right: Location


def main(is_test_mode: bool) -> None:
    print('haunted wasteland')
    print()

    (directions, current_locations) = get_input(is_test_mode)

    for i in range(0, len(current_locations)):
        steps = 0
        for direction in get_directions(directions):
            steps += 1
            current_locations[i] = current_locations[i].left if direction == 'L' else current_locations[i].right
            if current_locations[i].name.endswith('Z'):
                break
        print(f' {i} {steps} {steps / 281}')

    print()
    print(f'steps {steps}')


def get_directions(directions: str) -> str:
    while True:
        for direction in directions:
            yield direction


def get_input(get_test: bool) -> tuple[str, list[Location]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    lines = open(path, 'rt').read().splitlines()

    directions = lines[0]
    locations: dict[str, Location] = {}
    starting_locations: list[Location] = []
    for line in lines[2:]:
        elements = line.replace('=', '').replace('(', '').replace(')', '').replace(',', '').split(' ')
        locations[elements[0]] = Location(elements[0], None, None)

    map: Location = None
    for line in lines[2:]:
        elements = line.replace('=', '').replace('(', '').replace(')', '').replace(',', '').split(' ')
        locations[elements[0]].left = locations[elements[2]]
        locations[elements[0]].right = locations[elements[3]]
        if elements[0].endswith('A'):
            starting_locations.append(locations[elements[0]])

    return (directions, starting_locations)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
