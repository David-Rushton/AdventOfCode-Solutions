from enum import Enum
import sys
from dataclasses import dataclass
from typing import Iterator


class Direction(Enum):
    NORTH=0
    EAST=1
    SOUTH=2
    WEST=3


@dataclass(frozen=True)
class Point:
    x: int
    y: int


@dataclass
class Cell:
    point: Point
    value: str
    beams: list[Direction]


def main(is_test_mode: bool) -> None:
    print('the floor will be lava')
    print()

    (max_x, max_y, map) = get_input(is_test_mode)
    max_energised = 0

    for (point, direction) in get_starting_positions(max_x, max_y):
        reset_map(map)
        plot_beam(map, point, direction)
        energised = sum(min(1, len(map[point].beams)) for point in map)
        print(f' {point.x}x{point.y} {direction} {energised}')
        if energised > max_energised:
            max_energised = energised

    print()
    print(f'max energised cells: {max_energised}')


def get_starting_positions(max_x: int, max_y: int) -> tuple[Point, Direction]:
    for y in range(0, max_y):
        yield (Point(-1, y), Direction.EAST)
    for y in range(0, max_y):
        yield (Point(max_x, y), Direction.WEST)
    for x in range(0, max_x):
        yield (Point(x, -1), Direction.SOUTH)
    for x in range(0, max_x):
        yield (Point(x, max_y), Direction.NORTH)

def reset_map(map: dict[Point, Cell]):
    for point in map:
        map[point].beams = []

def plot_beam(
        map: dict[Point, Cell],
        starting_point: Point=Point(-1, 0),
        starting_direction: Direction=Direction.EAST) -> None:
    beams: list[tuple[Point, Direction]] = [(starting_point, starting_direction)]
    while len(beams) > 0:
        (point, direction) = move_beam(beams.pop())
        if point in map:
            if direction in map[point].beams:
                continue
            map[point].beams.append(direction)
            if map[point].value == '.':
                beams.append((point, direction))
            elif map[point].value == '-':
                if direction in [Direction.NORTH, Direction.SOUTH]:
                    beams.append((point, Direction.EAST))
                    beams.append((point, Direction.WEST))
                else:
                    beams.append((point, direction))
            elif map[point].value == '|':
                if direction in [Direction.EAST, Direction.WEST]:
                    beams.append((point, Direction.NORTH))
                    beams.append((point, Direction.SOUTH))
                else:
                    beams.append((point, direction))
            elif map[point].value == '/':
                beams.append((point, rotate_anticlockwise(direction)))
            elif map[point].value == '\\':
                beams.append((point, rotate_clockwise(direction)))


# \
def rotate_clockwise(direction: Direction) -> Direction:
    if direction == Direction.NORTH:
        return Direction.WEST
    if direction == Direction.EAST:
        return Direction.SOUTH
    if direction == Direction.SOUTH:
        return Direction.EAST
    # WEST
    return Direction.NORTH

# /
def rotate_anticlockwise(direction: Direction) -> Direction:
    if direction == Direction.NORTH:
        return Direction.EAST
    if direction == Direction.WEST:
        return Direction.SOUTH
    if direction == Direction.SOUTH:
        return Direction.WEST
    # EAST
    return Direction.NORTH


def move_beam(beam: tuple[Point, Direction]) -> tuple[Point, Direction]:
    point, direction = beam
    if direction == Direction.NORTH:
        return (Point(point.x, point.y - 1), direction)
    if direction == Direction.EAST:
        return (Point(point.x + 1, point.y), direction)
    if direction == Direction.SOUTH:
        return (Point(point.x, point.y + 1), direction)
    # WEST
    return (Point(point.x - 1, point.y), direction)


def print_map(map: dict[Point, Cell]):
    last_y = 0
    for point in map:
        if point.y > last_y:
            last_y = point.y
            print()
        value = '#' if len(map[point].beams) > 0 else map[point].value
        print(value, end='')
    print()


def get_input(get_test: bool) -> tuple[int, int, dict[Point, Cell]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    rows = open(path, 'rt').read().splitlines()
    result: dict[Point, Cell] = {}
    for y in range(0, len(rows)):
        for x in range(0, len(rows[y])):
            point = Point(x, y)
            cell = Cell(point, rows[y][x], [])
            result[point] = cell

    return (len(rows[0]), len(rows), result)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
