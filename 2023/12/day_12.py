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
        print(f' {report}')
        for arrangement in get_arrangements(report.pattern, level=0, result=[]):
            if list(get_damaged_groups(arrangement)) == report.damaged_groups:
                arrangement_count += 1
                print(f'  {arrangement} {arrangement_count}')
        # exit(0)

    print()
    print(f'arrangement count: {arrangement_count}')


def get_arrangements(pattern: str, level: int=0, result: list[str]=[]) -> list[str]:
    if pattern.find('?', level) == -1:
        result.append(pattern)
    else:
        get_arrangements(pattern.replace('?', '.', 1), level + 1, result)
        get_arrangements(pattern.replace('?', '#', 1), level + 1, result)
    return result


def get_damaged_groups(pattern: str) -> Iterator[int]:
    for group in pattern.split('.'):
        damaged_count = group.count('#')
        if damaged_count > 0:
            yield damaged_count


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
