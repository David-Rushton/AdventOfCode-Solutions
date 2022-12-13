import sys


def main(path: str):
    pairs = parse_input(path)
    iteration = 1

    while(len(pairs) > 0):
        left = pairs.pop(0)
        right = pairs.pop(0)
        is_in_order = True

        # for i in range(len(left)):
        print(f'\n== Pair {iteration} ==\n- Compare {left} vs {right}')
        is_in_order = in_order(left, right, depth=2)


        iteration += 1

def in_order(left, right, depth):

    padding = ''.rjust(depth)

    for i in range(len(left)):
        print(f'{padding}- Compare {left[i]} vs {right[i]}')

        if isinstance(left[i], int) and isinstance(right[i], int):
            if left[i] == right[i]:
                continue
            else:
                if left[i] < right[i]:
                    print(f'{padding}  - Left side is smaller, so inputs are in the right order')
                    return True
                else:
                    print(f'{padding}  - Right side is smaller, so inputs are not in the right order')
                    return False

        if isinstance(left[i], list) and isinstance(right[i], list):
            return in_order(left[i], right[i], depth + 2)

        int_count = count_of_integers(left[i], right[i])
        if int_count == 1:
            if isinstance(left[i], int):
                print(f'{padding}  - Mixed types; convert left to {[left[i]]} and retry comparison')
                return in_order([left[i]], right[i], depth + 2)
            else:
                print(f'{padding}  - Mixed types; convert right to {[right[i]]} and retry comparison')
                return in_order(left[i], [right[i]], depth + 2)

def count_of_integers(left, right):
    left_int_count = 1 if isinstance(left, int) else 0
    right_int_count = 1 if isinstance(right, int) else 0

    return left_int_count + right_int_count

def parse_input(path):
    lines = open(path, 'r').read().splitlines()
    result = []

    for line in lines:
        if len(line) > 0:
            result.append(eval(line))

    return result




if __name__ == '__main__':
    path = 'day13.input.txt.test' if sys.argv[1] == 'test' else 'day13.input.txt'
    main(path)

