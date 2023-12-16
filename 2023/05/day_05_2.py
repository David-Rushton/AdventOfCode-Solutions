import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Range:
    start: int
    end: int


@dataclass
class LookupRange:
    source: Range
    destination: Range


@dataclass
class Lookup:
    name: str
    ranges: list[LookupRange]


def main(is_test_mode: bool) -> None:
    print('seed to soil lookup')
    print('-------------------')
    print()

    (seeds, lookups) = get_input(is_test_mode)

    print(' seeds:')
    for seed in sorted(seeds, key=lambda s: s.start):
        print_range(seed)
    print()

    for lookup in lookups:
        print(f' {lookup.name}')
        seeds_updated: list[Range] = []
        for lookup_range in lookup.ranges:
            print_lookup_range(lookup_range)
            seeds_to_check = seeds.copy()
            seeds_checked: list[Range] = []
            while(len(seeds_to_check) > 0):
                seed = seeds_to_check.pop()
                for (seed_intersection, is_updated) in get_intersections(seed, lookup_range):
                    if is_updated:
                        seeds_updated.append(seed_intersection)
                    else:
                        seeds_checked.append(seed_intersection)
            seeds = seeds_checked
        print(' seeds:')
        seeds.extend(seeds_updated)
        for seed in sorted(seeds, key=lambda r: r.start):
            print_range(seed)
        print()

    print('-------------------')
    print(f'Lowest location number: {min(seed.start for seed in seeds)}')


def print_lookup_range(range: LookupRange) -> None:
    print(f'  {range.source.start} -> {range.source.end} = {range.destination.start} -> {range.destination.end}')


def print_range(range: Range, end: str='\n') -> None:
    print(f'  {range.start} -> {range.end}', end=end)


def get_intersections(seeds: Range, lookup_range: LookupRange) -> Iterator[tuple[Range, bool]]:
    left = seeds
    right = lookup_range.source
    if left.start < right.start:
        yield (Range(left.start, min(left.end, right.start - 1)), False)
    if left.end > right.start and left.start < right.end:
        new_range = Range(max(left.start, right.start), min(left.end, right.end))
        yield (apply_lookup_range(new_range, lookup_range), True)
    if left.end >= right.end:
        yield (Range(max(left.start, right.end), max(left.end, right.end)), False)


def apply_lookup_range(range: Range, lookup_range: LookupRange) -> Range:
    offset = range.start - lookup_range.source.start
    width = range.end - range.start
    r = Range(
        lookup_range.destination.start + offset,
        lookup_range.destination.start + offset + width)
    print(f'    {range.start}x{range.end} -> {r.start}x{r.end}')
    return r


def get_input(get_test: bool) -> tuple[list[Range], list[Lookup]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    lines = open(path, 'rt').read().splitlines()

    seeds = [int(number) for number in lines[0].split(':')[1].split(' ') if number.isdigit()]
    seed_ranges: list[Range] = []
    for i in range(1, len(seeds), 2):
        seed_ranges.append(Range(seeds[i - 1], seeds[i - 1] + seeds[i]))

    lookups: list[Lookup] = []
    for line in lines[1::]:
        if line == '':
            continue
        if line.endswith(':'):
            lookups.append(Lookup(line, []))
        if line[0].isdigit():
            values = [int(number) for number in line.split(' ') if number.isdigit()]
            lookups[-1].ranges.append(
                LookupRange(
                    Range(values[1], values[1] + values[2]),
                    Range(values[0], values[0] + values[2])))

    return (seed_ranges, lookups)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
