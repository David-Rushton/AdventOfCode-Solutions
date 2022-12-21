from __future__ import annotations
from ast import Tuple
from dataclasses import dataclass
import math
import sys

@dataclass
class Monkey:
    name: str
    value: int
    operator: str
    left_name: str
    left: Monkey
    right_name: str
    right: Monkey
    parent: Monkey
    is_human_or_ancestor: bool


def main(path: str) -> None:
    monkeys = parse_monkeys(path)
    root = monkeys['root']
    starting_human = monkeys['humn'].value
    value = get_value(root)
    target = get_value(root.right)

    # HACK! Assumes human is on the left.  Which is test and live input is true.  YMMV.
    human_value = solve_equation(root.left, target)

    print(f'\n== Monkey Maths\n- Star 1 result: {value}\n- Star 2 result: {human_value}\n- Human: {starting_human}')

def get_value(monkey: Monkey) -> int:
    if monkey.value == None:
        left = get_value(monkey.left)
        right = get_value(monkey.right)
        return get_monkey_operation(left, right, monkey.operator)
    else:
        return monkey.value

def is_equal(monkey: Monkey) -> bool:
    if monkey.value == None:
        left = get_value(monkey.left)
        right = get_value(monkey.right)

        if monkey.name == 'root':
            print(f'- left: {left}')
            print(f'- right: {right}')
            return left == right
        else:
            return get_monkey_operation(left, right, monkey.operator)
    else:
        return monkey.value

def solve_equation(monkey: Monkey, target: int) -> int:

    if monkey.name == 'humn':
        monkey.value = target
        return target

    if monkey.left.is_human_or_ancestor == monkey.right.is_human_or_ancestor:
        raise Exception('Human detected on both sides.  This is not possible.  Check input.')

    print()
    print_equation(monkey)
    print()


    left = get_value(monkey.left)
    right = get_value(monkey.right)

    inverse_operator = get_inverse_monkey_operator(monkey.operator)

    if monkey.left.is_human_or_ancestor:
        print_left = 'x'
        print_right = right
        inverse_left = target
        inverse_right = right
        next_monkey = monkey.left
    else:
        print_left = left
        print_right = 'x'
        inverse_left = left
        inverse_right = target
        next_monkey = monkey.right

    # for addition and multiplication we need to swap
    if monkey.operator in ('+', '*'):
        if monkey.right.is_human_or_ancestor:
            inverse_left = target
            inverse_right = left

    # negating subtraction
    if monkey.operator == '-':
        if monkey.left.is_human_or_ancestor:
            inverse_operator = '+'
        else:
            inverse_operator = '-'
            inverse_right = target
            print_right = target

    result = get_monkey_operation(inverse_left, inverse_right, inverse_operator)
    left_or_right = 'left' if monkey.left.is_human_or_ancestor else 'right'

    print (f'{target} = {print_left} {monkey.operator} {print_right}')
    print(f'  x = {inverse_left} {inverse_operator} {inverse_right}')
    print(f'  x = {result}')
    print(f'Human side: {left_or_right}')

    return solve_equation(next_monkey, result)

def print_equation(monkey: Monkey) -> None:
    if monkey.value == None:
        print('(', end='')
        print_equation(monkey.left)
        print(monkey.operator, end='')
        print_equation(monkey.right)
        print(')', end='')
        pass
    else:
        if monkey.is_human_or_ancestor:
            print(f'\033[1m\033[91mx\033[0m', end='')
        else:
            print(monkey.value, end='')

def get_inverse_monkey_operator(operator: str) -> str:
    if operator == '+':
        return '-'

    if operator == '-':
        return '+'

    if operator == '*':
        return '/'

    if operator == '/':
        return '*'

    raise Exception(f"Operator not supported: {operator}")

def get_monkey_operation(left: int, right: int, operator: str) -> int:
    if operator == '+':
        return left + right

    if operator == '-':
        return left - right

    if operator == '*':
        return left * right

    if operator == '/':
        return left / right

    raise Exception(f"Operator not supported: {operator}")

def parse_monkeys(path: str) -> dict[str, Monkey]:
    monkeys: dict[str, Monkey] = {}

    for monkey in open(path, 'r').read().splitlines():
        name_operation = monkey.split(':')
        name = name_operation[0]
        operation_elements =  name_operation[1][1:].split(' ')

        if len(operation_elements) == 1:
            monkeys[name] = Monkey(
                name=name,
                value=int(operation_elements[0]),
                operator=None,
                left_name=None,
                left=None,
                right_name=None,
                right=None,
                parent=None,
                is_human_or_ancestor=False
            )
        else:
            monkeys[name] = Monkey(
                name=name,
                value=None,
                operator=operation_elements[1],
                left_name=operation_elements[0],
                left=None,
                right_name=operation_elements[2],
                right=None,
                parent=None,
                is_human_or_ancestor=False
            )

    for monkey in monkeys:
        current = monkeys[monkey]
        if not current.left_name == None:
            current.left = monkeys[current.left_name]
            current.left.parent = current
            current.right = monkeys[current.right_name]
            current.right.parent = current

    current = monkeys['humn']
    while not current == None:
        current.is_human_or_ancestor = True
        current = current.parent

    return monkeys


if __name__ == '__main__':
    path = 'input.test.txt' if sys.argv[1] == 'test' else 'input.txt'
    main(path)
