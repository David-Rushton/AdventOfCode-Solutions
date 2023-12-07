import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class LookupRange:
    destination_range: int
    source_range: int
    range_length: int


@dataclass
class Lookup:
    name: str
    ranges: list[LookupRange]


def main(is_test_mode: bool) -> None:
    (seeds, lookups) = get_input(is_test_mode)

    print('seed to soil lookup')
    print('-------------------')

    for lookup in lookups:
        print(f'applying {lookup.name}')
        seeds_found: list[int] = []
        next: list[int] = []
        for seed in seeds:
            for range in lookup.ranges:
                if seed >= range.source_range and seed < range.source_range + range.range_length:
                    next.append(range.destination_range + (seed - range.source_range))
                    seeds_found.append(seeds.index(seed))
        for index in seeds_found[::-1]:
            seeds.pop(index)
        seeds.extend(next)

    print('-------------------')
    print(f'Lowest location number: {min(seeds)}')


def get_input(get_test: bool) -> tuple[list[int], list[Lookup]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    lines = open(path, 'rt').read().splitlines()

    seeds = [int(number) for number in lines[0].split(':')[1].split(' ') if number.isdigit()]
    lookups: list[Lookup] = []

    for line in lines[1::]:
        if line == '':
            continue

        if line.endswith(':'):
            lookups.append(Lookup(line, []))

        if line[0].isdigit():
            values = [int(number) for number in line.split(' ') if number.isdigit()]
            range = LookupRange(values[0], values[1], values[2])
            lookups[-1].ranges.append(range)

    return (seeds, lookups)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
