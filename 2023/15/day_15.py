import sys
from dataclasses import dataclass
from typing import Iterator


def main(is_test_mode: bool) -> None:
    print('lens library')
    print()

    words = get_input(is_test_mode)
    steps_total = 0
    for word in words:
        word_total = get_hash(word)
        print(f'  {word} {word_total}')
        steps_total += word_total

    print()
    print(f'sum of steps {steps_total}')


def get_hash(word: str):
    result = 0
    for char in word:
        result += ord(char)
        result *= 17
        result = result % 256
    return result


def get_input(get_test: bool) -> Iterator[str]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    for word in open(path, 'rt').read().replace('\n', '').replace('\t', '').replace(' ', '').split(','):
        yield word


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
