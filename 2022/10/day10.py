import sys

def main(path: str):
    instructions = open(path, 'r').read().splitlines()
    nextInstruction = ''
    cpuCycle = 0
    value = 1
    signalStrength = []

    while len(instructions) > 0 or len(nextInstruction) > 0:

        cpuCycle += 1

        if len(nextInstruction) > 0:
            instruction = nextInstruction
            nextInstruction = ''
        else:
            instruction = instructions.pop(0)
            if instruction.startswith('addx'):
                nextInstruction = instruction
                instruction = 'pass'

        position = (cpuCycle - 1) % 40
        pixel = '#' if value in (position - 1, position, position + 1) else '.'
        delimiter = '\n' if position == 39 else ''
        print(pixel, end = delimiter)

        if cpuCycle in [20, 60, 100, 140, 180, 220]:
            strength = cpuCycle * value
            signalStrength.append(strength)

        if instruction == 'pass':
            pass

        if instruction == 'noop':
            pass

        if instruction.startswith('addx'):
            value += int(instruction.split(' ')[1])

    print(f'Register Value: {value}')
    print(f'Signal Strength: {sum(signalStrength)}')




if __name__ == '__main__':
    path = 'day10.input.txt.test' if sys.argv[1] == 'test' else 'day10.input.txt'
    main(path)
