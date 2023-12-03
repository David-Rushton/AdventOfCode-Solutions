import sys
from dataclasses import dataclass
from enum import Enum
from typing import Iterator


class CellType(Enum):
    DIGIT = 1
    SPACE = 2
    SYMBOL = 3
    END_OF_LINE = 4


@dataclass(frozen=True)
class Point:
    x: int
    y: int


@dataclass
class Cell:
    point: Point
    cellType: CellType
    value: str


def main(is_test_mode: bool) -> None:
    part_number_points: list[Point] = []
    gear_points: dict[Point, Point] = {}
    for cell in get_input(is_test_mode):
        if cell.cellType == CellType.SYMBOL:
            part_number_points.extend(get_neighbours(cell.point))
            if cell.value == '*':
                for neighbour in get_neighbours(cell.point):
                    gear_points[neighbour] = cell.point

    buffer = ''
    is_part_number = False
    total_of_part_numbers = 0
    gear_location: Point = None
    gear_parts: dict[Point, list[int]] = {}
    for cell in get_input(is_test_mode):
        if cell.cellType == CellType.DIGIT:
            buffer += cell.value
            if cell.point in part_number_points:
                is_part_number = True
            if cell.point in gear_points:
                gear_location = gear_points[cell.point]
        elif len(buffer) > 0:
            print(f'{buffer} {is_part_number} {gear_location is None}')
            if is_part_number:
                total_of_part_numbers += int(buffer)
            if gear_location is not None:
                if gear_location not in gear_parts:
                    gear_parts[gear_location] = []
                gear_parts[gear_location].append(int(buffer))
            buffer = ''
            is_part_number = False
            gear_location = None

    print('----------')
    print(f'total of part numbers: {total_of_part_numbers}')

    gear_ratios = 0
    for gear_part in gear_parts:
        if len(gear_parts[gear_part]) == 2:
            gear_ratios += gear_parts[gear_part][0] * gear_parts[gear_part][1]
    print(gear_parts)
    print(f'gear ratios: {gear_ratios}')

def get_neighbours(point: Point) -> Iterator[Point]:
    for y in range(point.y - 1, point.y + 2):
        for x in range(point.x - 1, point.x + 2):
            yield Point(x, y)


def get_input(get_test: bool) -> Iterator[Cell]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path

    y = 0
    for line in open(path, 'rt').read().splitlines():
        x = 0
        for char in line:
            point = Point(x, y)
            if char == '.':
                yield Cell(point, CellType.SPACE, char)
            elif char.isdigit():
                yield Cell(point, CellType.DIGIT, char)
            else:
                yield Cell(point, CellType.SYMBOL, char)
            x +=1
        yield Cell(Point(x, y), CellType.END_OF_LINE, '')
        y += 1


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
