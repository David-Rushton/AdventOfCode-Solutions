import sys
import os

def main():
    totalRequired = 0
    totalRibbon = 0
    for arg in getInput():
        elements = arg.split('x')
        length = int(elements[0])
        width = int(elements[1])
        height = int(elements[2])

        facesAB = 2 * length * width
        facesCD = 2 * width * height
        facesEF = 2 * height * length

        excess = min([facesAB, facesCD, facesEF]) / 2

        ribbon = (sum(sorted([length, width, height])[0:2]) * 2) + (length * width * height)
        totalRibbon += ribbon

        required = facesAB + facesCD + facesEF + excess
        totalRequired += required
        print(f'{arg} requires {required} square feet of wrapping paper and {ribbon} feet of ribbon')

    print(f'\nTotal wrapping paper required {totalRequired} square feet\nTotal ribbon required {totalRibbon}')

def getInput():
    for arg in sys.argv[1:]:
        input = ''
        if os.path.exists(arg):
            for line in open(arg, 'r').readlines():
                if len(line) > 1:
                    yield line.replace('\n' ,'')
        else:
            yield arg

if __name__ == '__main__':
    main()
