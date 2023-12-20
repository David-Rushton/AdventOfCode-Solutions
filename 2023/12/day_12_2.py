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
    for report in reports:
        unfolded_report = unfold_report(report)
        print(f' {report.pattern}', end='')

        for arrangement in get_arrangements(unfolded_report.pattern, unfolded_report.damaged_groups, level=0, result=[]):
            if list(get_damaged_groups(arrangement)) == unfolded_report.damaged_groups:
                arrangement_count += 1

        print(f' {arrangement_count}')

    print()
    print(f'arrangement count: {arrangement_count}')


def get_arrangements(pattern: str, damaged_groups: list[int], level: int=0, result: list[str]=[]) -> list[str]:
    current_damaged_groups = list(get_damaged_groups(pattern))
    if len(current_damaged_groups) > 0:
        if current_damaged_groups != damaged_groups[0:len(current_damaged_groups)]:
            return result

    if pattern.find('?', level) == -1:
        result.append(pattern)
    else:
        get_arrangements(pattern.replace('?', '#', 1), damaged_groups, level + 1, result)
        get_arrangements(pattern.replace('?', '.', 1), damaged_groups, level + 1, result)

    return result


def unfold_report(report: Report) -> Report:
    return Report(
        '?'.join([report.pattern, report.pattern, report.pattern, report.pattern, report.pattern]),
        report.damaged_groups * 5)


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
