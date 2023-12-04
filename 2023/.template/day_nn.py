import sys
from dataclasses import dataclass
from typing import Iterator


def main(is_test_mode: bool) -> None:
    input = get_input(is_test_mode)
    for line in input:
        print(line)
    pass


def get_input(get_test: bool) -> str:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    return open(path, 'rt').read().splitlines()


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
