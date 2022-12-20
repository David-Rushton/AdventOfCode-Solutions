from __future__ import annotations
from ast import Tuple
from dataclasses import dataclass
import sys

@dataclass
class MixNumber:
    id: int
    value: int
    original_position: int
    next: MixNumber


def main(is_test_mode: bool, path: str):
    elements = read_mix_file(path)
    indexed_by_starting_position: dict[int, MixNumber] = elements[0]
    mix_numbers: list[MixNumber] = elements[1]
    original_order = mix_numbers.copy()

    for index in range(len(original_order)):
        print(f'Value = {mix_numbers[index].value} | Next Value = {mix_numbers[index].next.value}')


def read_mix_file(path: str) -> Tuple[dict[int, MixNumber], list[MixNumber]]:

    indexed_by_starting_position: dict[int, MixNumber] = {}
    mix_numbers: list[MixNumber] = []
    id_seed = -1
    last = None

    for value in open(path, 'r').read().splitlines():
        id_seed += 1

        current = MixNumber(
            id = id_seed,
            value = int(value),
            original_position = id_seed,
            next = None
        )

        indexed_by_starting_position[id_seed] = current
        mix_numbers.append(current)

        if last is not None:
            last.next = current

        last = current

    indexed_by_starting_position[id_seed].next = indexed_by_starting_position[0]

    return (indexed_by_starting_position, mix_numbers)


if __name__ == '__main__':
    is_test_mode = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, path)
