import sys
from dataclasses import dataclass
from typing import Iterator


def main(is_test_mode: bool) -> None:
    print('mirage maintenance')
    print()

    input = get_input(is_test_mode)
    running_total = 0
    lines = 0
    for line in input:
        numbers = run_numbers([line])
        extend_numbers(numbers)
        print_numbers(numbers)
        running_total += numbers[0][-1]
        lines += 1

    print()
    print(f'total {running_total} ({lines})')


def print_numbers(numbers: list[list[int]]):
    indent = 4
    for row in numbers:
        print(' ' * indent, end='')
        for number in row:
            print(str(number).rjust(8), end='')
        print()
        indent += 4
    print()


def run_numbers(numbers: list[list[int]], level: int=1) -> list[list[int]]:
    numbers.append([])
    go_deeper = False
    for i in range(1, len(numbers[-2])):
        next_number = numbers[-2][i] - numbers[-2][i -1]
        numbers[-1].append(next_number)
        if go_deeper == False and next_number != 0:
            go_deeper = True

    if go_deeper:
        return run_numbers(numbers, level + 1)
    else:
        return numbers


def extend_numbers(numbers: list[list[int]]):
    numbers[-1].append(0)
    for row in range(len(numbers) - 2, -1, -1):
        next_number = numbers[row][-1] + numbers[row + 1][-1]
        numbers[row].append(next_number)


def get_input(get_test: bool) -> list[list[int]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    for line in open(path, 'rt').read().splitlines():
        yield [int(number) for number in line.split(' ')]


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
