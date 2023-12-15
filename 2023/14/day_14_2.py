import sys
from dataclasses import dataclass
from typing import Iterator


def main(is_test_mode: bool) -> None:
    print('parabolic reflector dish')
    print()

    grid = get_input(is_test_mode)
    print_grid(grid)
    print()
    print('cycle 0')

    for cycle in range(0, 1_000_000_001):
        for y in range(0, len(grid)):
            for x in range(0, len(grid[y])):
                if grid[y][x] == 'O':
                    grid = roll_north(x, y, grid)

        for x in range(0, len(grid[y])):
            for y in range(0, len(grid)):
                if grid[y][x] == 'O':
                    grid = roll_west(x, y, grid)

        for y in range(len(grid) - 1, -1, -1):
            for x in range(0, len(grid[y])):
                if grid[y][x] == 'O':
                    grid = roll_south(x, y, grid)

        for x in range(len(grid[y]) -1 , -1, -1):
            for y in range(0, len(grid)):
                if grid[y][x] == 'O':
                    grid = roll_east(x, y, grid)

        score = score_grid(grid)
        print(f'\rcycle {cycle} {score}')

        if cycle == 999:
            break


    print()
    print()
    print_grid(grid)

    print()
    print(f'result {score_grid(grid)}')


def roll_north(x: int, y: int, grid: list[list[str]]):
    new_y = y - 1
    while new_y >= 0 and grid[new_y][x] == '.':
        new_y -= 1
    new_y += 1
    if new_y >= 0 and new_y != y:
        grid[new_y][x] = 'O'
        grid[y][x] = '.'
    return grid


def roll_west(x: int, y: int, grid: list[list[str]]):
    new_x = x - 1
    while new_x >= 0 and grid[y][new_x] == '.':
        new_x -= 1
    new_x += 1
    if new_x >= 0 and new_x != x:
        grid[y][new_x] = 'O'
        grid[y][x] = '.'
    return grid


def roll_south(x: int, y: int, grid: list[list[str]]):
    new_y = y + 1
    while new_y < len(grid) and grid[new_y][x] == '.':
        new_y += 1
    new_y -= 1
    if new_y < len(grid) and new_y != y:
        grid[new_y][x] = 'O'
        grid[y][x] = '.'
    return grid


def roll_east(x: int, y: int, grid: list[list[str]]):
    new_x = x + 1
    while new_x < len(grid[y]) and grid[y][new_x] == '.':
        new_x += 1
    new_x -= 1
    if new_x < len(grid[y]) and new_x != x:
        grid[y][new_x] = 'O'
        grid[y][x] = '.'
    return grid


def print_grid(grid: list[list[str]]):
    for y in range(0, len(grid)):
        for x in range(0, len(grid[y])):
            print(grid[y][x], end='')
        print()


def score_grid(grid: list[list[str]]) -> int:
    score = 0
    for y in range(0, len(grid)):
        for x in range(0, len(grid[y])):
            if grid[y][x] == 'O':
                score += len(grid) - y
    return score


def get_input(get_test: bool) -> list[list[str]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    grid = open(path, 'rt').read().splitlines()
    for row in range(0, len(grid)):
        grid[row] = list(grid[row])
    return grid


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
