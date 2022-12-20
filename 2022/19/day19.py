from dataclasses import dataclass
from enum import Enum
import os
import sys
import re

@dataclass
class Resources:
    ore: int
    clay: int
    obsidian: int
    geode: int

@dataclass
class Robot:
    ore_cost: int
    clay_cost: int
    obsidian_cost: int

@dataclass
class Blueprint:
    id_number: int
    quality_level: int
    ore_robot: Robot
    clay_robot: Robot
    obsidian_robot: Robot
    geode_robot: Robot


def main(is_test: bool, path: str) -> None:
    blueprints = list(parse_blueprints(path))

    for blueprint in blueprints:
        blueprint.quality_level = calculate_quality_level(blueprint)
        print_blueprint(blueprint)

def calculate_quality_level(blueprint: Blueprint) -> int:
    # prioritise higher level robots.
    # figure out the quickest way to add one to stock of highest unlocked
    #   - saving up
    #   - +1 next level down
    #   - maybe recursse?
    minutes = 1

    robots = Resources(ore = 1, clay = 0, obsidian = 0, geode = 0)
    balance = Resources(ore = 0, clay = 0, obsidian = 0, geode = 0)

    while minutes <= 24:
        is_factory_building = False


        # time_to_next_ore_robot
        # time_to_next_clay_robot
        # time_to_next_obsidian_robot
        # time_to_next_geode_robot
        #
        # can_build_obsidian_without_delaying_geode
        #   ore req = ore cost - ore held
        #   days to ore = ore req / ore robots
        #   clay req = ...
        #   days to clay = ...
        #   days = max( days to ore, days to clay)
        #   rebate = ...
        #
        # can_build_clay_without_delaying_obsidian
        # can_build_ore_without_delaying_clay



        # build?
        if can_purchase(is_factory_building, balance, blueprint.geode_robot):
            balance = purchase_robot(blueprint.geode_robot, balance)
            is_factory_building = True
            balance.geode -= 1
            robots.geode += 1

        if can_purchase(is_factory_building, balance, blueprint.obsidian_robot):
            balance = purchase_robot(blueprint.obsidian_robot, balance)
            is_factory_building = True
            balance.obsidian -= 1
            robots.obsidian += 1

        if balance.clay <= blueprint.obsidian_robot.clay_cost - robots.clay:
            if can_purchase(is_factory_building, balance, blueprint.clay_robot):
                balance = purchase_robot(blueprint.clay_robot, balance)
                is_factory_building = True
                balance.clay -= 1
                robots.clay += 1

            if can_purchase(is_factory_building, balance, blueprint.ore_robot):
                balance = purchase_robot(blueprint.ore_robot, balance)
                is_factory_building = True
                balance.ore -= 1
                robots.ore += 1

        # mine
        balance.ore += robots.ore
        balance.clay += robots.clay
        balance.obsidian += robots.obsidian
        balance.geode += robots.geode

        print(f'== Minute {minutes} ==')
        print(f'- ore = {balance.ore}')
        print(f'- clay = {balance.clay}')
        print(f'- obsidian = {balance.obsidian}')
        print(f'- geode = {balance.geode}')
        print(f'- ore-collectors {robots.ore} | clay-collectors {robots.clay} | obsidian-collectors {robots.obsidian} | geode-crackers {robots.geode}')
        print()

        minutes += 1

    return blueprint.id_number * balance.geode

def purchase_robot(robot: Robot, balance: Resources) -> Resources:
    return Resources(
        ore = balance.ore - robot.ore_cost,
        clay = balance.clay - robot.clay_cost,
        obsidian = balance.obsidian - robot.obsidian_cost,
        geode = balance.geode
    )

def can_purchase(is_factory_building: bool, balance: Resources, robot: Robot) -> bool:
    if not is_factory_building:
        if balance.ore >= robot.ore_cost:
            if balance.clay >= robot.clay_cost:
                if balance.obsidian >= robot.obsidian_cost:
                    return True

    return False

def should_buy_robot(ore_required: int, clay_required: int, obsidian_required: int) -> bool:
    return False

def parse_blueprints(path: str):
    for blueprint in open(path, 'r').read().splitlines():
        elements = list(map(int, re.findall('[0-9]+', blueprint)))

        if len(elements) != 7:
            message = f'Expected 7 numbers, found {len(elements)}.  Cannot parse blueprint: {blueprint}'
            raise Exception(message)

        yield Blueprint(
            id_number=int(elements[0]),
            quality_level=0,
            ore_robot = Robot(ore_cost = elements[1], clay_cost = 0, obsidian_cost = 0),
            clay_robot = Robot(ore_cost = elements[2], clay_cost = 0, obsidian_cost = 0),
            obsidian_robot = Robot(ore_cost = elements[3], clay_cost = elements[4], obsidian_cost = 0),
            geode_robot = Robot(ore_cost = elements[5], clay_cost = 0, obsidian_cost = elements[6])
        )

def print_blueprint(blueprint: Blueprint):
    print(f'Blueprint {blueprint.id_number}:')
    print(f'  Each ore robot costs {blueprint.ore_robot.ore_cost}.')
    print(f'  Each clay robot costs {blueprint.clay_robot.ore_cost} ore.')
    print(f'  Each obsidian robot costs {blueprint.obsidian_robot.ore_cost} ore and {blueprint.obsidian_robot.clay_cost} clay.')
    print(f'  Each geode robot costs {blueprint.geode_robot.ore_cost} ore and {blueprint.geode_robot.obsidian_cost} obsidian.')
    print(f'  Quality level = {blueprint.quality_level}.')
    print()

if __name__ == '__main__':
    is_test = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test else 'input.txt'
    main(is_test, path)
