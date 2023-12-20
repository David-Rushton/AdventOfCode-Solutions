from enum import Enum
import sys
from dataclasses import dataclass
from typing import Iterator


HIGH = True
LOW = False


ON = True
OFF = False


class ModuleType(Enum):
    FLIP_FLOP = 0
    CONJUNCTION = 1
    BROADCAST = 2


@dataclass
class Module:
    name: str
    type: ModuleType
    state: bool
    memory: dict[str, bool]
    destinations: list[str]


def main(is_test_mode: bool) -> None:
    print('pulse propagation')
    print()
    print()

    modules = get_input(is_test_mode)
    button_pressed = 0
    high_pluses_send = 0
    low_pluses_send = 0

    watch_list: dict[str, list[str]] = {}
    result: dict[str, int] = {}
    for junction in ['rl', 'rd', 'qb', 'nn']:
        watch_list[junction] = []
        for module_name in modules:
            # if module_name == junction:
            #     for destination in modules[junction].destinations:
            #         if modules[destination].type == ModuleType.FLIP_FLOP:
            #             if destination not in watch_list[junction]:
            #                 watch_list[junction].append(destination)
            if junction in modules[module_name].destinations:
                if modules[module_name].type == ModuleType.FLIP_FLOP:
                    if module_name not in watch_list[junction]:
                        watch_list[junction].append(module_name)

    while (True):
        button_pressed += 1
        low_pluses_send += 1
        q: list[tuple(str, str, bool)] = [('button', 'broadcaster', LOW)]
        while len(q) > 0:
            (source, module_name, pulse) = q.pop(0)
            module = modules[module_name]

            if module.type == ModuleType.FLIP_FLOP:
                if pulse == HIGH:
                    continue
                else:
                    module.state = not module.state
                    pulse = HIGH if module.state == ON else LOW
            elif module.type == ModuleType.CONJUNCTION:
                module.memory[source] = pulse
                pulse = LOW
                for src in module.memory:
                    if module.memory[src] == LOW:
                        pulse = HIGH
                        break
                # if pulse == LOW:
                #     print('woo hoo')

            for destination in module.destinations:
                if pulse == HIGH:
                    high_pluses_send += 1
                if pulse == LOW:
                    low_pluses_send += 1
                if destination == 'rx' and pulse == LOW:
                    print('CHICKEN DINNER')
                    exit(0)
                if destination in ['output', 'rx']:
                    continue
                q.append((module_name, destination, pulse))

            r = f' {button_pressed} '
            for watching in watch_list:
                r += f' | {watching} > '
                c = 0
                if watching in result:
                    r+= '111111111111'
                else:
                    for m in watch_list[watching]:
                        r += '1' if modules[m].state else '0'
                        c += 1 if modules[m].state else 0
                    if c == len(watch_list[watching]):
                        result[watching] = button_pressed
            print(r, end='\r')

        if len(result) == len(watch_list):
            break

    print()
    print(f'button pressed {button_pressed}')
    print(f'pluses sent\n\tlow {low_pluses_send}\n\thigh {high_pluses_send}\n\tvalue {high_pluses_send * low_pluses_send}')
    print('factors')
    for factor in result:
        print(f'\t{factor} {result[factor]}')


def print_rx(button_pressed: int, rx: list[bool]) -> None:
    print(f' button pressed {button_pressed} rx len {len(rx)} ', end='')
    for pulse in rx:
        print('^' if pulse else 'v', end='')
    print()


def get_input(get_test: bool) -> dict[str, Module]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    modules: dict[str, Module] = {}
    for line in open(path, 'rt').read().splitlines():
        elements = line.split('->')
        destinations = elements[1].replace(' ', '').split(',')
        (name, type) = get_module_properties(elements[0].rstrip())
        modules[name] = Module(name, type, OFF, {}, destinations)
    for module in modules:
        for destination in modules[module].destinations:
            if destination in modules:
                if modules[destination].type == ModuleType.CONJUNCTION:
                    modules[destination].memory[module] = LOW
    return modules


def get_module_properties(name_and_type: str) -> (str, ModuleType):
    if name_and_type.startswith('%'):
        return (name_and_type[1:], ModuleType.FLIP_FLOP)
    elif name_and_type.startswith('&'):
        return (name_and_type[1:], ModuleType.CONJUNCTION)
    elif name_and_type == 'broadcaster':
        return (name_and_type, ModuleType.BROADCAST)
    else:
        raise Exception(f'module type not supported: {name_and_type}')


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
