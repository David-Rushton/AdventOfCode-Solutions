import sys
import itertools
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Report:
    pattern: str
    damaged_groups: list[int]


# wrong: 17613380164860
def main(is_test_mode: bool) -> None:
    print('hot springs')
    print()

    reports = get_input(is_test_mode)
    for report in reports:
        print(f' {report.pattern} | {report.damaged_groups}')

        unfolded_report = unfold_report(report, factor=5)
        segments = segment_string(unfolded_report.pattern, size=5)

        all_segments = itertools.product(segments)
        # all_segments.



        for segment in all_segments:
            print(f'  {segment} ') #: {len(get_arrangements_2(segment))}')

    print()
    # print(f'segments {len(segments)}')


def unfold_report(report: Report, factor: int = 5) -> Report:
    def repeat_pattern() -> Iterator[str]:
        for _ in range(factor):
            yield report.pattern
    return Report(
        '?'.join(repeat_pattern()),
        report.damaged_groups * factor)

    result: list[str] = []
    q: list[str] = ['']
    while len(q) > 0:
        current = q.pop(0)

def segment_string(string: str, size: int=5) -> Iterator[str]:
    buffer = ''
    for i in range(len(string)):
        buffer += string[i]
        if len(buffer) >= size:
            yield buffer
            buffer = ''
    if len(buffer) > 0:
        yield buffer

        next = pattern[len(current)]
        candidates = [next] if next != '?' else ['#', '.']
        for candidate in candidates:
            q.append(f'{current}{candidate}')

cache: dict[str, set[str]] = {}
def get_arrangements_2(pattern: str) -> set[str]:
    global cache

    if pattern in cache:
        return cache[pattern]

    cache[pattern] = set()
    q: list[str] = ['']
    while len(q) > 0:
        current = q.pop(0)

        if len(current) == len(pattern):
            cache[pattern].add(current)
            continue

        next = pattern[len(current)]
        options = ['#', '.'] if next == '?' else next
        for option in options:
            q.append(f'{current}{option}')

    return cache[pattern]


def get_arrangements(pattern: str, damaged_groups: list[int]) -> Iterator[str]:
    q: list[str] = ['']
    while len(q) > 0:
        current = q.pop(0)

        if len(current) == len(pattern):
            current_damaged_groups = [len(s) for s in current.split('.') if len(s) > 0]
            if current_damaged_groups == damaged_groups:
                yield current
            continue

        next = pattern[len(current)]
        options = ['#', '.'] if next == '?' else next
        for option in options:
            q.append(f'{current}{option}')


def get_input(get_test: bool) -> Iterator[Report]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    reports = open(path, 'rt').read().splitlines()
    for report in reports:
        elements = report.split(' ')
        yield Report(
            elements[0],
            [int(number) for number in elements[1].split(',') if number.isdigit()])


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
