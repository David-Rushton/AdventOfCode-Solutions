import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Operation:
    label: str
    operator: str
    focal_length: int


@dataclass
class Box:
    id: int
    lens: dict[str, int]


def main(is_test_mode: bool) -> None:
    print('lens library')
    print()

    boxes: dict[int, Box] = {}
    operations = get_input(is_test_mode)

    for operation in operations:
        box_id = get_hash(operation.label)
        if box_id not in boxes:
            boxes[box_id] = Box(box_id, {})
        if operation.operator == '-':
            if operation.label in boxes[box_id].lens:
                boxes[box_id].lens.pop(operation.label)
        else:
            # operation.operator == '='
            boxes[box_id].lens[operation.label] = operation.focal_length

    print_boxes(boxes)
    score = score_boxes(boxes)

    print()
    print(f'score of steps {score}')


def score_boxes(boxes: list[Box]) -> int:
    result = 0
    box_number = 1
    for box in boxes:
        factor = 1
        box_number = box + 1
        for lens_key in boxes[box].lens:
            subtotal = (box_number * factor * boxes[box].lens[lens_key])
            print(f' {lens_key}: {box_number} (box {box}) * {factor} (slot) * {boxes[box].lens[lens_key]} (focal length) = {subtotal}')
            result += subtotal
            factor += 1
    print()

    return result


def print_boxes(boxes: list[Box]) -> None:
    for box in boxes:
        print(f' Box {box} ', end='')
        for lens in boxes[box].lens:
            print(f'[{lens} {boxes[box].lens[lens]}] ', end='')
        print()
    print()


def get_hash(word: str):
    result = 0
    for char in word:
        result += ord(char)
        result *= 17
        result = result % 256
    return result


def get_input(get_test: bool) -> Iterator[Operation]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    for word in open(path, 'rt').read().replace('\n', '').replace('\t', '').replace(' ', '').split(','):
        if word.endswith('-'):
            yield Operation(word[:-1], '-', 0)
        else:
            split = word.index('=')
            yield Operation(word[0:split], word[split:split + 1], int(word[split + 1:]))


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
