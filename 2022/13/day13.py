import sys


def main(path: str):
    pairs = parse_input(path)
    iteration = 1

    in_order_indexes = []
    in_order_count = 0
    not_in_order_count = 0

    while(len(pairs) > 0):
        left = pairs.pop(0)
        right = pairs.pop(0)
        is_in_order = True

        # for i in range(len(left)):
        print(f'\n== Pair {iteration} ==\n- Compare {left} vs {right}')
        if in_order(left, right, depth=0):
            in_order_indexes.append(iteration)
            in_order_count += 1
        else:
            not_in_order_count +=1

        iteration += 1

    print(f'\n\n== Results ==\n- Order Sum {in_order_indexes} = ({sum(in_order_indexes)})\n- In Order {in_order_count}\n- Not in Order {not_in_order_count}\n- Iterations {iteration - 1}\n')

def in_order(left, right, depth):

    padding = ''.rjust(depth)

    for i in range(len(left)):

        right_empty = i >= len(right)
        if (right_empty):
            print(f'{padding}  - Right side ran out of items, so inputs are not in the right order')
            return False

        print(f'{padding}  - Compare {left[i]} vs {right[i]}')

        if isinstance(left[i], int) and isinstance(right[i], int):
            if left[i] == right[i]:
                continue
            else:
                if left[i] < right[i]:
                    print(f'{padding}    - Left side is smaller, so inputs are in the right order')
                    return True
                else:
                    print(f'{padding}    - Right side is smaller, so inputs are not in the right order')
                    return False

        if isinstance(left[i], list) and isinstance(right[i], list):
            result = in_order(left[i], right[i], depth + 2)
            if result is not None:
                return result

        int_count = count_of_integers(left[i], right[i])
        if int_count == 1:
            if isinstance(left[i], int):
                print(f'{padding}    - Mixed types; convert left to {[left[i]]} and retry comparison')
                print(f'{padding}    - Compare {[left[i]]} vs {right[i]}')
                result = in_order([left[i]], right[i], depth + 2)
                if result is not None:
                    return result
            else:
                print(f'{padding}    - Mixed types; convert right to {[right[i]]} and retry comparison')
                print(f'{padding}    - Compare {left[i]} vs {[right[i]]}')
                result = in_order(left[i], [right[i]], depth + 2)
                if result is not None:
                    return result

    if len(left) < len(right):
        print(f'{padding}  - Left side ran out of items, so inputs are in the right order')
        return True

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
