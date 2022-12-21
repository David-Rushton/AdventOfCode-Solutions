from __future__ import annotations
from ast import Tuple
from dataclasses import dataclass
import math
import sys


IS_VERBOSE_MODE = True

@dataclass
class MixNumber:
    id: int
    value: int
    last: MixNumber
    next: MixNumber


def main(is_test_mode: bool, is_star_two: bool, path: str):
    elements = read_mix_file(is_star_two, path)
    indexed_by_starting_position: dict[int, MixNumber] = elements[0]
    mix_numbers: list[MixNumber] = elements[1]
    original_order = mix_numbers.copy()
    round = 0
    round_limit = 1 if is_star_two else 1

    print('== Decrypting Mix File ==\n')

    for round in range(round_limit):
        print(f'  - Round: {round +  1}')
        print(f'  - Initial: {original_order[0].value}, {original_order[1].value}, {original_order[2].value}, {original_order[3].value}, {original_order[4].value}, {original_order[5].value}, {original_order[6].value}\n')

        for index in range(len(original_order)):

            print(f'  - Decrypted item {index + 1} of {len(mix_numbers)}: {original_order[index].value}')

            source = original_order[index]
            source_index = get_current_index(source, mix_numbers, index)
            target_index = get_index_offset(source_index, source.value, len(mix_numbers))
            target = mix_numbers[target_index]

            print(f'    - Source index {source_index}')
            print(f'    - Target index {target_index}')

            # reorder array
            if source.value != 0:
                mix_numbers.insert(target_index, mix_numbers.pop(source_index))

            print(f'    - New order: ', end='')
            for x in mix_numbers:
                print(x.value, end=', ')

            print()
            print()

        coordinates = 0
        index_of_zero = get_mix_number_zero_index(mix_numbers)
        print(f'  - Searching for coordinates')

        for offset in (1000, 2000, 3000):
            index = get_index_offset(index_of_zero, offset, len(mix_numbers), disable_wraps=True)
            print(f'  - {offset}th number after 0 is {mix_numbers[index].value} ({index})')
            coordinates += mix_numbers[index].value

        print(f'  - Coordinates {coordinates}\n')

    exit(0)

def get_mix_number_zero_index(mix_numbers: list[MixNumber]) -> int:
    for index in range(len(mix_numbers)):
        if mix_numbers[index].value == 0:
            return index

def get_index_offset(index: int, offset: int, array_length: int, disable_wraps: bool=False) -> int:

    if disable_wraps:
        wraps = 0
    else:
        wraps = math.floor((index + offset) / array_length)

    result = (index + offset + wraps) % array_length

    if result == 0 and offset < 0:
        result = array_length - 1

    return result

def get_current_index(mix_number: MixNumber, mix_numbers: list[MixNumber], start_from_index: int):
    offset_seed = 0
    while True:
        for factor in (1, -1):
            offset = ((start_from_index + offset_seed) * factor) % len(mix_numbers)
            if mix_numbers[offset] == mix_number:
                return offset
            offset_seed += 1

            if offset_seed % 100 == 0:
                print(f'      - Large offset detected: {offset_seed}')

def read_mix_file(is_star_two: bool, path: str) -> Tuple[dict[int, MixNumber], list[MixNumber]]:

    indexed_by_starting_position: dict[int, MixNumber] = {}
    mix_numbers: list[MixNumber] = []
    id_seed = -1
    decryption_key = 811589153 if is_star_two else 1
    last = None

    for value in open(path, 'r').read().splitlines():
        id_seed += 1

        current = MixNumber(
            id = id_seed,
            value = int(value) * decryption_key,
            last = last,
            next = None
        )

        indexed_by_starting_position[id_seed] = current
        mix_numbers.append(current)

        if last is not None:
            last.next = current

        last = current

    indexed_by_starting_position[id_seed].next = indexed_by_starting_position[0]
    indexed_by_starting_position[0].last = indexed_by_starting_position[id_seed]

    return (indexed_by_starting_position, mix_numbers)

def print_ex(message: str = ''):
    if IS_VERBOSE_MODE:
        print(message)

if __name__ == '__main__':
    is_test_mode = sys.argv[1] == 'test'
    is_star_two = sys.argv[2] == 'star2'
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, is_star_two, path)
