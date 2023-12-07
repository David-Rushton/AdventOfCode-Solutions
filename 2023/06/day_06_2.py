import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Race:
    time_ms: int
    record_distance: int
    wins: int


def main(is_test_mode: bool) -> None:
    races = get_input(is_test_mode)
    total_wins = 1

    print('wait for it')
    print('-----------')

    for race in races:
        for held_ms in range(1, race.time_ms):
            speed = held_ms
            distance_travelled = speed * (race.time_ms - held_ms)
            if (distance_travelled > race.record_distance):
                race.wins += 1
        if race.wins > 0:
            total_wins = total_wins * race.wins
        print(f' race time: {race.time_ms}.  distance beat: {race.wins}.')

    print('-----------')
    print(f'Total wins: {total_wins}')


def get_input(get_test: bool) -> Iterator[Race]:
    if get_test:
        yield Race(71530, 940200, 0)
    else:
        yield Race(54946592, 302147610291404, 0)

if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
