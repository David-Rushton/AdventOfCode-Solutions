from data_types import *
import dataclasses


class CachedBlizzardMap:
    def __init__(self) -> None:
        self._blizzards: list[Blizzard] = []
        self._max_x = 0
        self._max_y = 0
        self._max_time = -1
        self._cache: dict[int, list[list[str]]] = {}

    def set_boundaries(self, max_x, max_y):
        self._max_x = max_x
        self._max_y = max_y

    def add_blizzard(self, location: Location, direction: Direction):
        self._blizzards.append(Blizzard(location, direction))

    def get_map(self, time: int) -> dict[Location, str]:
        while time > self._max_time:
            self._advance_time()
            self._cache[self._max_time] = self._map_blizzard()

        return self._cache[time]

    def _advance_time(self):
        if self._max_time >= 0:
            for blizzard in self._blizzards:
                if blizzard.direction not in DIRECTIONS:
                    raise Exception(f'Unsupported blizzard direction: {blizzard.direction}')

                if blizzard.direction == UP:
                    new_location = dataclasses.replace(blizzard.location, y = blizzard.location.y - 1)
                    if new_location.y == 0:
                        new_location = dataclasses.replace(blizzard.location, y = self._max_y - 1)

                if blizzard.direction == DOWN:
                    new_location = dataclasses.replace(blizzard.location, y = blizzard.location.y + 1)
                    if new_location.y == self._max_y:
                        new_location = dataclasses.replace(blizzard.location, y = 1)

                if blizzard.direction == LEFT:
                    new_location = dataclasses.replace(blizzard.location, x = blizzard.location.x - 1)
                    if new_location.x == 0:
                        new_location = dataclasses.replace(blizzard.location, x = self._max_x - 1)

                if blizzard.direction == RIGHT:
                    new_location = dataclasses.replace(blizzard.location, x = blizzard.location.x + 1)
                    if new_location.x == self._max_x:
                        new_location = dataclasses.replace(blizzard.location, x = 1)

                blizzard.location = new_location

        self._max_time += 1

    def _map_blizzard(self):
        result: dict[Location, str] = {}

        for blizzard in self._blizzards:
            if blizzard.location in result:
                current = result[blizzard.location]
                if current.isnumeric():
                    current = int(current) + 1
                else:
                    current = 2
                result[blizzard.location] = str(current)
            else:
                result[blizzard.location] = blizzard.direction

        return result;
