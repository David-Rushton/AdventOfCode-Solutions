from dataclasses import dataclass
import sys
import os
import time

@dataclass
class Monkey:
    name: str
    items: list[int]
    operation_action: str
    operation_value: str
    test_divider: int
    when_true: int
    when_false: int


def main(path: str, starTwo: bool):
    monkeys = parse_monkeys(open(path, 'r').read().splitlines())
    round = 1
    inspected = {}
    roundLimit = 10000 if starTwo else 20

    os.system('cls')

    devisor = 1
    for monkey in monkeys:
        devisor *= monkey.test_divider

    while round <= roundLimit:
        for monkey in monkeys:
            for index in range(len(monkey.items)):
                if monkey.name not in inspected:
                    inspected[monkey.name] = 0
                inspected[monkey.name] += 1
                item = monkey.items[index]
                item = calculate_worry_level(item, monkey.operation_action, monkey.operation_value, starTwo, devisor)

                if item % monkey.test_divider == 0:
                    monkeys[monkey.when_true].items.append(item)
                else:
                    monkeys[monkey.when_false].items.append(item)

            monkey.items = []

        print_monkeys(round, monkeys, inspected, starTwo, round == roundLimit)
        round += 1

    print_monkey_business(inspected)

def print_monkey_business(inspected: dict[str, int]):
    sorted_inspections = sorted(inspected.items(), key=lambda item: item[1], reverse=True)
    monkey_business = sorted_inspections[0][1] * sorted_inspections[1][1]
    print(f'\nMonkey business: {sorted_inspections[0][1]} * {sorted_inspections[1][1]} = {monkey_business}')


def print_monkeys(round: int, monkeys: list[Monkey], inspected: dict[str, int], starTwo: bool, lastRound: bool):
    print(f'\033[1;1HRound: {round}\n')

    if not starTwo or lastRound:
        for monkey in monkeys:
            items = ", ".join(map(str, monkey.items))
            print(f'{monkey.name} {items}                                 ')

        print()

    if not starTwo or lastRound or round % 500 == 0:
        for key in inspected:
            print(f'{key} inspected items {inspected[key]} times          ')

        print()
        time.sleep(.3)


def calculate_worry_level(starting_level: int, operation: str, value: str, starTwo: bool, devisor: int):
    value_number = starting_level if value == 'old' else int(value)

    if operation == '+':
        result = (int(starting_level) + int(value_number))

    if operation == '*':
        result = (int(starting_level) * int(value_number))

    if starTwo:
        return result % devisor
    else:
        return result // 3


def parse_monkeys(text: list[str]) -> list[Monkey]:
    result = []
    name = operation_action = ''
    items = []
    operation_value = test_divider = when_true = when_false = 0
    for line in text:

        if line.startswith('Monkey'):
            name = line

        if line.startswith('  Starting items: '):
            items = line.replace('  Starting items: ', '').split(', ')

        if line.startswith('  Operation: new = old '):
            elements = line.replace('  Operation: new = old ', '').split(' ')
            operation_action = elements[0]
            operation_value = elements[1]

        if line.startswith('  Test: divisible by'):
            test_divider = int(line.replace('  Test: divisible by', ''))

        if line.startswith('    If true: throw to monkey '):
            when_true = int(line.replace('    If true: throw to monkey ', ''))

        if line.startswith('    If false: throw to monkey '):
            when_false = int(line.replace('    If false: throw to monkey ', ''))

        if line == '':
            result.append(Monkey(
                name,
                items,
                operation_action,
                operation_value,
                test_divider,
                when_true,
                when_false
            ))

    return result;


if __name__ == '__main__':
    path = 'input.test.txt' if sys.argv[1] == 'test' else 'input.txt'
    starTwo = sys.argv[2] == 'star2'
    main(path, starTwo)
