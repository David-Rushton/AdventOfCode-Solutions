import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass(frozen=True)
class Point:
    x: int
    y: int


@dataclass
class Cell:
    heat_loss: int


def main(is_test_mode: bool) -> None:
    print('clumsy crucible')
    print()

    map = get_input(is_test_mode)
    last_y = 0
    for point in map:
        if point.y > last_y:
            print()
            last_y = point.y
        print(map[point].heat_loss, end='')
    print()

    max_x = max(point.x for point in map)
    max_y = max(point.y for point in map)
    find_best_route_3(map, Point(max_x, max_y))
    print()

    print()
    print()

best = 800
def find_best_route_3(map: dict[Point, Cell], destination: Point):
    global best
    start = Point(0, 0)
    q: list[tuple[Point, str, int]] = []
    q.append((start, '>', 0))
    q.append((start, 'v', 0))
    cache = get_cache(map)
    x_offsets = {'^': 0,  '>': 1, 'v': 0, '<': -1}
    y_offsets = {'^': -1, '>': 0, 'v': 1, '<': 0 }
    turns = {'^': ['>', '<'], '>': ['v', '^'], 'v': ['>', '<'], '<': ['v', '^']}

    counter = 0

    print()

    while len(q) > 0:
        (point, direction, heat_loss) = q.pop()
        counter +=1
        print(f'\r {counter}: {len(q)}    {best}                 \t\t\t', end='')
        if heat_loss > cache[point][direction]:
            continue
        if heat_loss > best:
            continue

        for _ in range(3):
            point = Point(point.x + x_offsets[direction], point.y + y_offsets[direction])
            if point not in map:
                break
            heat_loss += map[point].heat_loss
            for turn in turns[direction]:
                if heat_loss < cache[point][turn]:
                    cache[point][turn] = heat_loss
                    q.append((point, turn, heat_loss))

                    if point == destination:
                        if heat_loss < best:
                            best = heat_loss

    print()
    for direction in cache[destination]:
        print(f' {destination} {direction} {cache[destination][direction]}')



def get_cache(map: dict[Point, Cell]) -> dict[Point,dict[str,int]]:
    result: dict[Point,dict[str,int]] = {}
    for point in map:
        result[point] = {}
        for direction in ['^', '>', 'v', '<']:
            result[point][direction] = sys.maxsize
    return result


def get_input(get_test: bool) -> dict[Point, Cell]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    result: dict[Point, Cell] = {}
    rows = open(path, 'rt').read().splitlines()
    for y in range(0, len(rows)):
        for x in range(0, len(rows[y])):
            result[Point(x, y)] = Cell(int(rows[y][x]))
    return result


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
