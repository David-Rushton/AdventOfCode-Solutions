import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Range:
    starting: int
    ending: int
    length: int


@dataclass
class LookupRange:
    id: int
    destination_range: int
    source_range: int
    range_length: int


@dataclass
class Lookup:
    name: str
    ranges: list[LookupRange]


def main(is_test_mode: bool) -> None:
    is_test_mode = True
    (seeds, lookups) = get_input(is_test_mode)

    print('seed to soil lookup')
    print('-------------------')
    print()

    seed_ranges: list[Range] = []
    for i, value in enumerate(seeds):
        if i % 2 == 1:
            seed_ranges.append(Range(
                seeds[i - 1],
                seeds[i - 1] + value - 1,
                value))

    for lookup in lookups:
        print(lookup.name)

        temp_range: list[Range] = []
        for seed_range in seed_ranges:
            print_range(seed_range)
            for lu_range in lookup.ranges:
                temp_range.extend(split_range(seed_range, to_range(lu_range)))

        print(f't {len(temp_range)}')
        for r in temp_range:
            print_range(r)


        seed_range = temp_range
        print()


    print()
    print('-------------------')
    print(f'Lowest location number: {min(seeds)}')


def print_range(range: Range) -> None:
    print(f' seed: {range.starting} -> {range.starting +  range.length}')


def print_lookup(range: LookupRange) -> None:
    print(f' range: {range.source_range} -> {range.source_range + range.range_length} : {range.destination_range} -> {range.destination_range +  range.range_length}')


def split_range(source_range: Range, target_range: Range) -> Iterator[Range]:
    # before
    if source_range.starting <= target_range.starting:
        start = source_range.starting
        end = min(source_range.ending, target_range.starting - 1)
        yield Range(start, end, end - source_range.starting)

    # overlap
    if source_range.ending >= target_range.starting and source_range.starting < target_range.ending:
        start = max(source_range.starting, target_range.starting)
        end = min(source_range.ending, target_range.ending)
        yield Range(start, end, end - start)

    # after
    if source_range.ending > target_range.ending:
        start = max(source_range.starting, target_range.ending)
        end = source_range.ending
        yield(start, end, end - start)


def ending(range: Range) -> int:
    return range.starting + range.length - 1


def to_range(lookup_range: LookupRange) -> Range:
    return Range(
        lookup_range.source_range,
        lookup_range.source_range + lookup_range.range_length - 1,
        lookup_range.range_length)


def get_input(get_test: bool) -> tuple[list[int], list[Lookup]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    lines = open(path, 'rt').read().splitlines()

    seeds = [int(number) for number in lines[0].split(':')[1].split(' ') if number.isdigit()]
    lookups: list[Lookup] = []
    id = 0

    for line in lines[1::]:
        if line == '':
            continue

        if line.endswith(':'):
            lookups.append(Lookup(line, []))

        if line[0].isdigit():
            values = [int(number) for number in line.split(' ') if number.isdigit()]
            id += 1
            range = LookupRange(id, values[0], values[1], values[2])
            lookups[-1].ranges.append(range)

    return (seeds, lookups)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
