import os
import re
import sys

def main(isStarTwo: bool, stacks: list[list[str]], instructionsPath: str):
    iteration = 1
    instructions = getInput(instructionsPath)

    for instruction in instructions:
        steps = re.findall('\d+', instruction)
        move = int(steps[0])
        fromStack = int(steps[1]) - 1
        toStack = int(steps[2]) - 1

        while move > 0:
            if isStarTwo:
                item = stacks[fromStack].pop()
                stacks[toStack].append(item)
                move -= 1
            else:
                item = stacks[fromStack].pop(move * -1)
                stacks[toStack].append(item)
                move -= 1

        printStacks(iteration, instruction, stacks)
        iteration += 1

    result = ""
    for stack in stacks:
        result += stack.pop()

    print(f'\nResult: {result}')


def printStacks(iteration: int, instruction: str, stacks: list[list[str]]):
    i = 1
    print(f'\n\titeration {iteration}: {instruction}')
    for stack in stacks:
        print(f'\t{i}: {stack}')
        i += 1

def getInput(path: str):
    for line in open(path, 'r').readlines():
        content = line.replace('\n' ,'')
        if len(content) > 0:
            yield content

def getTestStacks():
    return [
        ['Z', 'N'],
        ['M', 'C', 'D'],
        ['P']
    ]

def getStacks():
    return [
        ['D', 'H', 'N', 'Q', 'T', 'W', 'V', 'B'],
        ['D', 'W', 'B'],
        ['T', 'S', 'Q', 'W', 'J', 'C'],
        ['F', 'J', 'R', 'N', 'Z', 'T', 'P'],
        ['G', 'P', 'V', 'J', 'M', 'S', 'T'],
        ['B', 'W', 'F', 'T', 'N'],
        ['B', 'L', 'D', 'Q', 'F', 'H', 'V', 'N'],
        ['H', 'P', 'F', 'R'],
        ['Z', 'S', 'M', 'B', 'L', 'N', 'P', 'H'],
    ]


if __name__ == '__main__':
    isTest = sys.argv[1] in ['true', 'test', 'yes', 'on']
    isStarTwo = sys.argv[1] in ['star2', 'star-2', 'startwo', 'star-two']
    instructionsPath = 'day05.input.txt.test' if isTest else 'day05.input.txt'
    stacks = getTestStacks() if isTest else getStacks()
    main(isStarTwo, stacks, instructionsPath)
