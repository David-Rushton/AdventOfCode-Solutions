from dataclasses import dataclass
from enum import Enum


EXPLORER = 'E'
WALL = '#'
PATH = '.'
UP = '^'
DOWN = 'v'
LEFT = '<'
RIGHT = '>'
DIRECTIONS = [UP, DOWN, LEFT, RIGHT]

class Direction(Enum):
    UP = UP
    DOWN = DOWN
    LEFT = LEFT
    RIGHT = RIGHT

@dataclass(eq=True, frozen=True)
class Location:
    x: int
    y: int

@dataclass
class ValleyMap:
    entry: Location
    exit: Location
    max_x: int
    max_y: int
    locations: dict[Location, str]

@dataclass
class Blizzard:
    location: Location
    direction: Direction
