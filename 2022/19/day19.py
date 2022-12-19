from dataclasses import dataclass
import os
import sys
import re

@dataclass
class Robot:
    ore: int
    clay: int
    obsidian: int

@dataclass
class Blueprint:
    id_number: int
    ore_robot: Robot
    clay_robot: Robot
    obsidian_robot: Robot
    geode_robot: Robot


def main(is_test: bool, path: str) -> None:
    blueprints = list(parse_blueprints(path))

    for blueprint in blueprints:
        print_blueprint(blueprint)

def parse_blueprints(path: str):
    for blueprint in open(path, 'r').read().splitlines():
        elements = re.findall('[0-9]+', blueprint)

        if len(elements) != 7:
            message = f'Expected 7 numbers, found {len(elements)}.  Cannot parse blueprint: {blueprint}'
            raise Exception(message)

        yield Blueprint(
            id_number=int(elements[0]),
            ore_robot=Robot(int(elements[1]), 0, 0),
            clay_robot=Robot(int(elements[2]), 0, 0),
            obsidian_robot=Robot(int(elements[3]), int(elements[4]), 0),
            geode_robot=Robot(int(elements[5]), 0, int(elements[6]))
        )

def print_blueprint(blueprint: Blueprint):
    print(f'Blueprint  {blueprint.id_number}:')
    print(f'  Each ore robot costs {blueprint.ore_robot.ore}.')
    print(f'  Each clay robot costs {blueprint.clay_robot.ore} ore.')
    print(f'  Each obsidian robot costs {blueprint.obsidian_robot.ore} ore and {blueprint.obsidian_robot.clay} clay.')
    print(f'  Each geode robot costs {blueprint.geode_robot.ore} ore and {blueprint.geode_robot.obsidian} obsidian.')
    print()

if __name__ == '__main__':
    is_test = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test else 'input.txt'
    main(is_test, path)
