from blizzard_map import CachedBlizzardMap
from data_types import *
from typing import Generator
import dataclasses
import sys
import time


SLEEP_INTERVAL = 0.01


class RoutePlanner:
    def __init__(self) -> None:
        self._best_time = sys.maxsize
        self._best_history: list[Location] = []
        self._iterations = 0
        self._visited: set[tuple[int, int, int]] = {}
        self._get_neighbours = None

    def get_best_time_to_exit(
        self,
        explorer: Location,
        valley_map: ValleyMap,
        blizzard_map: CachedBlizzardMap,
        start_time: int=0,
        use_reverse_mode=False
    ) -> int:
        self._best_time = sys.maxsize
        self._best_history: list[Location] = []
        self._iterations = 0
        self._visited: set[tuple[int, int, int]] = {}

        self._get_neighbours = self._get_neighbours_reverse if use_reverse_mode else self._get_neighbours_forward
        best_time, history = self._calculate_best_time(explorer, valley_map, blizzard_map, history=[explorer], time=start_time)
        self._print_valley_map(valley_map, explorer, blizzard_map, history)

        return best_time

    def _calculate_best_time(
        self,
        explorer: Location,
        valley_map: ValleyMap,
        blizzard_map: CachedBlizzardMap,
        history: list[Location],
        time: int,
        depth=0
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


        self._iterations += 1
        print(f'- Best time: {self._best_time} | Current time: {time} | Iteration: {self._iterations}                               \r', end='')

        # 🔥🔥🔥 main loop 🔥🔥🔥
        if explorer == valley_map.exit:
            if time < self._best_time:
                self._best_time = time
                self._best_history = history
        else:
            current_blizzard_map = blizzard_map.get_map(time + 1)
            for neighbour in self._get_neighbours(explorer, valley_map.max_x, valley_map.max_y):
                if valley_map.locations[neighbour] == PATH and neighbour not in current_blizzard_map:
                    new_history = history.copy()
                    new_history.append(neighbour)
                    self._calculate_best_time(
                        neighbour,
                        valley_map,
                        blizzard_map,
                        new_history,
                        time + 1,
                        depth + 1
                        )

        return (self._best_time, self._best_history)

    def _get_neighbours_forward(self, location: Location, max_x: int, max_y: int) -> Generator[Location, None, None]:
        # down
        if location.y < max_y:
            yield dataclasses.replace(location, y = location.y + 1)

        # right
        if location.x < max_x - 1:
            yield dataclasses.replace(location, x = location.x + 1)

        # up
        if location.y > 0:
            yield dataclasses.replace(location, y = location.y - 1)

        # left
        if location.x > 1:
            yield dataclasses.replace(location, x = location.x - 1)

        # wait
        yield location

    def _get_neighbours_reverse(self, location: Location, max_x: int, max_y: int) -> Generator[Location, None, None]:
        # up
        if location.y > 0:
            yield dataclasses.replace(location, y = location.y - 1)

        # left
        if location.x > 1:
            yield dataclasses.replace(location, x = location.x - 1)

        # down
        if location.y < max_y:
            yield dataclasses.replace(location, y = location.y + 1)

        # right
        if location.x < max_x - 1:
            yield dataclasses.replace(location, x = location.x + 1)

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
            print(f'\n- Minute: {z}\n')

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
                        cell = '\033[1;30;101m \033[0m'

                    if current_location in history[z:z + 1]:
                        cell = '\033[1;30;103mE\033[0m'

                    print(cell, end='')
                print()
            time.sleep(SLEEP_INTERVAL)
        print()
