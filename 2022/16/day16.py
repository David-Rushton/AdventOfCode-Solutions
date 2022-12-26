from __future__ import annotations
from dataclasses import dataclass
import sys


TIME_LIMIT = 30


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
    required_caves = list(get_required_caves(caves))
    starting_cave_name ='AA'
    cave_router = CaveRouter(caves, cave_map)

    routes = get_routes(starting_cave_name, caves, cave_map, required_caves, cave_router)
    routes.sort(key=lambda x: x.score)
    for route in routes:
        print(f'- {" > ".join(route.path)}= {route.score}')

    # print(f'Best score: {best}')
    exit(0)

def get_routes(
    starting_cave_name: str,
    caves: dict[str, Cave],
    cave_map: dict[str, list[str]],
    required_caves: list[str],
    cave_router: CaveRouter,
    remaining_time: int=TIME_LIMIT,
    score=0,
    path: str='',
    routes: list[CaveRoute]=[],
    depth: int=0
) -> list[CaveRoute]:

    padding = ''.ljust(depth * 2)
    delimiter = ' > '

    if depth == 0:
        path = starting_cave_name
        print(f'\n== Routes ==\n')

    if len(required_caves) == 0 or remaining_time <= 2:
        # if score > best_score:
        #     print(f'New best route: {path} with a score of {score}')
        #     best_score = score

        routes.append(CaveRoute(
            source_name='AA',
            destination_name=starting_cave_name,
            destination_flow_rate=0,
            path=path.split(delimiter),
            steps=-1,
            score=score))
        print(f'\r{len(routes)}', end='')

    # TODO:
    # We need a heuristic or two to speed things up.
    # Cull if:
    # - sum of remaining caves is not enough to beat best
    # - outside of top x
    # if len(routes) > 100:
        # routes.sort(key=lambda x: x.score)
        # best_score = { score: 1500 }
    # current_score = score + get_remaining_score(caves, required_caves, remaining_time)
    # if best_score > 0 and current_score < best_score:
    #     # print(f'\nCannot beat score {best_score.score} with {current_score}')
    #     # print(f'- Best: {best_score.path}')
    #     # print(f'- Current: {path}\n')
    #     return

    for i in range(len(required_caves)):
        required_cave = required_caves[i]
        remaining_caves = required_caves[:i] + required_caves[i +1:]
        route = cave_router.get_route(starting_cave_name, required_cave, remaining_time)
        route_score = score + route.score

        get_routes(
            required_cave,
            caves,
            cave_map,
            remaining_caves,
            cave_router,
            remaining_time=remaining_time - route.steps - 1,
            score=route_score,
            path=f'{path} > {required_cave}',
            # best_score=max(best_score, route_score),
            routes=routes,
            depth=depth + 1
        )

    if depth == 0:
        return routes

# def get_best_route(cave: Cave, seed_route: Route) -> Route:
#     tabu = [seed_route]

#     # TODO: Stopping condition
#     # TODO: https://en.wikipedia.org/wiki/Tabu_search#Pseudocode
#     while True:
#         pass

def get_remaining_score(caves: dict[str, Cave], required_caves: list[str], remaining_time) -> int:
    score = 0

    for cave_name in required_caves:
        score += caves[cave_name].flow_rate * remaining_time

    return score

def get_seed_path(cave: Cave) -> str:
    steps = []

    while len(steps) < 30:
        steps.append(cave.name)
        cave = cave.connected_to[0]

    return '>'.join(steps)

def get_route(
    source_name: str,
    destination_name: str,
    caves: dict[str, Cave],
    cave_map: dict[str, list[str]],
    remaining_time: int
) -> CaveRoute:
    """
    Finds the shortest path between two caves.
    """
    queue: list[CaveRouteStep] = []
    queue_index: dict[str, CaveRouteStep] = {}
    visited = {}

    for cave_name in caves:
        distance = 0 if cave_name == source_name else sys.maxsize
        cave_route = CaveRouteStep(cave_name, distance, previous=None)
        queue.append(cave_route)
        queue_index[cave_name] = cave_route

    while len(queue) > 0:
        queue.sort(key=lambda x: x.distance)
        current = queue.pop(0)
        visited[current.name] = current

        for neighbour_name in cave_map[current.name]:
            neighbour = queue_index[neighbour_name]
            if neighbour in queue:
                neighbour_distance = current.distance + 1
                if neighbour_distance < neighbour.distance:
                    neighbour.distance = neighbour_distance
                    neighbour.previous = current

        if current.name == destination_name:
            break

    route = CaveRoute(
        source_name,
        destination_name,
        path=get_path(current),
        steps=current.distance,
        score=score_path(remaining_time, current.distance, caves[current.name].flow_rate)
    )

    return route

def get_path(step: CaveRouteStep) -> list[str]:
    result = []

    while step is not None:
        result.insert(0, step.name)
        step = step.previous

    return result

# def score_path(remaining_time: int, steps: int, flow_rate: int) -> int:
#     factor = remaining_time - steps - 1
#     if factor > 0:
#         return factor * flow_rate
#     else:
#         return 0

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
