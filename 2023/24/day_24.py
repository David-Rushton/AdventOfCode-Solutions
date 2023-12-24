import sys
import math
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Point:
    x: int
    y: int
    z: int


@dataclass
class Hail:
    id: int
    position: Point
    velocity: Point


def main(is_test_mode: bool) -> None:
    print('never tell me the odds')
    print()

    start = 7 if is_test_mode else 200000000000000
    end = 27 if is_test_mode else 400000000000000

    print(f'comparing hail stone trajectories')
    hail_stones = list(get_input(is_test_mode))
    checked: set(tuple(int, int)) = set()
    intersections_count = 0
    for hail_1 in hail_stones:
        for hail_2 in hail_stones:
            if hail_1.id == hail_2.id:
                continue
            key = (min(hail_1.id, hail_2.id), max(hail_1.id, hail_2.id))
            if key in checked:
                continue
            checked.add(key)
            (has_intersection, intersection) = get_intersection(get_line(hail_1), get_line(hail_2))
            if has_intersection:
                if within_area(intersection, start, end):
                    if is_future_intersection(hail_1, intersection) and is_future_intersection(hail_2, intersection):
                        intersections_count += 1
                        print(f' {hail_1.position.x}x{hail_1.position.y} | {hail_2.position.x}x{hail_2.position.y} | intersects @ {intersection.x}x{intersection.y}')
                    else:
                        print(f' {hail_1.position.x}x{hail_1.position.y} | {hail_2.position.x}x{hail_2.position.y} | intersects @ {intersection.x}x{intersection.y} | in the past')
                else:
                    print(f' {hail_1.position.x}x{hail_1.position.y} | {hail_2.position.x}x{hail_2.position.y} | intersects @ {intersection.x}x{intersection.y} | not in range')
            else:
                print(f' {hail_1.position.x}x{hail_1.position.y} | {hail_2.position.x}x{hail_2.position.y} | does not intersect')

    print()
    print(f'intersections count {intersections_count}')


def is_future_intersection(hail: Hail, intersecton: Point) -> bool:
    if hail.velocity.x > 0:
        return intersecton.x > hail.position.x
    else:
        return intersecton.x < hail.position.x


def within_area(point: Point, start: int, end: int):
    if point.x < start or point.y < start:
        return False
    if point.x > end or point.y > end:
        return False
    return True


def get_line(hail: Hail) -> tuple[Point, Point]:
   return (
       hail.position,
       Point(hail.position.x + hail.velocity.x, hail.position.y + hail.velocity.y, 0))


def get_intersection(line_1: tuple[Point, Point], line_2: tuple[Point, Point]) -> tuple[bool, Point]:
    x1 = line_1[0].x
    x2 = line_1[1].x
    x3 = line_2[0].x
    x4 = line_2[1].x
    y1 = line_1[0].y
    y2 = line_1[1].y
    y3 = line_2[0].y
    y4 = line_2[1].y
    x12 = x1 - x2
    x34 = x3 - x4
    y12 = y1 - y2
    y34 = y3 - y4
    c = x12 * y34 - y12 * x34

    # no intersection
    if math.fabs(c) < 0.01:
        return (False, None)

    a = x1 * y2 - y1 * x2
    b = x3 * y4 - y3 * x4
    x = round((a * x34 - b * x12) / c, 3)
    y = round((a * y34 - b * y12) / c, 3)
    z = round(0, 3)
    return (True, Point(x, y, z))


def get_input(get_test: bool) -> Iterator[Hail]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    next_id = 1
    for hail in open(path, 'rt').read().splitlines():
        sections = hail.replace(' ','').split('@')
        positions = [int(value) for value in sections[0].split(',')]
        velocities = [int(value) for value in sections[1].split(',')]
        yield Hail(
            next_id,
            Point(positions[0], positions[1], positions[2]),
            Point(velocities[0], velocities[1], velocities[2]))
        next_id += 1


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
