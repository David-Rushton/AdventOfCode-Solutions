import os
import re
import sys

def main(isStarTwo: bool, streamPath: str):
    streams = getInput(streamPath)

    for stream in streams:
        recentCharacters = ['', '', '', '']
        currentCharacter = 0

        for character in stream:
            recentCharacters[currentCharacter % 4] = character

            if(isStartOfPacket(recentCharacters)):
                print (f'Start of packet <{stream[0:5]}...> found: {currentCharacter + 1}')
                break

            currentCharacter += 1


def isStartOfPacket(recentCharacters: list[str]):
    if recentCharacters[3] == '':
        return False
    uniqueCharacters = {recentCharacters[0], recentCharacters[1], recentCharacters[2], recentCharacters[3]}
    return len(uniqueCharacters) == 4

def getInput(path: str):
    for line in open(path, 'r').readlines():
        content = line.replace('\n' ,'')
        if len(content) > 0:
            yield content

if __name__ == '__main__':
    isTest = sys.argv[1] in ['true', 'test', 'yes', 'on']
    isStarTwo = sys.argv[2] in ['star2', 'star-2', 'startwo', 'star-two']
    streamPath = 'day06.input.txt.test' if isTest else 'day06.input.txt'
    main(isStarTwo, streamPath)
