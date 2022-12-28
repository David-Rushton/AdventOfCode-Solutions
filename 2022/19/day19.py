from dataclasses import dataclass
import re
import sys
from typing import Generator


TIME_LIMIT = 24


@dataclass
class RobotCost:
    ore_required: int
    clay_required: int
    obsidian_required: int

@dataclass
class Blueprint:
    id: int
    ore_robot: RobotCost
    clay_robot: RobotCost
    obsidian_robot: RobotCost
    geode_robot: RobotCost
    maximum_required_per_minute: RobotCost

@dataclass
class BlueprintQualityLevel:
    id: int
    geodes_collected: int
    quality_level: int

@dataclass
class Robots:
    ore_robots: int
    clay_robots: int
    obsidian_robots: int
    geode_robots: int

@dataclass
class Resources:
    time_remaining: int
    ore_collected: int
    clay_collected: int
    obsidian_collected: int
    geode_collected: int


def main(is_test_mode: bool, path: str) -> None:
    print('\n== Not Enough Minerals ==\n')
    for quality_level in get_quality_levels(parse_blueprints(path)):
        print(f'- Blueprint #{quality_level.id} collected {quality_level.geodes_collected} geodes.  Quality level = {quality_level.quality_level}.')
    print()

def get_quality_levels(blueprints: Generator[Blueprint, None, None]) -> Generator[BlueprintQualityLevel, None, None]:
    for blueprint in blueprints:
        resources = Resources(time_remaining=TIME_LIMIT, ore_collected=0, clay_collected=0, obsidian_collected=0, geode_collected=0)
        robots = Robots(ore_robots=1, clay_robots=0, obsidian_robots=0, geode_robots=0)
        yield get_blueprint_quality_level(blueprint, resources, robots)

def get_blueprint_quality_level(blueprint: Blueprint, resources: Resources, robots: Robots) -> BlueprintQualityLevel:
    while resources.time_remaining > 0:

        # discard less productive routes early
        # figure out what robots we can build
            # include unlocked robots
                # where unlocked = resource collected per minute < max required per minute
                # &&
                # we are mining required build resources
        # recursively try each unlocked robot
        # once geode cracker is created we can calculate its total lifetime yield
            # we can generate a heuristic best case here for future yield, to filter less productive routes early
                # heuristic = sum of numbers 1 to time remaining + geodes already collected

        # TODO:
        # implement above

        pass

    return BlueprintQualityLevel(id=blueprint.id, geodes_collected=0, quality_level=0)

def get_geodes_available(geode_robots: int, time_remaining: int) -> int:
    """
    Best case; we can build one additional geode cracking robot every minute.
    Each robot cracks one geode per minute, from creation until the time limit.
    If there are 5 minutes left we can build 5 robots:
        min    robots  collected
        1      1        0
        2      2        1
        3      3        3
        4      4        10
        5      5        15
    Where collected is equal to the sum of robots available in each minute.
    """
    # step squad, hush, hush...
    max_yield_from_new_robots = time_remaining * (time_remaining + 1) / 2
    return max_yield_from_new_robots + (geode_robots * time_remaining)

def print_blueprint(blueprint: Blueprint) -> None:
    print(f'\nBlueprint: {blueprint.id}')
    print(f'- Each ore robot costs {blueprint.ore_robot.ore_required} ore')
    print(f'- Each clay robot costs {blueprint.clay_robot.ore_required} ore')
    print(f'- Each obsidian robot costs {blueprint.obsidian_robot.ore_required} ore and {blueprint.obsidian_robot.clay_required}')
    print(f'- Each geode robot costs {blueprint.geode_robot.ore_required} ore and {blueprint.geode_robot.obsidian_required}\n')

def parse_blueprints(path: str) -> Generator[Blueprint, None, None]:
    # Format:
    # Blueprint x: Each ore robot costs x ore. Each clay robot costs x ore. Each obsidian robot costs x ore and x clay. Each geode robot costs x ore and x obsidian.
    #           0                       1                            2                                3         4                              5         6
    blueprints = open(path, 'r').read().splitlines()
    for blueprint in blueprints:
        numbers = re.findall('[0-9]+', blueprint)
        max_ore = max(int(numbers[1]), int(numbers[2]), int(numbers[3]), int(numbers[5]))
        max_clay = int(numbers[4])
        max_obsidian = int(numbers[6])

        yield Blueprint(
            id=int(numbers[0]),
            ore_robot=RobotCost(
                ore_required=int(numbers[1]),
                clay_required=0,
                obsidian_required=0
            ),
            clay_robot=RobotCost(
                ore_required=int(numbers[2]),
                clay_required=0,
                obsidian_required=0
            ),
            obsidian_robot=RobotCost(
                ore_required=int(numbers[3]),
                clay_required=int(numbers[4]),
                obsidian_required=0
            ),
            geode_robot=RobotCost(
                ore_required=int(numbers[5]),
                clay_required=0,
                obsidian_required=int(numbers[6])
            ),
            maximum_required_per_minute=RobotCost(
                ore_required=max_ore,
                clay_required=max_clay,
                obsidian_required=max_obsidian
            )
        )


if __name__ == '__main__':
    is_test_mode = True if sys.argv[1] == 'test' else False
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, path)
