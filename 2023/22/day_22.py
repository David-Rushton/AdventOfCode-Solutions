import sys
import matplotlib as mpl
import matplotlib.pyplot as plt
import numpy as np
from mpl_toolkits.mplot3d import Axes3D
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Point:
    x: int
    y: int
    z: int


@dataclass
class Brick:
    id: int
    start: Point
    end: Point
    points: list[Point]


def main(is_test_mode: bool) -> None:
    print('sand slabs')
    print()

    (max_x, max_y, max_z, bricks) = get_input(is_test_mode)

    print(' dropping bricks')
    should_repeat = True
    iterations = 0
    while should_repeat:
        should_repeat = False
        iterations += 1
        print(f'  iteration {iterations}')
        bricks.sort(key=lambda b: b.start.z)
        for brick in bricks:
            if brick.start.z <= 1:
                continue
            overlaps_found = 0
            for other in bricks:
                if other.id == brick.id:
                    continue
                if brick.start.z - 1 >= other.start.z and brick.start.z - 1 <= other.end.z:
                    if is_overlapping(brick, other):
                        overlaps_found += 1
            if overlaps_found == 0:
                should_repeat = True
                brick.start.z -= 1
                brick.end.z -= 1
                brick.points = list(fill_brick(brick))

    print(' building index')
    index: dict[tuple(int, int, int), int] = {}
    for brick in bricks:
        for point in brick.points:
            index[(point.x, point.y, point.z)] = brick.id

    print(' searching for disintegration candidates')
    required: set[int] = set([])
    for brick in bricks:
        supported_by: set[int] = set([])
        for point in brick.points:
            key = (point.x, point.y, point.z -1)
            if key in index:
                if index[key] != brick.id:
                    supported_by.add(index[key])
        if len(supported_by) == 1:
            required.add(supported_by.pop())
    candidates: set[int] = set([])
    for brick in bricks:
        if brick.id not in required:
            candidates.add(brick.id)

    # print(' plotting bricks')
    # plot_bricks(max_x, max_y, max_z, bricks)

    print()
    print(f'disintegration candidates found {len(candidates)}')


def is_overlapping(brick: Brick, other: Brick) -> bool:
    # entirely under
    if brick.start.y > other.end.y:
        return False
    # entirely over
    if brick.end.y < other.start.y:
        return False
    # entirely left
    if brick.end.x  < other.start.x:
        return False
    # entirely right
    if brick.start.x > other.end.x:
        return False
    # overlaps
    return True


def plot_bricks(max_x: int, max_y: int, max_z: int, bricks: list[Brick]) -> None:
    fig = plt.figure()
    ax = fig.add_subplot(111, projection='3d')
    data = np.zeros([max_x, max_y, max_z], dtype=np.bool_)
    for brick in bricks:
        for point in brick.points:
            data[point.x, point.y, point.z] = True
    ax.set_xlabel('x')
    ax.set_ylabel('y')
    ax.set_zlabel('z')
    ax.voxels(data, edgecolors='gray')
    plt.show()


def get_input(get_test: bool) -> tuple[int, int, int, list[Brick]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    max_x = 0
    max_y = 0
    max_z = 0
    bricks: list[Brick] = []
    next_id = 1
    for row in open(path, 'rt').read().splitlines():
        elements = row.split('~')
        start = [int(coordinate) for coordinate in elements[0].split(',')]
        end = [int(coordinate) for coordinate in elements[1].split(',')]
        brick = Brick(next_id, Point(start[0], start[1], start[2]), Point(end[0], end[1], end[2]), [])
        next_id += 1
        brick.points = list(fill_brick(brick))
        bricks.append(brick)
        max_x = max(max_x, brick.start.x, brick.end.x)
        max_y = max(max_y, brick.start.y, brick.end.y)
        max_z = max(max_z, brick.start.z, brick.end.z)
    return (max_x + 1, max_y + 1, max_z + 1, bricks)


def fill_brick(brick: Brick) -> Iterator[Point]:
    for x in range(min(brick.start.x, brick.end.x), max(brick.start.x, brick.end.x) + 1):
        for y in range(min(brick.start.y, brick.end.y), max(brick.start.y, brick.end.y) + 1):
            for z in range(min(brick.start.z, brick.end.z), max(brick.start.z, brick.end.z) + 1):
                yield(Point(x, y, z))


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
