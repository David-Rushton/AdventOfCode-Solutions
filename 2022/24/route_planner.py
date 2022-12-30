from blizzard_map import CachedBlizzardMap
from data_types import *
from typing import Generator
import dataclasses
import sys


class RoutePlanner:
    def __init__(self) -> None:
        self._best_time = sys.maxsize
        self._iterations = 0

    def get_best_time_to_exit(self, explorer: Location, valley_map: ValleyMap, blizzard_map: CachedBlizzardMap) -> int:
        self._best_time = sys.maxsize
        self._iterations = 0
        return self._calculate_best_time(explorer, valley_map, blizzard_map, time=0)

    def _calculate_best_time(self, explorer: Location, valley_map: ValleyMap, blizzard_map: CachedBlizzardMap, time: int) -> int:
        # there are better routes, don't explore this one any further

        best_possible_time = time + (valley_map.exit.y - explorer.y) + (valley_map.exit.x - explorer.x)
        if best_possible_time > self._best_time:
            return

        self._iterations += 1
        print(f'- Best time: {self._best_time} | Current time: {time} | Iteration: {self._iterations}                               \r', end='')

        # print()
        # self._print_valley_map(valley_map, explorer, blizzard_map.get_map(time))
        # input()

        # 🔥🔥🔥 main loop 🔥🔥🔥
        if explorer == valley_map.exit:
            if time < self._best_time:
                self._best_time = time
                # print(f'- New best time found: {self._best_time} minutes   \r', end='')
                # print()
                # print_valley_map(valley_map, explorer, blizzard_map.get_map(time + 1))
                # input()
        else:
            # print(f'- Current time: {time} minutes   \r', end='')
            routes_explored = 0
            while routes_explored == 0:
                for neighbour in self._get_neighbours(explorer, valley_map.max_x, valley_map.max_y):
                    if valley_map.locations[neighbour] == PATH and neighbour not in blizzard_map.get_map(time + 1):
                        # print()
                        # print_valley_map(valley_map, explorer, blizzard_map.get_map(time + 1))
                        # input()
                        routes_explored += 1
                        self._calculate_best_time(neighbour, valley_map, blizzard_map, time + 1)
                time += 1

        if time==0:
            return self._best_time

    def _get_neighbours(self, location: Location, max_x: int, max_y: int) -> Generator[Location, None, None]:
        # down
        if location.y < max_y:
            yield dataclasses.replace(location, y = location.y + 1)

        # right
        if location.x < max_x - 1:
            yield dataclasses.replace(location, x = location.x + 1)

        # left
        if location.x > 1:
            yield dataclasses.replace(location, x = location.x - 1)

        # up
        if location.y > 1:
            yield dataclasses.replace(location, y = location.y - 1)

        # wait
        # yield location

    def _print_valley_map(self, valley_map: ValleyMap, explorer: Location, blizzards: dict[Location, str]) -> None:
        for y in range(valley_map.max_y + 1):
            for x in range(valley_map.max_x + 1):
                current_location = Location(x, y)
                cell = valley_map.locations[Location(x, y)]

                if cell == WALL:
                    cell = f'\033[32m{WALL}\033[0m'

                if current_location in blizzards:
                    cell = f'\033[94m{blizzards[current_location]}\033[0m'

                # Always consider our explorer's location last.
                # To ensure it is printed when a cell contain more than 1 item.
                if current_location == explorer:
                    cell = '\033[1;30;103mP\033[0m'

                print(cell, end='')
            print()
