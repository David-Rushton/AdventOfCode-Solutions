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
    print(' > AAA')

    (directions, current_location) = get_input(is_test_mode)
    steps = 0
    for direction in get_directions(directions):
        steps += 1
        current_location = current_location.left if direction == 'L' else current_location.right
        print(f' > {current_location.name}')
        if current_location.name == 'ZZZ':
            break;

    print()
    print(f'steps {steps}')


def get_directions(directions: str) -> str:
    while True:
        for direction in directions:
            yield direction


def get_input(get_test: bool) -> tuple[str, Location]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    lines = open(path, 'rt').read().splitlines()

    directions = lines[0]
    locations: dict[str, Location] = {}
    for line in lines[2:]:
        elements = line.replace('=', '').replace('(', '').replace(')', '').replace(',', '').split(' ')
        locations[elements[0]] = Location(elements[0], None, None)

    map: Location = None
    for line in lines[2:]:
        elements = line.replace('=', '').replace('(', '').replace(')', '').replace(',', '').split(' ')
        locations[elements[0]].left = locations[elements[2]]
        locations[elements[0]].right = locations[elements[3]]

    return (directions, locations['AAA'])


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
