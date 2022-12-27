from __future__ import annotations
from dataclasses import dataclass
import sys
import math


TIME_LIMIT = 26


@dataclass
class Cave:
    name: str
    flow_rate: int

@dataclass
class CaveRouteStep:
    name: str
    distance: int
    previous: CaveRouteStep

@dataclass
class CaveRoute:
    source_name: str
    destination_name: str
    destination_flow_rate: int
    path: list[str]
    steps: int
    score: int

class CaveRouter:
    def __init__(self, caves: dict[str, Cave], cave_map: dict[str, Cave]) -> None:
        self._caves = caves
        self._cave_map = cave_map
        self._cache: dict[tuple[str, str], CaveRoute] = {}

    def get_route(self, source_name: str, destination_name: str, remaining_time: int) -> CaveRoute:
        """
        Returns the shortest path between two caves.
        """

        if (source_name, destination_name) not in self._cache:
            self._add_route_to_cache(source_name, destination_name)

        route = self._cache[(source_name, destination_name)]

        return self._copy_and_add_score(route, remaining_time)

    def _add_route_to_cache(self, source_name: str, destination_name: str) -> CaveRoute:
        queue: list[CaveRouteStep] = []
        queue_index: dict[str, CaveRouteStep] = {}
        visited = {}

        for cave_name in self._caves:
            distance = 0 if cave_name == source_name else sys.maxsize
            cave_route = CaveRouteStep(cave_name, distance, previous=None)
            queue.append(cave_route)
            queue_index[cave_name] = cave_route

        while len(queue) > 0:
            queue.sort(key=lambda x: x.distance)
            current = queue.pop(0)
            visited[current.name] = current

            for neighbour_name in self._cave_map[current.name]:
                neighbour = queue_index[neighbour_name]
                if neighbour in queue:
                    neighbour_distance = current.distance + 1
                    if neighbour_distance < neighbour.distance:
                        neighbour.distance = neighbour_distance
                        neighbour.previous = current

            if current.name == destination_name:
                break

        self._cache[(source_name, destination_name)] = CaveRoute(
            source_name,
            destination_name,
            destination_flow_rate=self._caves[destination_name].flow_rate,
            path=get_path(current),
            steps=current.distance,
            score=0
        )

    def _copy_and_add_score(self, original: CaveRoute, remaining_time: int) -> CaveRoute:
        return CaveRoute(
            source_name=original.source_name,
            destination_name=original.destination_name,
            destination_flow_rate=original.destination_flow_rate,
            path=original.path,
            steps=original.steps,
            score=self._score_path(remaining_time, original.steps, original.destination_flow_rate)
        )

    def _score_path(self, remaining_time: int, steps: int, flow_rate: int) -> int:
        factor = remaining_time - steps - 1
        if factor > 0:
            return factor * flow_rate
        else:
            return 0


def main(is_test_mode: bool, path: str) -> None:
    caves, cave_map = parse_caves(path)
    required_caves = set(get_required_caves(caves))
    human_starting_cave_name ='AA'
    elephant_starting_cave_name ='AA'
    cave_router = CaveRouter(caves, cave_map)

    best_routes = get_best_route(
        human_starting_cave_name,
        elephant_starting_cave_name,
        caves,
        cave_map,
        required_caves,
        cave_router
    )

    print('\n== Result ==')
    print(f'- Human')
    print(f'  - Best path: {" > ".join(best_routes[0].path[0])}')
    print(f'  - Best score: {best_routes[0].score}')
    print(f'- Elephant')
    print(f'  - Best path: {" > ".join(best_routes[1].path[0])}')
    print(f'  - Best score: {best_routes[1].score}')
    print(f'- Elephant')
    print(f'  - Best score: {best_routes[0].score + best_routes[1].score}')
    print()

    exit(0)

def get_best_route(
    human_starting_cave_name: str,
    elephant_starting_cave_name: str,
    caves: dict[str, Cave],
    cave_map: dict[str, list[str]],
    required_caves: set[str],
    cave_router: CaveRouter,
    human_remaining_time: int=TIME_LIMIT,
    elephant_remaining_time: int=TIME_LIMIT,
    human_score=0,
    elephant_score=0,
    human_path: str='',
    elephant_path: str='',
    human_best_route: CaveRoute=CaveRoute('AA', 'AA', 0, [], -1, -1),
    elephant_best_route: CaveRoute=CaveRoute('AA', 'AA', 0, [], -1, -1),
    depth: int=0
) -> tuple[CaveRoute, CaveRoute]:

    delimiter = ' > '

    if depth == 0:
        human_path = human_starting_cave_name
        elephant_path = elephant_starting_cave_name
        print(f'\n== Routes ==\n')

    if len(required_caves) == 0:
        best_score = human_best_route.score + elephant_best_route.score
        current_score = human_score + elephant_score

        if current_score > best_score:
            print(f'New best route found with a score of {current_score}')
            print(f'  - Human route   : {human_path}')
            print(f'  - Elephant route: {elephant_path}')

            human_best_route.source_name='AA',
            human_best_route.destination_name=human_starting_cave_name,
            human_best_route.destination_flow_rate=0,
            human_best_route.path=human_path.split(delimiter),
            human_best_route.steps=-1,
            human_best_route.score=human_score

            elephant_best_route.source_name='AA',
            elephant_best_route.destination_name=elephant_starting_cave_name,
            elephant_best_route.destination_flow_rate=0,
            elephant_best_route.path=elephant_path.split(delimiter),
            elephant_best_route.steps=-1,
            elephant_best_route.score=elephant_score

    if len(required_caves) == 1:
        # try both paths
        # whoever can reach the valve in the sortest number of steps wins
        human_destination_cave_name = required_caves.pop()
        human_route = cave_router.get_route(human_starting_cave_name, human_destination_cave_name, human_remaining_time)

        elephant_destination_cave_name = human_destination_cave_name
        elephant_route = cave_router.get_route(elephant_starting_cave_name, elephant_destination_cave_name, elephant_remaining_time)

        if human_route.score >= elephant_route.score:
            get_best_route(
                human_starting_cave_name=human_destination_cave_name,
                elephant_starting_cave_name='',
                caves=caves,
                cave_map=cave_map,
                required_caves={},
                cave_router=cave_router,
                human_remaining_time=0,
                elephant_remaining_time=0,
                human_score=human_score + human_route.score,
                elephant_score=elephant_score,
                human_path=f'{human_path}{delimiter}{human_destination_cave_name}',
                elephant_path=elephant_path,
                human_best_route=human_best_route,
                elephant_best_route=elephant_best_route,
                depth=depth + 1
            )
        else:
            get_best_route(
                    human_starting_cave_name='',
                    elephant_starting_cave_name=elephant_destination_cave_name,
                    caves=caves,
                    cave_map=cave_map,
                    required_caves={},
                    cave_router=cave_router,
                    human_remaining_time=0,
                    elephant_remaining_time=0,
                    human_score=human_route.score,
                    elephant_score=elephant_score + elephant_route.score,
                    human_path=human_path,
                    elephant_path=f'{elephant_path}{delimiter}{elephant_destination_cave_name}',
                    human_best_route=human_best_route,
                    elephant_best_route=elephant_best_route,
                    depth=depth + 1
                )
    else:
        for destination_names in get_destination_pairs(required_caves):
            remaining_caves = required_caves.copy()

            human_destination_cave_name = destination_names[0]
            remaining_caves.remove(human_destination_cave_name)
            human_route = cave_router.get_route(human_starting_cave_name, human_destination_cave_name, human_remaining_time)

            elephant_destination_cave_name = destination_names[1]
            remaining_caves.remove(elephant_destination_cave_name)
            elephant_route = cave_router.get_route(elephant_starting_cave_name, elephant_destination_cave_name, elephant_remaining_time)

            combined_remaining_time = human_remaining_time + elephant_remaining_time
            max_remaining_score = human_score + human_route.score + elephant_score + elephant_route.score + get_remaining_score(caves, required_caves, combined_remaining_time)
            if human_best_route.score + elephant_best_route.score > max_remaining_score:
                return

            get_best_route(
                human_destination_cave_name,
                elephant_destination_cave_name,
                caves,
                cave_map,
                remaining_caves,
                cave_router,
                human_remaining_time=human_remaining_time - human_route.steps - 1,
                elephant_remaining_time=elephant_remaining_time - elephant_route.steps - 1,
                human_score=human_score + human_route.score,
                elephant_score=elephant_score + elephant_route.score,
                human_path=f'{human_path}{delimiter}{human_destination_cave_name}',
                elephant_path=f'{elephant_path}{delimiter}{elephant_destination_cave_name}',
                human_best_route=human_best_route,
                elephant_best_route=elephant_best_route,
                depth=depth + 1
            )

    if depth == 0:
        return (human_best_route, elephant_best_route)

def get_destination_pairs(required_caves: set[str]) -> tuple[str, str]:
    for human in required_caves:
        for elephant in required_caves:
            if human != elephant:
                yield (human, elephant)

def get_remaining_score(caves: dict[str, Cave], required_caves: set[str], remaining_time: int) -> int:
    """
    This could be improved.  If we limited the number of caves considered to the maximum number
    we could visit in the remaining time (remaining_time / 2).
    """
    score = 0

    for cave_name in required_caves:
        score += caves[cave_name].flow_rate * remaining_time

    return score

def get_path(step: CaveRouteStep) -> list[str]:
    result = []

    while step is not None:
        result.insert(0, step.name)
        step = step.previous

    return result

def get_required_caves(caves: dict[str, Cave]) -> list[str]:
    for cave_name in caves:
        if caves[cave_name].flow_rate > 0:
            yield cave_name

def parse_caves(path: str) -> tuple[dict[str, Cave], dict[str, list[str]]]:
    # format:
    # Valve <NAME> has flow rate=<FLOW_RATE>; tunnels lead to valves <NAME1>, ..., <NAMEn>
    # 0     1      2   3    4                 5       6    7  8      9+
    raw_caves = open(path, 'r').read().replace(';', '').replace(',', '').replace('rate=', '').splitlines()
    caves = {}
    cave_map = {}

    for cave in raw_caves:
        elements = cave.split(' ')
        name = elements[1]

        caves[name] = Cave(name, flow_rate=int(elements[4]))

        cave_map[name] = []
        for neighbour in elements[9:]:
            cave_map[name].append(neighbour)

    return (caves, cave_map)


if __name__ == '__main__':
    is_test_mode = sys.argv[1] == 'test'
    path = 'input.test.txt' if is_test_mode else 'input.txt'
    main(is_test_mode, path)
