import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Report:
    pattern: str
    damaged_groups: list[int]


@dataclass
class Tracker:
    pattern: str
    damaged_groups: list[int]
    weight: int


def main(is_test_mode: bool) -> None:
    print('hot springs')
    print()

    is_test_mode = True

    total_arrangement_count = 0
    reports = get_input(is_test_mode)
    for report in reports:
        unfolded_report = unfold_report(report, size=5)
        print(f' {report.pattern} | ', end='')

        segments = list(split_str(unfolded_report.pattern, size=10))
        arrangement_count = sum(get_arrangement_count(segments, unfolded_report))
        total_arrangement_count += arrangement_count
        print(arrangement_count)

    print()
    print(f'arrangement count: {total_arrangement_count}')


def unfold_report(report: Report, size:int = 5) -> Report:
    unfolded_pattern = []
    for _ in range(size):
        unfolded_pattern.append(report.pattern)
    return Report(
        '?'.join(unfolded_pattern),
        report.damaged_groups * size)


def get_arrangement_count(
        segments: list[str],
        report: Report,
        prefix: str='',
        current_damaged_groups: list[int] = [],
        level: int=0,
        matches: list[int]=[],
        weight: int=1) -> list[int]:
    if level < len(segments):
        segment = segments[level]
        cache: dict[str, Tracker] = {}
        for arrangement in get_arrangements(segment):

            candidate = f'{prefix}{arrangement}'
            arrangement_damaged_groups = get_damaged_groups(arrangement)
            candidate_damaged_groups = add_damaged_groups(prefix, current_damaged_groups, arrangement, arrangement_damaged_groups)

            if len(candidate) == len(report.pattern):
                if is_exact_match(candidate_damaged_groups, report.damaged_groups):
                    matches.append(weight)
            else:
                if is_match(candidate[-1] == '.', candidate_damaged_groups, report.damaged_groups):
                    key = f'{candidate[-1]}-{str(candidate_damaged_groups)}'
                    if key not in cache:
                        cache[key] = Tracker(f'{candidate[0]}{"." * (len(candidate) - 2)}{candidate[-1]}', candidate_damaged_groups, weight)
                    else:
                        cache[key].weight += 1
                    # print(f' {key} {cache[key].damaged_groups} {cache[key].weight}')

        for key in cache:
            item = cache[key]
            cumulative_weight = weight * item.weight
            matches = get_arrangement_count(segments, report, item.pattern, item.damaged_groups, level + 1, matches, item.weight)

    return matches


def add_damaged_groups(
        left_pattern: str,
        left_damaged_groups: list[int],
        right_pattern: str,
        right_damaged_groups: list[int]) -> list[int]:
    if len(left_pattern) == 0:
        return right_damaged_groups
    # .#. + .#. [1] + [1] == [1, 1]
    # ..# + #.. [1] + [1] ==    [1]
    # ..# + .#. [1] + [1] == [1, 1]
    # .#. + #.. [1] + [1] == [1, 1]
    if left_pattern.endswith('#') and right_pattern.startswith('#'):
        new_damaged_groups = left_damaged_groups[:-1]
        new_damaged_groups.append(left_damaged_groups[-1] + right_damaged_groups[0])
        new_damaged_groups.extend(right_damaged_groups[1:])
        return new_damaged_groups
    else:
        return  left_damaged_groups + right_damaged_groups


def is_match(closed: bool, compare: list[int], against: list[int]) -> bool:
    if len(compare) == 0:
        return True
    elif len(compare) > len(against):
        return False
    elif len(compare) == 1 and len(against) > 1:
        return compare[0] <= against[0]
    else:
        start_matches = compare[:-1] == against[:len(compare) - 1]
        if closed:
            last_matches = compare[-1] == against[len(compare) - 1]
        else:
            last_matches = compare[-1] <= against[len(compare) - 1]
        return start_matches and last_matches


def is_exact_match(compare: list[int], against: list[int]) -> bool:
    return compare == against


def split_str(string: str, size: int = 10) -> Iterator[str]:
    start = 0
    end = 0
    while end < len(string):
        start = end
        end = end + size
        end if end <= len(string) else len(string)
        yield string[start:end]


damaged_groups_cache: dict[str, list[int]] = {}
def get_damaged_groups(pattern: str) -> list[int]:
    global damaged_groups_cache
    if pattern in damaged_groups_cache:
        return damaged_groups_cache[pattern]

    result = [len(section) for section in pattern.split('.') if len(section) > 0]
    damaged_groups_cache[pattern] = result
    return result


arrangement_cache: dict[str, list[str]] = {}
def get_arrangements(pattern: str) -> list[str]:
    global arrangement_cache
    if pattern in arrangement_cache:
        return arrangement_cache[pattern]

    result: list[str] = []
    q: list[str] = ['']
    while len(q) > 0:
        current = q.pop(0)

        if len(current) == len(pattern):
            result.append(current)
            continue

        next = pattern[len(current)]
        candidates = [next] if next != '?' else ['#', '.']
        for candidate in candidates:
            q.append(f'{current}{candidate}')

    arrangement_cache[pattern] = result
    return result


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
