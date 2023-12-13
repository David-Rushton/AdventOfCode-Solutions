import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Map:
    id: int
    rows: list[str]
    columns: list[str]


def main(is_test_mode: bool) -> None:
    print('point of incidence')
    print()

    maps = get_input(is_test_mode)
    summary = 0
    for map in maps:
        print(f'map {map.id}')
        for column in map.columns:
            print(column)
        row_summary = find_reflection(map.rows)
        if row_summary >= 0:
            summary += row_summary
        column_summary = find_reflection(map.columns)
        if column_summary >= 0:
            summary += column_summary * 100
        print(max(row_summary, column_summary))

        print()

    print()
    print(f'summary: {summary}')


def get_input(get_test: bool) -> Iterator[Map]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    lines = open(path, 'rt').read().splitlines()
    lines.append('')
    id = 0
    buffer = []
    for line in lines:
        if line == '':
            if len(buffer) > 0:
                yield Map(id, rows_to_columns(buffer), buffer)
            id += 1
            buffer.clear()
        else:
            buffer.append(line)


def find_reflection(image: list[str]) -> int:
    for i in range(0, len(image) -1):
        if image[i] == image[i + 1]:
            offset = 1
            while True:
                if i - offset >= 0 and i + 1 + offset < len(image):
                    if image[i - offset] == image[i + 1 + offset]:
                        offset += 1
                    else:
                        break
                else:
                    return i + 1
    return -1


def rows_to_columns(rows: list[str]) -> list[str]:
    result = []
    for x in range(0, len(rows[0])):
        result.append('')
        for y in range(0, len(rows)):
            result[x] += rows[y][x]
    return result


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
