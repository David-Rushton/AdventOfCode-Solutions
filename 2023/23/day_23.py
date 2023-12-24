import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass(frozen=True)
class Point:
    x: int
    y: int


def main(is_test_mode: bool) -> None:
    print('a long walk')
    print()

    (start, end, map) = get_input(is_test_mode)

    max_steps = 0
    q: list[tuple[Point, int, set[Point]]] = [(start, 0, set([start]))]
    while (len(q) > 0):
        (current, steps, route) = q.pop(0)
        steps += 1
        for neighbour in get_neighbours(current, map):
            if neighbour in route:
                continue
            new_route = route.copy()
            new_route.add(neighbour)
            if neighbour == end:
                if steps > max_steps:
                    max_steps = steps
                print_map(map, new_route)
                print()
                print(f' route found {steps} steps')
                print()
            else:
                q.append((neighbour, steps, new_route))

    print()
    print(f'longest hike {max_steps}')


def get_neighbours(start: Point, map: dict[Point, str]) -> Iterator[Point]:
    if map[start] == '^':
        yield Point(start.x, start.y - 1)
    elif map[start] == '>':
        yield Point(start.x + 1, start.y)
    elif map[start] == 'v':
        yield Point(start.x, start.y + 1)
    elif map[start] == '<':
        yield Point(start.x - 1, start.y)
    else:
        for offset in [(0, -1), (1, 0), (0, 1), (-1 , 0)]:
            next = Point(start.x + offset[0], start.y + offset[1])
            if next in map and map[next] in ['.', '^', '>', 'v', '<']:
                yield next


def print_map(map: dict[Point, str], route: set[Point]) -> None:
    last_y = 0
    for point in map:
        if point.y != last_y:
            last_y = point.y
            print()
        if point in route:
            print('x', end='')
        else:
            print(map[point], end='')
    print()


def get_input(get_test: bool) -> tuple[Point, Point, dict[Point, str]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    rows = open(path, 'rt').read().splitlines()
    start: Point = None
    end: Point = None
    map: dict[Point, str] = {}
    for y in range(len(rows)):
        for x in range(len(rows[y])):
            point = Point(x, y)
            map[point] = rows[y][x]
            if rows[y][x] == '.':
                if start == None:
                    start = point
                end = point
    return (start, end, map)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
