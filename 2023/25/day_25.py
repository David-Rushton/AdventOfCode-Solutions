from __future__ import annotations
from dataclasses import dataclass
from typing import Iterator
import random
import sys


@dataclass
class Module:
    name: str
    connections: list[str]


def main(is_test_mode: bool) -> None:
    print('snowverload')
    print()

    map = get_input(is_test_mode)
    modules: dict[str, Module] = {}

    # load the modules
    for key in map:
        if key not in modules:
            modules[key] = Module(key, [])
        for connections_module in map[key]:
            if connections_module not in modules:
                modules[connections_module] = Module(connections_module, [])
            modules[connections_module].connections.append(key)
            modules[key].connections.append(connections_module)

    # look for the min cut
    cut_count = sys.maxsize
    while cut_count != 3:
        (cut_count, left_count, right_count) = kargers_min_cut(copy_modules(modules))

    print()
    print(f'min cut {cut_count} produces groups of {left_count} and {right_count} worth {left_count * right_count}')


def copy_modules(modules: dict[str, Module]) -> dict[str, Module]:
    result: dict[str, Module] = {}
    for module in modules:
        result[module] = Module(module, modules[module].connections.copy())
    return result


def kargers_min_cut(modules: dict[str, Module]) -> tuple[int, int, int]:
    """
    https://www.cs.princeton.edu/courses/archive/fall13/cos521/lecnotes/lec2final.pdf
    """
    while len(modules) != 2:
        # chose an edge at random
        left_key = random.choice(list(modules))
        right_key = random.choice(list(modules[left_key].connections))
        new_key = f'{left_key}.{right_key}'

        # remove old
        left = modules.pop(left_key)
        right = modules.pop(right_key)

        # join connections
        new_connections = []
        new_connections.extend(c for c in left.connections if c != right_key)
        new_connections.extend(c for c in right.connections if c != left_key)

        # update existing connections
        for connection in new_connections:
            if connection in modules:
                if left_key in modules[connection].connections:
                    modules[connection].connections.remove(left_key)
                    modules[connection].connections.append(new_key)
                if right_key in modules[connection].connections:
                    modules[connection].connections.remove(right_key)
                    modules[connection].connections.append(new_key)

        # add new replacement node
        modules[new_key] = Module(new_key, new_connections)

    connection_count = len(modules[next(iter(modules))].connections)
    left_count = modules[next(iter(modules))].name.count('.') + 1
    right_count = modules[next(reversed(modules))].name.count('.') + 1
    print(f'  cut of {connection_count} produces two groups of {left_count} and {right_count}')
    return (connection_count, left_count, right_count)


def get_input(get_test: bool) -> dict[str, list[str]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    result: dict[str, list[str]] = {}
    for row in open(path, 'rt').read().splitlines():
        elements = row.split(':')
        result[elements[0]] = [element for element in elements[1].split(' ') if element != '']
    return result


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
