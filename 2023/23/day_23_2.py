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

    print(f' locating junctions')
    junctions: list[Point] = [start, end]
    for point in map:
        if map[point] != '#':
            neighbours = list(get_neighbours(point, map))
            if len(neighbours) > 2:
                junctions.append(point)

    print(f' plotting segments')
    segments: dict[Point, list[tuple[Point, int]]] = {}
    for junction in junctions:
        for (junction_start, junction_end, junction_steps) in find_route(junction, junctions, map):
            if junction_start not in segments:
                segments[junction_start] = []
            segments[junction_start].append((junction_end, junction_steps))

    print(f' calculating best case')
    best_case = 0
    for segment in segments:
        segment_best_case = 0
        for (_, segment_steps)  in segments[segment]:
            segment_best_case = max(segment_best_case, segment_steps)
        best_case += segment_best_case


    print(f' finding longest route')
    q: list[tuple[Point, set[Point], int, int]] = [(start, set([start]), 0, best_case)]
    max_steps = 0
    while len(q) > 0:
        q.sort(key=lambda k: k[2], reverse=True)
        (current, route, steps, route_best_case) = q.pop(0)
        segments_to_check = [segment for segment in segments[current] if segment[0] not in route]
        segments_steps = sum(segment[1] for segment in segments_to_check)
        for (segment_end, segment_steps) in segments_to_check:
            new_route = route.copy()
            new_route.add(segment_end)
            new_steps = steps + segment_steps
            new_best_case = route_best_case - segments_steps + segment_steps
            if new_best_case <= max_steps:
                continue
            if segment_end == end:
                if new_steps > max_steps:
                    print(f'  new best route found in {new_steps}')
                    # print(' > '.join(f'{step.x}x{step.y}'  for step in route))
                    max_steps = new_steps
            else:
                q.append((segment_end, new_route, new_steps, route_best_case))

    print()
    print(f'longest route found {max_steps} steps')


def find_route(start: Point, destinations: list[Point], map: dict[Point, str]) -> Iterator[tuple[Point, Point, int]]:
    q: list[tuple[Point, set[Point], int]] = [(start, set([start]), 0)]
    while (len(q) > 0):
        (current, route, steps) = q.pop(0)
        steps += 1
        for neighbour in get_neighbours(current, map):
            if neighbour in route:
                continue
            new_route = route.copy()
            new_route.add(neighbour)
            if neighbour in destinations:
                end = neighbour
                yield (start, end, steps)
            else:
                q.append((neighbour, new_route, steps))


def get_neighbours(start: Point, map: dict[Point, str]) -> Iterator[Point]:
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
