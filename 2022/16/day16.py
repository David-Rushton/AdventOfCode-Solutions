from __future__ import annotations
from dataclasses import dataclass
import sys

TOTAL_TIME_IN_MINUTES = 30

@dataclass
class Cave:
    name: str
    flow_rate: int
    connected_to_names: list[str]
    connected_to: list[Cave]

@dataclass
class Route:
    path: str
    value: int


def main(is_test_mode: bool, path: str) -> None:
    caves = parse_caves(path)
    required_caves = get_required_caves(caves)
    cave_aa = caves['AA']
    seed_route = score_path(get_seed_path(cave_aa), caves)

    print(seed_route)

    exit(0)

def get_best_route(cave: Cave, seed_route: Route) -> Route:
    tabu = [seed_route]

    # TODO: Stopping condition
    # TODO: https://en.wikipedia.org/wiki/Tabu_search#Pseudocode
    while True:




def get_seed_path(cave: Cave) -> str:
    steps = []

    while len(steps) < 30:
        steps.append(cave.name)
        cave = cave.connected_to[0]

    return '>'.join(steps)

def score_path(path: str, caves: dict[str, Cave]) -> Route:
    caves_names = path.split('>')
    minutes_passed = 0
    cave = None
    score = 0
    valves_opened = []

    for caves_name in caves_names:
        cave = caves[caves_name]
        minutes_passed += 1

        if minutes_passed >= 29:
            break

        if cave.flow_rate > 0:
            if cave not in valves_opened:
                minutes_passed += 1
                valves_opened.append(cave)
                score += cave.flow_rate * (TOTAL_TIME_IN_MINUTES - minutes_passed)
                print(f'Time: {minutes_passed}')
                print(f'opened value {cave.name}, with flow rate of {cave.flow_rate} adding {cave.flow_rate * (TOTAL_TIME_IN_MINUTES - minutes_passed)} to score')

    return Route(path, score)

def get_required_caves(caves: dict[str, Cave]) -> list[Cave]:
    result = []

    for cave_name in caves:
        if caves[cave_name].flow_rate > 0:
            result.append(caves[cave_name])

    return result

def parse_caves(path: str) -> dict[str, Cave]:
    # format:
    # Valve <NAME> has flow rate=<FLOW_RATE>; tunnels lead to valves <NAME1>, ..., <NAMEn>
    # 0     1      2   3    4                 5       6    7  8      9+
    caves = open(path, 'r').read().replace(';', '').replace(',', '').replace('rate=', '').splitlines()
    result: dict[str, Cave] = {}

    for cave in caves:
        elements = cave.split(' ')
        name = elements[1]
        result[name] = Cave(
            name=name,
            flow_rate=int(elements[4]),
            connected_to_names=elements[9:],
            connected_to=[]
        )

    for cave_name in result:
        cave = result[cave_name]
        for connected_to_name in cave.connected_to_names:
            cave.connected_to.append(result[connected_to_name])

    return result


if __name__ == '__main__':
    is_test_mode = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, path)
