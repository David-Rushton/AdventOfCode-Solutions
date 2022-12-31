from dataclasses import dataclass
from typing import Generator
import math
import sys


SNAFU_DIGITS = {
    '2': 2,
    '1': 1,
    '0': 0,
    '-': -1,
    '=': -2
}

DECIMAL_DIGITS = {
     2: '2',
     1: '1',
     0: '0',
    -1: '-',
    -2: '='
}

@dataclass
class SnafuNumber:
    snafu: str
    decimal: int
    hint: str


def main(is_test_mode: bool, path: str):
    snafu_numbers = list(parse_snafu_numbers(path))

    for snafu_number in snafu_numbers:
        snafu_number.decimal = convert_from_snafu(snafu_number.snafu)
        snafu_number.hint = convert_to_snafu(snafu_number.decimal)

    print_snafu_numbers(snafu_numbers)

def convert_to_snafu(decimal: int) -> str:
    multiplier = 1
    decimal_digits = []

    while multiplier * 2 < decimal:
        multiplier *= 5

    running_total = 0
    remaining = decimal
    while multiplier >= 1:
        if remaining > 0:
            decimal_digit = math.floor(max(multiplier, remaining) / min(multiplier, remaining))
        else:
            decimal_digit = math.floor(min(multiplier, remaining) / max(multiplier, remaining))
        decimal_digits.append(DECIMAL_DIGITS[decimal_digit])
        running_total += multiplier * decimal_digit
        remaining = decimal - running_total
        multiplier /= 5


    return ''.join(decimal_digits[::-1])

def convert_from_snafu(snafu: str) -> int:
    multiplier = 1
    result = 0

    for snafu_digit in snafu[::-1]:
        result += SNAFU_DIGITS[snafu_digit] * multiplier
        multiplier *= 5

    return result

def print_snafu_numbers(snafu_numbers: list[SnafuNumber]) -> None:
    print()
    print('+--------------------+--------------------+--------------------+')
    print('| decimal            | snafu              | hint               |')
    print('+--------------------+--------------------+--------------------+')

    for snafu_number in snafu_numbers:
        print(f'| {str(snafu_number.decimal).rjust(18)} | {snafu_number.snafu.rjust(18)} | {snafu_number.hint.rjust(18)} |')

    print('+--------------------+--------------------+--------------------+')
    print()

def parse_snafu_numbers(path: str) -> Generator[SnafuNumber, None, None]:
    for snafu_number in open(path, 'r').read().splitlines():
        elements = snafu_number.split(',')
        if len(elements) == 2:
            yield SnafuNumber(snafu=elements[0], decimal=None, hint=elements[1])
        else:
            yield SnafuNumber(snafu=elements[0], decimal=None, hint='')


if __name__ == '__main__':
    is_test_mode = True if sys.argv[1] == 'test' else False
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, path)
