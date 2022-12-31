from blizzard_map import CachedBlizzardMap
from data_types import *
from typing import Generator
import dataclasses
import sys
import time


SLEEP_INTERVAL = .01


class RoutePlanner:
    def __init__(self) -> None:
        self._best_time = sys.maxsize
        self._best_history: list[Location] = []
        self._iterations = 0
        self._visited: set[tuple[int, int, int]] = {}

    def get_best_time_to_exit(
        self,
        explorer: Location,
        valley_map: ValleyMap,
        blizzard_map: CachedBlizzardMap,
        start_time: int=0
    ) -> int:
        self._best_time = sys.maxsize
        self._best_history: list[Location] = []
        self._iterations = 0
        self._visited: set[tuple[int, int, int]] = {}

        best_time, history = self._calculate_best_time(explorer, valley_map, blizzard_map, history=[explorer], time=0)
        self._print_valley_map(valley_map, explorer, blizzard_map, history)

        return best_time

    def _calculate_best_time(
        self,
        explorer: Location,
        valley_map: ValleyMap,
        blizzard_map: CachedBlizzardMap,
        history: list[Location],
        time: int
    ) -> tuple[int, list[Location]]:
        # if we have already found a better route then stop exploring this one
        if self._best_time < sys.maxsize:
            best_possible_time = time + (valley_map.exit.y - explorer.y) + (valley_map.exit.x - explorer.x)
            if best_possible_time >= self._best_time:
                return (self._best_time, self._best_history)

        # if we have already visited this path then stop early
        key = (explorer.x, explorer.y, time)
        if key in self._visited:
            return
        self._visited[key] = True

        # # if the route isn't making progress then stop exploring it
        if time > 500:
            return

        self._iterations += 1
        print(f'- Best time: {self._best_time} | Current time: {time} | Iteration: {self._iterations}                               \r', end='')

        # 🔥🔥🔥 main loop 🔥🔥🔥
        if explorer == valley_map.exit:
            if time < self._best_time:
                self._best_time = time
                self._best_history = history
        else:
            routes_explored = 0
            while routes_explored == 0:
                for neighbour in self._get_neighbours(explorer, valley_map.max_x, valley_map.max_y):
                    current_blizzard_map = blizzard_map.get_map(time + 1)
                    if valley_map.locations[neighbour] == PATH and neighbour not in current_blizzard_map:
                        routes_explored += 1
                        new_history = history.copy()
                        new_history.append(neighbour)
                        self._calculate_best_time(neighbour, valley_map, blizzard_map, new_history, time + 1)

                # we have nowhere to go.
                # if a blizzard moves into our cell this route has failed
                if explorer in current_blizzard_map:
                    return

                history.append(explorer)
                time += 1

        return (self._best_time, self._best_history)

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
        yield location

    def _print_valley_map(
        self,
        valley_map: ValleyMap,
        explorer: Location,
        blizzard_map: CachedBlizzardMap,
        history: list[Location]=[],
        force_time: int=-1
    ) -> None:
        for z in range(max(1, len(history))):
            if z > 0:
                # moves the cursor back to the top.
                # we overprint minute (z) iteration over the top of the last.
                print(f'\033[{valley_map.max_y + 5}A')

            blizzards = blizzard_map.get_map(max(force_time, z))
            print(f'\nMinute: {z}\n')

            for y in range(valley_map.max_y + 1):
                for x in range(valley_map.max_x + 1):
                    current_location = Location(x, y)
                    cell = valley_map.locations[Location(x, y)]

                    if cell == WALL:
                        cell = f'\033[32m{WALL}\033[0m'

                    if current_location in blizzards:
                        cell = f'\033[94m{blizzards[current_location]}\033[0m'

                    if current_location == explorer:
                        cell = '\033[1;30;103mE\033[0m'

                    if current_location in history[0:z + 1]:
                        cell = '\033[1;30;101m-\033[0m'

                    if current_location in history[z:z + 1]:
                        if current_location in blizzards:
                            cell = '\033[1;30;101mX\033[0m'
                        else:
                            cell = '\033[1;30;103mE\033[0m'

                    print(cell, end='')
                print()
            time.sleep(SLEEP_INTERVAL)
