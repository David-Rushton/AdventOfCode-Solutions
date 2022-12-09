import hashlib


def main():
    key = 'yzbqklnj'
    number = 0

    while True:
        candidate = f'{key}{number}'.encode('utf-8')
        result = hashlib.md5(candidate).hexdigest()
        print(f'\r{candidate} = {result}', end = '')

        if result.startswith("000000"):
            break

        number += 1

if __name__ == '__main__':
    main()

