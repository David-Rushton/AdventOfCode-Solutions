from dataclasses import dataclass
from mpl_toolkits.mplot3d import Axes3D
import matplotlib.pyplot as plt
import numpy as np
import sys

@dataclass
class Cube:
    x: int
    y: int
    z: int


def main(path: str, is_test: bool):

    print('\n== Lava Scan==')
    print('- Scanning lava')

    cubes = list(parse_input(path))
    edges = 0
    edge_cubes = []
    cubes_considered = 0
    cap = 8 if is_test else 23

    print('- Scanning grid')

    status = ''
    to_visit = [Cube(0, 0, 0)]
    visited = []

    while len(to_visit):

        current = to_visit.pop(0)

        for neighbour in get_cube_neighbours(current):

            if neighbour.x < -1 or neighbour.y < -1 or neighbour.z < -1 or neighbour.x >= cap or neighbour.y >= cap or neighbour.z >= cap:
                continue

            if neighbour in cubes:
                if neighbour not in edge_cubes:
                    edge_cubes.append(neighbour)
                edges += 1
            else:
                if neighbour not in visited:
                    if neighbour not in to_visit:
                        to_visit.append(neighbour)

        cubes_considered += 1
        status = f'  - Considered = {cubes_considered} | Cells in grid = {cap ** 3} | Edges found = {edges} | To visit = {len(to_visit)} | Visited = {len(visited)} | current = {current}     \r'
        print(status, end='')
        visited.append(current)

    print(status)

    test_assertions(cubes, edge_cubes, visited)

    print('- Results')
    print(f'  - Exposed edges {edges}')

    plot(edge_cubes, [], cap)
    exit(0)

def test_assertions(cubes: list[Cube], edge_cubes: list[Cube], visited: list[Cube]):
    # assert visited and cubes do not overlap
    print('- Asserting\n  - Visited and cubes do not overlap')

    for cube in cubes:
        if cube in visited:
            print('    - Failed')
            raise Exception(f'Found cube {cube} in visited')

    print('    - Passed')

    # assert edges in cubes
    print('  - Edges cubes always overlap')

    for edge in edge_cubes:
        if edge not in cubes:
            print('    - Failed')
            raise Exception(f'Cannot find edge {edge} in cubes')

    print('    - Passed')

def plot(left_cubes: list[Cube], right_cubes: list[Cube], cap: int):
    fig = plt.figure()
    ax = fig.add_subplot(111, projection='3d')

    axes = [cap, cap, cap]
    data = np.ones(axes, dtype=np.bool_)
    colours = np.zeros(axes, dtype=np.object_)

    for x in range(cap):
        for y in range(cap):
            for z in range(cap):
                xyz = Cube(x, y, z)

                data[x, y, z] = (xyz in left_cubes or xyz in right_cubes)

                colours[x, y, z] = [0, 0, 0, 0]

                if xyz in left_cubes:
                    colours[x, y , z] = (0, 1, 0, .5)
                else:
                    if xyz in right_cubes:
                        colours[x, y , z] = (1, 0, 0, 1.)

    ax.set_xlabel('x')
    ax.set_ylabel('y')
    ax.set_zlabel('z')

    ax.voxels(data, facecolors=colours)

    plt.show()

def get_cube_neighbours(cube: Cube):
    yield Cube(cube.x + 1, cube.y    , cube.z    )
    yield Cube(cube.x - 1, cube.y    , cube.z    )
    yield Cube(cube.x    , cube.y + 1, cube.z    )
    yield Cube(cube.x    , cube.y - 1, cube.z    )
    yield Cube(cube.x    , cube.y    , cube.z + 1)
    yield Cube(cube.x    , cube.y    , cube.z - 1)

def parse_input(path: str):

    min_v = sys.maxsize
    max_v = sys.maxsize * -1

    for cube in open(path, 'r').read().splitlines():

        elements = cube.split(',')
        min_v = min([int(elements[0]), int(elements[1]), int(elements[2]), min_v])
        max_v = max([int(elements[0]), int(elements[1]), int(elements[2]), max_v])

        yield(Cube(
            x=int(elements[0]),
            y=int(elements[1]),
            z=int(elements[2])
            )
        )

    print(f'  - min_v: {min_v}')
    print(f'  - max_v: {max_v}')


if __name__ == '__main__':
    is_test = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test else 'input.txt'
    main(path, is_test)
