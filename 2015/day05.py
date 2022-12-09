import sys

def starOne(path):
    words = open(path, 'r').read().splitlines()
    naughty = 0
    nice = 0

    for word in words:

        isNice = False

        vowelCount = 0
        if 'a' in word:
            vowelCount += word.count('a')
        if 'e' in word:
            vowelCount += word.count('e')
        if 'i' in word:
            vowelCount += word.count('i')
        if 'o' in word:
            vowelCount += word.count('o')
        if 'u' in word:
            vowelCount += word.count('u')

        if vowelCount < 3:
            print(f'\t{word} naughty 👿')
            naughty += 1
            continue

        naughtyCombinations = ['ab', 'cd', 'pq', 'xy']
        isNaughty = False
        for naughtyCombination in naughtyCombinations:
            if naughtyCombination in word:
                print(f'\t{word} naughty 👿')
                isNaughty = True
                naughty += 1
                break

        if isNaughty:
            continue

        for i in range(1, len(word)):
            if word[i] == word[i - 1]:
                nice += 1
                isNice = True
                print(f'\t{word} nice 👼')
                break

        if not isNice:
            print(f'\t{word} naughty 👿')
            naughty += 1

    print(f'| naughty 👿 | {naughty} | nice 👼 | {nice} |')

def starTwo(path):
    words = open(path, 'r').read().splitlines()
    naughty = 0
    nice = 0

    for word in words:

        # pair of two and non-overlapping
        containsPair = False
        for i in range(len(word) - 1):
            pair = word[i:i+2]
            if word.count(pair) >= 2:
                containsPair = True
                break

        # repeated letter separated by another
        containsRepeated = False
        for i in range(0, len(word) - 2):
            if word[i] == word[i + 2]:
                repeated = word[i:i +3]
                containsRepeated = True
                break

        if containsPair and containsRepeated:
            nice += 1
            print(f'\t{word} nice 👼 | pair | {pair} | repeated | {repeated}')
        else:
            naughty += 1
            print(f'\t{word} naughty 👿')

    print(f'| naughty 👿 | {naughty} | nice 👼 | {nice} |')



if __name__ == '__main__':
    path = 'day05.input.txt.test' if sys.argv[1] == 'test' else 'day05.input.txt'
    isStarTwo = sys.argv[2] == 'star2'

    if isStarTwo:
        starTwo(path)
    else:
        starOne(path)

