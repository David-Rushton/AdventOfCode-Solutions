import sys
from dataclasses import dataclass
from typing import Iterator


@dataclass
class Part:
    values: dict[str, int]


@dataclass
class CategoryRange:
    start: int
    end: int


@dataclass
class PartRange:
    workflow: str
    categories: dict[str, CategoryRange]


@dataclass
class Check:
    input: str
    operator: str
    value: int
    workflow: str


@dataclass
class Workflow:
    name: str
    checks: list[Check]


def main(is_test_mode: bool) -> None:
    print('aplenty')
    print()

    (_, workflows) = get_input(is_test_mode)
    q: list[PartRange] = [PartRange(
        'in',
        {
            'x': CategoryRange(1, 4001),
            'm': CategoryRange(1, 4001),
            'a': CategoryRange(1, 4001),
            's': CategoryRange(1, 4001)
        })]

    approved: list[PartRange] = []
    rejected: list[PartRange] = []
    while len(q) > 0:
        current = q.pop(0)
        for check in workflows[current.workflow].checks:
            if check.input in current.categories:
                (pass_range, fail_range) = split_part_range(current, check)
                if fail_range is not None:
                    current = fail_range
                if pass_range is not None:
                    if check.workflow == 'A':
                        approved.append(pass_range)
                    elif check.workflow == 'R':
                        rejected.append(pass_range)
                    else:
                        pass_range.workflow = check.workflow
                        q.append(pass_range)
            else:
                if check.workflow == 'A':
                    approved.append(current)
                elif check.workflow == 'R':
                    rejected.append(current)
                else:
                    current.workflow = check.workflow
                    q.append(current)


    running_total = 0
    for part_range in approved:
        total = \
            (part_range.categories['x'].end - part_range.categories['x'].start) * \
            (part_range.categories['m'].end - part_range.categories['m'].start) * \
            (part_range.categories['a'].end - part_range.categories['a'].start) * \
            (part_range.categories['s'].end - part_range.categories['s'].start)
        running_total += total

        print(part_range.workflow, end=',')
        for cat in part_range.categories:
            print(f'{part_range.categories[cat].start},{part_range.categories[cat].end},', end='')
        print()


    print(running_total)


def split_part_range(part_range: PartRange, check: Check) -> tuple[PartRange, PartRange]:
    category_range = part_range.categories[check.input]
    if check.operator == '<':
        if category_range.end <= check.value:
            return (part_range, None)
        else:
            if category_range.start >= check.value:
                return (None, part_range)
            else:
                return (
                    part_range_with(part_range, check.input, category_range.start, check.value),
                    part_range_with(part_range, check.input, check.value, category_range.end))
    else:
        # >
        if part_range.categories[check.input].start > check.value:
            return (part_range, None)
        else:
            if category_range.end <= check.value + 1:
                return(None, part_range)
            else:
                return (
                    part_range_with(part_range, check.input, check.value + 1, category_range.end),
                    part_range_with(part_range, check.input, category_range.start, check.value + 1))


def part_range_with(part_range: PartRange, category: str, start: int, end: int) -> PartRange:
    result = PartRange(
        part_range.workflow,
        {
            'x': CategoryRange(part_range.categories['x'].start, part_range.categories['x'].end),
            'm': CategoryRange(part_range.categories['m'].start, part_range.categories['m'].end),
            'a': CategoryRange(part_range.categories['a'].start, part_range.categories['a'].end),
            's': CategoryRange(part_range.categories['s'].start, part_range.categories['s'].end)
        })
    result.categories[category].start = start
    result.categories[category].end = end
    return result


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
