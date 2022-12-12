from dataclasses import dataclass
import sys
import os
import time
from typing import Tuple


@dataclass
class Cell:
    row: int
    column: int
    index: Tuple[int, int]
    distance: int
    label: str
    value: int
    visited: bool

def main(path: str):
    os.system('cls')

    grid = open(path, 'r').read().splitlines()
    iteration = 0
    unvisited = []
    visited = []
    index = {}
    priority_queue = []
    iteration = 0

    for row in range(len(grid)):
        for column in range(len(grid[row])):
            cell = Cell(
                row,
                column,
                index=(row, column),
                distance = 0 if grid[row][column] == 'E' else sys.maxsize,
                label = grid[row][column],
                value = ord(grid[row][column]),
                visited = False
            )
            unvisited.append(cell)
            index[(row, column)] = cell

    while len(unvisited) > 0:
        unvisited.sort(key=lambda x: x.distance)
        current = unvisited[0]

        for neighbour in get_neighbours(grid, (current.row, current.column)):
            if neighbour not in visited:
                if can_move(get_cell(grid, current.index), get_cell(grid, neighbour)):
                    neighbour_cell = index[neighbour]
                    neighbour_cell.distance = get_distance(current, index[neighbour])

        current.visited = True
        visited.append(unvisited.pop(0).index)
        iteration += 1
        print_iteration(grid, index, iteration)

        if current.label in ['S', 'a']:
            print(f'\nDestination reached in {current.distance} steps')
            break

def get_distance(visited: Cell, unvisited: Cell):
    distance = visited.distance + 1;
    return distance if distance < unvisited.distance else unvisited.distance

def print_iteration(grid, index, iteration):
    print('\033[1;1H')
    print(f'iteration: {iteration}\n\n')
    for row in range(len(grid)):
        for column in range(len(grid[row])):
            cell = index[(row, column)]
            value = '#' if cell.visited == True else '.'
            if cell.label == 'S':
                value = 'S'
            if cell.label == 'E':
                value = 'E'
            print(value, end='')
        print()

def get_cell(grid: list[str], location):
    return grid[location[0]][location[1]]

def get_neighbours(grid: list[str], location):
    row = location[0]
    column = location[1]

    row_offset = 1
    new_row = row + row_offset
    if new_row >= 0 and new_row < len(grid):
        yield (new_row, column)

    column_offset = 1
    new_column = column + column_offset
    if new_column >= 0 and new_column < len(grid[row]):
        yield (row, new_column)

    column_offset = -1
    new_column = column + column_offset
    if new_column >= 0 and new_column < len(grid[row]):
        yield (row, new_column)

    row_offset = -1
    new_row = row + row_offset
    if new_row >= 0 and new_row < len(grid):
        yield (new_row, column)


def find_start(grid: list[str]):
    return find_cell(grid, 'S')

def find_end(grid: list[str]):
    return find_cell(grid, 'E')

def find_cell(grid: list[str], find: str):
    for row in range(len(grid)):
        for column in range(len(grid[row])):
            if grid[row][column] == find:
                return (row, column)

    raise Exception(f'Cannot find {find}')

def can_move(from_value, destination_value):

    if from_value == 'E':
        from_value = 'z'

    if destination_value == 'S':
        destination_value = 'a'

    return ord(from_value) - ord(destination_value) <= 1


if __name__ == '__main__':
    path = 'day12.input.txt.test' if sys.argv[1] == 'test' else 'day12.input.txt'
    main(path)
