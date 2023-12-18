import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass(frozen=True)
class Point:
    x: int
    y: int


@dataclass
class Instruction:
    direction: str
    distance: int
    colour: str


def main(is_test_mode: bool) -> None:
    print('lavaduct lagoon')
    print()


    moves = {'U': (0, -1), 'L': (-1, 0), 'D': (0, 1), 'R': (1, 0)}
    point = Point(0, 0)
    points: set[Point] = set([point])
    instructions = get_input(is_test_mode)



    print(' calculating outline')
    for instruction in instructions:
        instruction = convert_hex(instruction)
        offset = moves[instruction.direction]
        for _ in range(instruction.distance):
            point = Point(point.x + offset[0], point.y + offset[1])
            if point not in points:
                points.add(point)

    print(' filling lagoon')
    fill_lagoon(points)

    print_points(points)

    print()
    print(f'lagoon size {len(points)}m2')


def convert_hex(instruction: Instruction) -> Instruction:
    directions: dict[str, str] = {'0': 'R', '1': 'D', '2': 'L', '3': 'U'}
    return Instruction(
        directions[instruction.colour[-1]],
        int(instruction.colour[1:-1], 16),
        instruction.colour)


def fill_lagoon(points: set[Point]) -> None:
    q: list[Point] = [Point(1, 1)]
    while len(q) > 0:
        point = q.pop()
        for offset in [(-1 , -1), (0, -1), (1, -1), (-1, 0), (1, 0), (-1, 1), (0, 1), (1, 1)]:
            next_point = Point(point.x + offset[0], point.y + offset[1])
            if next_point not in points:
                q.append(next_point)

        if point not in points:
            points.add(point)


def print_points(points: set[Point]) -> None:
    min_x = min(point.x for point in points)
    max_x = max(point.x for point in points)
    min_y = min(point.y for point in points)
    max_y = max(point.y for point in points)

    for y in range(min_y, max_y + 1):
        for x in range(min_x, max_x + 1):
            value = '#' if Point(x, y) in points else '.'
            if (x, y) == (1, 1):
                value = '*'
            print(value, end='')
        print()


def get_input(get_test: bool) -> Iterator[Instruction]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    instructions = open(path, 'rt').read().splitlines()
    for instruction in instructions:
        elements = instruction.split(' ')
        yield Instruction(elements[0], int(elements[1]), elements[2].replace('(', '').replace(')', ''))


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
