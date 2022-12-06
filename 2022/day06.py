import os
import re
import sys

def main(isStarTwo: bool, streamPath: str):
    streams = getInput(streamPath)

    for stream in streams:
        last14 = ['', '', '', '', '', '', '', '', '', '', '', '', '', '']
        last4 = ['', '', '', '']
        currentCharacter = 0
        packetFound = False
        messageFound = False

        for character in stream:
            last4[currentCharacter % 4] = character
            last14[currentCharacter % 14] = character

            if not packetFound and isStartOfPacket(last4):
                print (f'Start of packet <{stream[0:5]}...> found: {currentCharacter + 1}')
                packetFound = True

            if not messageFound and isStartOfMessage(last14):
                print (f'Start of message <{stream[0:5]}...> found: {currentCharacter + 1}')
                messageFound = True

            currentCharacter += 1


def isStartOfPacket(recentCharacters: list[str]):
    if recentCharacters[3] == '':
        return False

    uniqueCharacters = {recentCharacters[0], recentCharacters[1], recentCharacters[2], recentCharacters[3]}
    return len(uniqueCharacters) == 4

def isStartOfMessage(recentCharacters: list[str]):
    if recentCharacters[13] == '':
        return False

    uniqueCharacters = []
    for character in recentCharacters:
        if character in uniqueCharacters:
            return False
        else:
            uniqueCharacters.append(character)

    return True

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
