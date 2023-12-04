from dataclasses import dataclass
from enum import Enum
import sys
from typing import Iterator


@dataclass
class Cubes:
    def __init__(self):
        self.red = 0
        self.green = 0
        self.blue = 0
    red: int
    green: int
    blue: int


@dataclass
class Game:
    def __init__(self, id: int):
        self.id = id
        self.cubes = []
        self.fail = False
    id: int
    cubes: list[Cubes]
    fail: bool


def main(is_test_mode: bool) -> None:
    required_red = 12
    required_green = 13
    required_blue = 14
    running_id_total = 0
    running_powerset_total = 0

    games = get_input(is_test_mode)
    for game in games:
        minimum_red = 0
        minimum_green = 0
        minimum_blue = 0
        for cube in game.cubes:
            if cube.red > required_red or cube.green > required_green or cube.blue > required_blue:
                game.fail = True
            if cube.red > minimum_red:
                minimum_red = cube.red
            if cube.green > minimum_green:
                minimum_green = cube.green
            if cube.blue > minimum_blue:
                minimum_blue = cube.blue

        if game.fail == False:
            running_id_total += int(game.id)

        running_powerset_total += minimum_red * minimum_green * minimum_blue

    print(f'The sum of ids {running_id_total}')
    print(f'The sum of power sets {running_powerset_total}')



def get_input(get_test: bool) -> Iterator[Game]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path

    for line in open(path, 'rt').read().splitlines():
        game = Game(line.split(':')[0].split(' ')[1])
        for cube_set in line.split(':')[1].split(';'):
            cubes = Cubes()
            colour = ''
            for cube_colour in cube_set.split(' ')[::-1]:
                if cube_colour.isdigit():
                    if colour == 'red':
                        cubes.red = int(cube_colour)
                    elif colour == 'green':
                        cubes.green = int(cube_colour)
                    elif colour == 'blue':
                        cubes.blue = int(cube_colour)
                    colour = ''
                else:
                    colour = cube_colour.replace(',', '')
            game.cubes.append(cubes);
        yield game


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
