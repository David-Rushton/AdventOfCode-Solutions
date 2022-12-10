import os
import re
import sys

def main(isStarTwo: bool, path: str):
    workingDirectory = []
    files = {}
    lines = getInput(path)

    for line in lines:
        if line == '$ cd /':
            workingDirectory = ['~']
            print(f'\tNew working directory {"/".join(workingDirectory)}')
            continue

        if line == '$ cd ..':
            workingDirectory.pop()
            print(f'\tNew working directory {"/".join(workingDirectory)}')
            continue

        if line.startswith('$ cd'):
            workingDirectory.append(line.replace('$ cd ', ''))
            print(f'\tNew working directory {"/".join(workingDirectory)}')
            continue

        if line == '$ ls':
            print(f'\tlisting contents')
            continue

        elements = line.split(' ')
        fileSize = elements[0]
        fileName = '/'.join(workingDirectory) + '/' + elements[1]

        if fileSize != 'dir':
            files[fileName] = int(fileSize)
            print(f'\tFound file {fileName} ({fileSize})')


    print('\nCalculating results\n')
    directories = {}
    for fileInfo in files:
        subDirectories = fileInfo.split('/')[0:-1]

        while len(subDirectories) > 0:
            subDirectory = "/".join(subDirectories)

            if subDirectory in directories:
                directories[subDirectory] += files[fileInfo]
            else:
                directories[subDirectory] = files[fileInfo]

            subDirectories.pop()

    freeSpace = 70000000 - directories["~"]
    requiredSpace = 30000000 - freeSpace
    smallestDirectoryName = "!"
    smallestDirectorySize = sys.maxsize

    result = 0
    for subDirectory in directories:
        print(f'\t{subDirectory} = {directories[subDirectory]}')

        if directories[subDirectory] <= 100000:
            result += directories[subDirectory]

        if directories[subDirectory] >= requiredSpace and directories[subDirectory] < smallestDirectorySize:
            smallestDirectoryName = subDirectory
            smallestDirectorySize = directories[subDirectory]


    print(f'\nTotal of directories <= 100,000 = {result}')
    print(f'Total free space = {freeSpace}')
    print(f'Total required space = {requiredSpace}')
    print(f'Smallest directory to delete {smallestDirectoryName} {smallestDirectorySize}')


def getInput(path: str):
    for line in open(path, 'r').readlines():
        content = line.replace('\n' ,'')
        if len(content) > 0:
            yield content

if __name__ == '__main__':
    isTest = sys.argv[1] in ['true', 'test', 'yes', 'on']
    isStarTwo = sys.argv[2] in ['star2', 'star-2', 'startwo', 'star-two']
    path = 'day07.input.txt.test' if isTest else 'day07.input.txt'
    main(isStarTwo, path)
