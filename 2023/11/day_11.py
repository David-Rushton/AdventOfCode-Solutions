import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass(eq=True,frozen=True)
class Galaxy:
    x: int
    y: int


@dataclass
class Space:
    location: Galaxy
    distance: int


@dataclass
class Universe:
    width: int
    height: int
    galaxies: dict[Galaxy, int]


def main(is_test_mode: bool) -> None:
    universe = get_input(is_test_mode)

    print('cosmic expansion')
    print()
    # print_universe(universe)

    sum_of_lengths = 0
    for (left_galaxy, right_galaxy) in get_unique_galaxy_pairs(universe):
        length = get_shortest_path(universe, left_galaxy, right_galaxy)
        print(f' {left_galaxy} ->  {right_galaxy} == {length}')
        sum_of_lengths += length

    print()
    print(f'sum of lengths: {sum_of_lengths}')


def print_universe(universe: Universe) -> None:
    for y in range(0, universe.height):
        for x in range(0, universe.width):
            print('#' if Galaxy(x, y) in universe.galaxies else '.', end='')
        print()


def get_shortest_path(universe: Universe, left_galaxy: Galaxy, right_galaxy: Galaxy) -> int:
    return (max(left_galaxy.x, right_galaxy.x) - min(left_galaxy.x, right_galaxy.x)) + (max(left_galaxy.y, right_galaxy.y) - min(left_galaxy.y, right_galaxy.y))


def get_unique_galaxy_pairs(universe: Universe) -> Iterator[tuple[Galaxy, Galaxy]]:
    keys = set()
    count = 0
    for outer_galaxy in universe.galaxies:
        for inner_galaxy in universe.galaxies:
            if outer_galaxy != inner_galaxy:
                key = f'{min(universe.galaxies[outer_galaxy], universe.galaxies[inner_galaxy])}x{max(universe.galaxies[outer_galaxy], universe.galaxies[inner_galaxy])}'
                if not key in keys:
                    count += 1
                    keys.add(key)
                    yield (outer_galaxy, inner_galaxy)


def get_input(get_test: bool) -> Universe:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    universe = open(path, 'rt').read().splitlines()
    empty_rows = list(get_empty_rows(universe))
    y_offset = 0
    galaxies: dict[Galaxy, int] = {}
    galaxy_next_id = 1
    for y in range(0, len(universe)):
        galaxies_found = 0
        x_offset = 0
        for x in range(0, len(universe[y])):
            if x in empty_rows:
                x_offset += 999999
            if universe[y][x] == '#':
                galaxies[Galaxy(x + x_offset, y + y_offset)] = galaxy_next_id
                galaxy_next_id += 1
                galaxies_found +=1
        if galaxies_found == 0:
            y_offset += 999999

    return Universe(
        width=len(universe[0]) + len(empty_rows),
        height=len(universe) + y_offset,
        galaxies=galaxies)


def get_empty_rows(universe: list[str]) -> Iterator[int]:
    for x in range(0, len(universe[0])):
        galaxies_found = 0
        for y in range(0, len(universe)):
            if universe[y][x] == '#':
                galaxies_found += 1
        if galaxies_found == 0:
            yield x


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
