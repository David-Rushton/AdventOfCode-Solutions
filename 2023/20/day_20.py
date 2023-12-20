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

    modules = get_input(is_test_mode)
    button_pressed = 0
    high_pluses_send = 0
    low_pluses_send = 0

    for _ in range(1000):
        button_pressed += 1
        low_pluses_send += 1
        print()
        print(f' button pressed {button_pressed} times')
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

            print(f'  {module_name} sending {"high" if pulse == HIGH else "low"}: ', end='')
            for destination in module.destinations:
                if pulse == HIGH:
                    high_pluses_send += 1
                if pulse == LOW:
                    low_pluses_send += 1
                if destination not in modules:
                    print(f' {destination} ', end='')
                    continue
                print(destination, end=', ')
                q.append((module_name, destination, pulse))
            print()

        print()
        print(f'button pressed {button_pressed}')
        print(f'pluses sent\n\tlow {low_pluses_send}\n\thigh {high_pluses_send}\n\tvalue {high_pluses_send * low_pluses_send}')


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
