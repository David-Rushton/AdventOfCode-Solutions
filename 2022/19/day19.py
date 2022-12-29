import dataclasses
from dataclasses import dataclass
from enum import Enum
import math
import re
import sys
from typing import Generator


STAR_ONE_TIME_LIMIT = 24
STAR_TWO_TIME_LIMIT = 32

class RobotType(Enum):
    ORE_ROBOT = 0
    CLAY_ROBOT = 1
    OBSIDIAN_ROBOT = 2
    GEODE_ROBOT = 3

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

@dataclass
class BlueprintQualityLevel:
    id: int
    robots: Robots
    geodes_collected: int
    quality_level: int

@dataclass
class BestBlueprintConfiguration:
    resources: Resources
    robots: Robots
    geode_collected: int

def main(is_test_mode: bool, is_star_two: bool, path: str) -> None:
    print('\n== Not Enough Minerals ==')

    quality_levels_sum = 0
    first_three_geodes_cracked_product = 1
    blueprints_checked = 0
    time_limit = STAR_TWO_TIME_LIMIT if is_star_two else STAR_ONE_TIME_LIMIT

    for quality_level in get_quality_levels(parse_blueprints(path), time_limit):
        robots = quality_level.robots
        quality_levels_sum += quality_level.quality_level
        blueprints_checked += 1
        first_three_geodes_cracked_product = first_three_geodes_cracked_product * quality_level.geodes_collected
        print(f'- Geodes collected: {quality_level.geodes_collected}')
        print(f'- Robots: ore {robots.ore_robots}, clay {robots.clay_robots}, obsidian {robots.obsidian_robots}, geode {robots.geode_robots}')
        print(f'- Quality level: {quality_level.quality_level}.')

        if blueprints_checked == 3:
            break

    print(f'\nQuality level sum: {quality_levels_sum}')
    print(f'Product of geodes cracked in first three blueprints: {first_three_geodes_cracked_product}\n')

def get_quality_levels(blueprints: Generator[Blueprint, None, None], time_limit: int) -> Generator[BlueprintQualityLevel, None, None]:
    for blueprint in blueprints:
        print_blueprint(blueprint)
        resources = Resources(time_remaining=time_limit, ore_collected=0, clay_collected=0, obsidian_collected=0, geode_collected=0)
        robots = Robots(ore_robots=1, clay_robots=0, obsidian_robots=0, geode_robots=0)
        best_config = BestBlueprintConfiguration(resources, robots, geode_collected=0)
        yield get_blueprint_quality_level(blueprint, resources, robots, best_config)

def get_blueprint_quality_level(
    blueprint: Blueprint,
    resources: Resources,
    robots: Robots,
    best_blueprint_config: BestBlueprintConfiguration,
    depth: int=0
) -> BlueprintQualityLevel:

    best_possible_return = get_geodes_available(resources.geode_collected, resources.time_remaining)
    if best_possible_return < best_blueprint_config.geode_collected:
        return

    for robot_type in get_available_robots_types(blueprint.maximum_required_per_minute, robots):
        time_required = get_time_to_build_robot(robot_type, blueprint, resources, robots)

        if time_required >= resources.time_remaining:
            continue

        updated_resources = mine_resources(time_required, resources, robots)

        updated_resources, updated_robots = build_robot(robot_type, blueprint, updated_resources, robots)

        if updated_resources.geode_collected > best_blueprint_config.geode_collected:
            best_blueprint_config.geode_collected = updated_resources.geode_collected
            best_blueprint_config.resources = updated_resources
            best_blueprint_config.robots = updated_robots

        get_blueprint_quality_level(blueprint, updated_resources, updated_robots, best_blueprint_config, depth + 1)

    if depth == 0:
        return BlueprintQualityLevel(
            id=blueprint.id,
            robots=best_blueprint_config.robots,
            geodes_collected=best_blueprint_config.geode_collected,
            quality_level=blueprint.id * best_blueprint_config.geode_collected
        )

def mine_resources(time_passed: int, resources: Resources, robots: Robots) -> Resources:
    return Resources(
        time_remaining=resources.time_remaining - time_passed,
        ore_collected=resources.ore_collected + (time_passed * robots.ore_robots),
        clay_collected=resources.clay_collected + (time_passed * robots.clay_robots),
        obsidian_collected=resources.obsidian_collected + (time_passed * robots.obsidian_robots),
        geode_collected=resources.geode_collected
    )

def build_robot(
    robot_type: RobotType, blueprint: Blueprint, resources: Resources, robots: Robots
) -> tuple[Resources, Robots]:
    updated_robots = dataclasses.replace(robots)
    updated_resources = dataclasses.replace(resources)

    if robot_type == RobotType.GEODE_ROBOT:
        updated_resources.ore_collected -= blueprint.geode_robot.ore_required
        updated_resources.obsidian_collected -= blueprint.geode_robot.obsidian_required
        updated_robots.geode_robots += 1
        updated_resources.geode_collected += updated_resources.time_remaining

    if robot_type == RobotType.OBSIDIAN_ROBOT:
        updated_resources.ore_collected -= blueprint.obsidian_robot.ore_required
        updated_resources.clay_collected -= blueprint.obsidian_robot.clay_required
        updated_robots.obsidian_robots += 1

    if robot_type == RobotType.CLAY_ROBOT:
        updated_resources.ore_collected -= blueprint.clay_robot.ore_required
        updated_robots.clay_robots += 1

    if robot_type == RobotType.ORE_ROBOT:
        updated_resources.ore_collected -= blueprint.ore_robot.ore_required
        updated_robots.ore_robots += 1

    return (updated_resources, updated_robots)

def get_time_to_build_robot(robot_type: RobotType, blueprint: Blueprint, resources: Resources, robots: Robots) -> int:
    if robot_type == RobotType.GEODE_ROBOT:
        days_until_enough_ore = 0
        if resources.ore_collected < blueprint.geode_robot.ore_required:
            days_until_enough_ore = math.ceil((blueprint.geode_robot.ore_required - resources.ore_collected) / robots.ore_robots) + 1

        days_until_enough_obsidian = 0
        if resources.obsidian_collected < blueprint.geode_robot.obsidian_required:
            days_until_enough_obsidian = math.ceil((blueprint.geode_robot.obsidian_required - resources.obsidian_collected) / robots.obsidian_robots) + 1

        return max(1, days_until_enough_ore, days_until_enough_obsidian)

    if robot_type == RobotType.OBSIDIAN_ROBOT:
        days_until_enough_ore = 0
        if resources.ore_collected < blueprint.obsidian_robot.ore_required:
            days_until_enough_ore = math.ceil((blueprint.obsidian_robot.ore_required - resources.ore_collected) / robots.ore_robots) + 1

        days_until_enough_clay = 0
        if resources.clay_collected < blueprint.obsidian_robot.clay_required:
            days_until_enough_clay = math.ceil((blueprint.obsidian_robot.clay_required - resources.clay_collected) / robots.clay_robots) + 1

        return max(1, days_until_enough_ore, days_until_enough_clay)

    if robot_type == RobotType.CLAY_ROBOT:
        if resources.ore_collected < blueprint.clay_robot.ore_required:
            return math.ceil((blueprint.clay_robot.ore_required - resources.ore_collected) / robots.ore_robots) + 1
        return 1

    if robot_type == RobotType.ORE_ROBOT:
        if resources.ore_collected < blueprint.ore_robot.ore_required:
            return math.ceil((blueprint.ore_robot.ore_required - resources.ore_collected) / robots.ore_robots) + 1
        return 1

def get_available_robots_types(robot_cost: RobotCost, robots: Robots) -> Generator[RobotType, None, None]:
    """
    Returns robots we can build and need.
    We can build a robot if we have the necessary miners.
    We will build a robot if the amount mined per minute is less the maximum we can spend in a minute.
    """
    if robots.obsidian_robots > 0:
        yield RobotType.GEODE_ROBOT

    if robots.clay_robots > 0 and robots.obsidian_robots < robot_cost.obsidian_required:
        yield RobotType.OBSIDIAN_ROBOT

    if robots.clay_robots < robot_cost.clay_required:
        yield RobotType.CLAY_ROBOT

    if robots.ore_robots < robot_cost.ore_required:
        yield RobotType.ORE_ROBOT

def get_geodes_available(geode_collected: int, time_remaining: int) -> int:
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
    time_remaining -= 1
    max_yield_from_new_robots = time_remaining * (time_remaining + 1) / 2
    return max_yield_from_new_robots + geode_collected

def print_blueprint(blueprint: Blueprint) -> None:
    print(f'\nBlueprint: {blueprint.id}')
    print(f'- Each ore robot costs {blueprint.ore_robot.ore_required} ore')
    print(f'- Each clay robot costs {blueprint.clay_robot.ore_required} ore')
    print(f'- Each obsidian robot costs {blueprint.obsidian_robot.ore_required} ore and {blueprint.obsidian_robot.clay_required} clay')
    print(f'- Each geode robot costs {blueprint.geode_robot.ore_required} ore and {blueprint.geode_robot.obsidian_required} obsidian')

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
    is_star_two = True if sys.argv[2] == 'star2' else False
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, is_star_two, path)
