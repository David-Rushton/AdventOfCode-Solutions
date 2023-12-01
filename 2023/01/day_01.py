import sys


def main(is_test_mode: bool) -> None:
    input = get_input(is_test_mode)

    total = 0
    for line in input:
        first_number = -1
        last_number = -1
        start = 0
        for i in range(len(line) + 1):
            value = get_value(line[start:i])
            if (value.isnumeric()):
                start = i - 1
                if (first_number == -1):
                    first_number = value
                    last_number = value
                else:
                    last_number = value
        print(f'{first_number}{last_number} {line}')
        total += int(f'{first_number}{last_number}')
    print('==========')
    print(total)

def get_value(buffer: str) -> str:
    if buffer.endswith('one') or buffer.endswith('1'):
        return '1'
    if buffer.endswith('two') or buffer.endswith('2'):
        return '2'
    if buffer.endswith('three') or buffer.endswith('3'):
        return '3'
    if buffer.endswith('four') or buffer.endswith('4'):
        return '4'
    if buffer.endswith('five') or buffer.endswith('5'):
        return '5'
    if buffer.endswith('six') or buffer.endswith('6'):
        return '6'
    if buffer.endswith('seven') or buffer.endswith('7'):
        return '7'
    if buffer.endswith('eight') or buffer.endswith('8'):
        return '8'
    if buffer.endswith('nine') or buffer.endswith('9'):
        return '9'

    return ''


def get_input(get_test: bool) -> str:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    return open(path, 'rt').read().splitlines()


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
