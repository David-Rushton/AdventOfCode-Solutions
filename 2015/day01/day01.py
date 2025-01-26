import sys
import os

def main():
    for arg in getInput():
        step = 0
        current = 0;
        enteredBasement = 0
        
        for char in arg:
            step += 1
            current += 1 if char == '(' else -1

            if current == -1 and enteredBasement == 0:
                enteredBasement = step

        preview = f'{arg[0:10]}...' if len(arg) > 10 else arg
        print(f'{preview} | Final floor {current} | First basement visit {enteredBasement}')

def getInput():
    for arg in sys.argv[1:]:
        input = ''
        if os.path.exists(arg):
            yield open(arg, 'r').read()
        else:
            yield arg

if __name__ == '__main__':
    main()