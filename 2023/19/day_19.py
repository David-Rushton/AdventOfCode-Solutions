import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Part:
    values: dict[str, int]


@dataclass
class Check:
    input: str
    operator: str
    value: int
    workflow: int


@dataclass
class Workflow:
    name: str
    checks: list[Check]


def main(is_test_mode: bool) -> None:
    print('aplenty')
    print()

    (parts, workflows) = get_input(is_test_mode)
    approved: list[Part] = []
    rejected: list[Part] = []
    for part in parts:
        print(f' checking part x={part.values["x"]}, m={part.values["m"]}, a={part.values["a"]}, s={part.values["s"]}\t', end='')
        if is_part_approved(part, 'in', workflows):
            print('approved')
            approved.append(part)
        else:
            print('rejected')
            rejected.append(part)

    print()
    print(f'approved parts rating {sum(part.values[value] for value in part.values for part in approved)}')


def is_part_approved(part: Part, workflow: str, workflows: dict[str, Workflow]) -> bool:
    for check in workflows[workflow].checks:
        next = None
        if check.operator == '<':
            if part.values[check.input] < check.value:
                next = check.workflow
        elif check.operator == '>':
            if part.values[check.input] > check.value:
                next = check.workflow
        else:
            next = check.workflow

        if next == None:
            continue
        elif next == 'A':
            return True
        elif next == 'R':
            return False
        else:
            return is_part_approved(part, next, workflows)
    print(part)
    raise('unable to approve or reject part')


def get_input(get_test: bool) -> tuple[list[Part], dict[str, Workflow]]:
    test_path = './input.test.txt'
    prod_path = './input.txt'
    path = test_path if get_test else prod_path
    parts: list[Part] = []
    workflows: dict[str, Workflow] = {}
    lines = open(path, 'rt').read().splitlines()
    for line in lines:
        if line.startswith('{'):
            part = Part({})
            parts.append(part)
            for element in line.replace('{', '').replace('}', '').split(','):
                part.values[element[0]] = int(element[2:])
        else:
            if len(line) > 0:
                elements = line.replace('}', '').split('{')
                workflow = Workflow(elements[0], [])
                for element in elements[1].split(','):
                    if element.find(':') == -1:
                        workflow.checks.append(Check('*', '*', -1, element))
                    else:
                        sub_elements = element.split(':')
                        workflow.checks.append(Check(
                            element[0],
                            element[1],
                            int(sub_elements[0][2:]),
                            sub_elements[1]))
                workflows[workflow.name] = workflow
    return (parts, workflows)


if __name__ == '__main__':
    is_test_mode = sys.argv.count('test') > 0
    main(is_test_mode)
