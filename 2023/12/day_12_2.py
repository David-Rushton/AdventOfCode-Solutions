import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Report:
    pattern: str
    damaged_groups: list[int]


def main(is_test_mode: bool) -> None:
    print('hot springs')
    print()

    arrangement_count = 0
    reports = get_input(is_test_mode)
    print('  a   b  extend')
    print(' --- --- ------')
    for report in reports:
        a = get_arrangement_count(f'{report.pattern}', report.damaged_groups)
        if report.pattern.endswith('.'):
            b = get_arrangement_count(f'?{report.pattern}', report.damaged_groups)
        else:
            b = get_arrangement_count(f'{report.pattern}?', report.damaged_groups)

        print(f' {str(len(a)).rjust(3)} {str(len(b)).rjust(3)} {str(len(a) * len(b) * len(b) * len(b) * len(b)).rjust(6)}')
        arrangement_count += len(a) * len(b) * len(b) * len(b) * len(b)

    print()
    print(f'arrangement count: {arrangement_count}')


def get_arrangement_count(pattern: str, damaged_groups: list[int]) -> list[str]:
    result: list[str] = []

    q: list[str] = ['']
    sum_damaged_groups = sum(damaged_groups)
    total_matches = 0
    iterations = 0
    while len(q) > 0:
        iterations += 1
        current = q.pop(0)

        if len(current) == len(pattern):
            if equals_damaged_groups(current, damaged_groups):
                result.append(current)
                total_matches += 1
            continue

        next = pattern[len(current)]
        if next == '?':
            options = ['#', '.']
        else:
            options = [next]
        for option in options:
            candidate = f'{current}{option}'

            damaged = f'{candidate}{pattern[len(candidate):]}'.count('#')
            if damaged > sum_damaged_groups:
                continue

            if matches_damaged_groups(candidate, damaged_groups, len(pattern)):
                q.append(candidate)

    return result


def matches_damaged_groups(pattern: str, damaged_groups: list[int], limit: int) -> bool:
    pattern_damaged_groups = ([len(s) for s in pattern.split('.') if len(s) > 0])
    required = len(pattern) + sum(damaged_groups[len(pattern_damaged_groups):]) + len(damaged_groups[len(pattern_damaged_groups):]) - 1
    if required > limit:
        return False

    if len(pattern_damaged_groups) == 0:
        return True
    elif len(pattern_damaged_groups) > len(damaged_groups):
        return False
    elif len(pattern_damaged_groups) == 1:
        return pattern_damaged_groups[0] <= damaged_groups[0]
    else:
        return pattern_damaged_groups[:-1] == damaged_groups[:len(pattern_damaged_groups) - 1] \
            and pattern_damaged_groups[-1] <= damaged_groups[len(pattern_damaged_groups) - 1]


def equals_damaged_groups(pattern: str, damaged_groups: list[int]) -> bool:
    current_damaged_groups = ([len(s) for s in pattern.split('.') if len(s) > 0])
    return current_damaged_groups == damaged_groups


def unfold_report(report: Report) -> Report:
    return Report(
        '?'.join([report.pattern, report.pattern]), #report.pattern, report.pattern, report.pattern]),
        report.damaged_groups * 2)


def get_damaged_groups(pattern: str) -> Iterator[int]:
    current_count = 0
    for char in pattern:
        if char == '#':
            current_count += 1
        if char == '.':
            if current_count > 0:
                yield current_count
                current_count = 0
        if char == '?':
            return
    if current_count > 0:
        yield current_count


def print_reports(reports: list[Report]):
    for report in reports:
        print(f'{report.pattern} {report.damaged_groups}')


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
