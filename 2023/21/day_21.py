import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass(frozen=True)
class Point:
    x: int
    y: int


def main(is_test_mode: bool) -> None:
    print('step counter')
    print()

    (start, map) = get_input(is_test_mode)
    q: list[tuple[int, Point]] = [(0, start)]
    destinations: list[Point] = []
    visited: list[tuple(int, Point)] = []
    target = 6 if is_test_mode else 64

    while len(q) > 0:
        (steps, current) = q.pop(0)
        steps += 1

        print(f'checking {current.x}x{current.y} for step {steps}, with {len(q)} remaining to check', end='\r')
        for neighbour in get_neighbours(current):
            if neighbour in map:
                if map[neighbour] in ['S', '.']:
                    if (steps, neighbour) in visited:
                        continue
                    visited.append((steps, neighbour))
                    if steps == target:
                        if neighbour not in destinations:
                            destinations.append(neighbour)
                    else:
                        q.append((steps, neighbour))

    print_map(map, destinations)

    print()
    print(f'garden plots visited {len(destinations)}')


def print_map(map: dict[Point, str], destinations: list[Point]):
    last_y = 0
    for point in map:
        if point.y != last_y:
            last_y = point.y
            print()
        value = map[point]
        if point in destinations:
            value = '0'
        print(value, end='')
    print()


def get_neighbours(point: Point) -> Iterator[Point]:
    for offset in [(0, -1), (1, 0), (0, 1), (-1, 0)]:
        yield Point(point.x + offset[0], point.y + offset[1])


def get_input(get_test: bool) -> tuple[Point, dict[Point, str]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    rows = open(path, 'rt').read().splitlines()
    start: Point = None
    map: dict[Point, str] = {}
    for y in range(len(rows)):
        for x in range(len(rows[y])):
            map[Point(x, y)] = rows[y][x]
            if rows[y][x] == 'S':
                start = Point(x, y)
    return (start, map)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
