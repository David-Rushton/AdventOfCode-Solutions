import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass(frozen=True)
class Location:
    x: int
    y: int


@dataclass
class Cell:
    location: Location
    pipe: str
    steps: int
    enclosed: bool
    is_enclosed: bool


def main(is_test_mode: bool) -> None:
    print('pipe maze')
    print()

    (start, cells) = get_input(is_test_mode)
    cells[start].pipe = get_start_connector(start, cells)
    cells_to_check: list[Cell] = [cells[start]]

    while len(cells_to_check) > 0:
        cell_to_check = cells_to_check.pop()
        for connection in get_connections(cell_to_check, cells):
            steps = cell_to_check.steps + 1
            if connection.location != start:
                if connection.steps == 0 or steps < connection.steps:
                    connection.steps = steps
                    cells_to_check.append(connection)

    # HACK: we use steps elsewhere to determine the main pipe from others.
    cells[start].steps = 1

    print_cells(cells)

    stretched_cells = stretch(cells)
    stretched_start = get_internal_cell(stretched_cells)
    flood_fill(stretched_start, stretched_cells)
    enclosed_cells = count_enclosed(stretched_cells)
    print_cells(stretched_cells)

    print()
    print(f'max steps {max(cells[cell].steps for cell in cells if cells[cell].steps > 0)}')
    print(f'enclosed {enclosed_cells}')


def get_start_connector(start: Location, cells: dict[Location, Cell]) -> str:
    connections: list[Location] = []
    for offsets in [(-1, 0), (0, -1), (0, 1), (1, 0)]:
        location = Location(start.x + offsets[1], start.y + offsets[0])
        if (location in cells):
            for connection in get_connections(cells[location], cells):
                if connection.location == start:
                    connections.append(Location(location.x - start.x, location.y - start.y))
    return get_connecting_pipe(connections)


def get_internal_cell(cells: dict[Location, Cell]) -> Location:
    for cell in cells:
        if cells[cell].pipe == 'F' and cells[cell].steps > 0:
            return Location(cell.x + 1, cell.y + 1)


def flood_fill(start: Location, cells: dict[Location, Cell]) -> None:
    q: list[Location] = [start]
    visited: set[Location] = set()
    while len(q) > 0:
        current = q.pop(0)
        cells[current].pipe = '#'
        if current in visited:
            continue
        for offset in [(0, -1), (1, 0), (0, 1), (-1, 0)]:
            next = Location(current.x + offset[0], current.y + offset[1])
            if next in visited:
                continue
            if cells[next].pipe == '.':
                q.append(next)
        visited.add(current)


def count_enclosed(cells: dict[Location, Cell]) -> int:
    result = 0
    for cell in cells:
        if cell.x % 2 == 0 and cell.y % 2 == 0:
            if cells[cell].pipe == '#':
                cells[cell].pipe = 'I'
                result += 1
    return result


def print_cells(cells: dict[Location, Cell]) -> None:
    last_y = 0
    for cell in sorted(cells, key=lambda k: k.y * 1_000_000 + k.x):
        if cell.y != last_y:
            last_y = cell.y
            print()
        value = cells[cell].pipe
        print(value, end='')
    print()
    print()


def stretch(cells: dict[Location, Cell]) -> dict[Location, Cell]:
    return stretch_vertical(stretch_horizontal(cells))


def stretch_vertical(cells: dict[Location, Cell]) -> dict[Location, Cell]:
    m = {
        '|': '|',
        '-': '.',
        'L': '.',
        'J': '.',
        '7': '|',
        'F': '|',
        'S': 'S',
        '.': '.'}
    result: dict[Location, cell] = {}
    for cell in cells:
        value = '.' if cells[cell].steps == 0 else cells[cell].pipe
        result[Location(cell.x, cell.y * 2)] = Cell(Location(cell.x, cell.y * 2), value, cells[cell].steps, False, False)
        result[Location(cell.x, cell.y * 2 + 1)] = Cell(Location(cell.x, cell.y * 2 + 1), m[value], cells[cell].steps, False, False)
    return result


def stretch_horizontal(cells: dict[Location, Cell]) -> dict[Location, Cell]:
    m = {
        '|': '.',
        '-': '-',
        'L': '-',
        'J': '.',
        '7': '.',
        'F': '-',
        'S': 'S',
        '.': '.'}
    result: dict[Location, cell] = {}
    for cell in cells:
        value = '.' if cells[cell].steps == 0 else cells[cell].pipe
        result[Location(cell.x * 2, cell.y)] = Cell(Location(cell.x * 2, cell.y), value, cells[cell].steps, False, False)
        result[Location(cell.x * 2 + 1, cell.y)] = Cell(Location(cell.x * 2 + 1, cell.y), m[value], cells[cell].steps, False, False)
    return result


def get_connections(start: Cell, cells: dict[Location, Cell]) -> list[Cell]:
    north = Location(0, -1)
    east = Location(1, 0)
    south = Location(0, 1)
    west = Location(-1, 0)
    offsets: list[Location] = []

    if start.pipe == '|':
        offsets.append(north)
        offsets.append(south)
    elif start.pipe == '-':
        offsets.append(east)
        offsets.append(west)
    elif start.pipe == 'L':
        offsets.append(north)
        offsets.append(east)
    elif start.pipe == 'J':
        offsets.append(north)
        offsets.append(west)
    elif start.pipe == '7':
        offsets.append(south)
        offsets.append(west)
    elif start.pipe == 'F':
        offsets.append(south)
        offsets.append(east)

    for offset in offsets:
        candidate = Location(start.location.x + offset.x, start.location.y + offset.y)
        if candidate in cells:
            yield cells[candidate]


def get_connecting_pipe(offsets: list[Location]) -> str:
    north = Location(0, -1)
    east = Location(1, 0)
    south = Location(0, 1)
    west = Location(-1, 0)

    if north in offsets and south in offsets:
        return '|'
    if east in offsets and west in offsets:
        return '-'
    if north in offsets and east in offsets:
        return 'L'
    if north in offsets and west in offsets:
        return 'J'
    if south in offsets and west in offsets:
        return '7'
    if south in offsets and east in offsets:
        return 'F'


def get_input(get_test: bool) -> tuple[Location, dict[Location, Cell]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    rows = open(path, 'rt').read().splitlines()
    start: Location = None
    cells: dict[Location, Cell] = {}
    for y in range(0, len(rows)):
        for x in range(0, len(rows[y])):
            location = Location(x, y)
            cells[location] = Cell(location, rows[y][x], 0, False, False)
            if rows[y][x] == 'S':
                start = location
    return (start, cells)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
