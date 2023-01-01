

x = [1, 2, -3, 0, 3, 4, -2]


for i in range(5):
    y = x.copy()


    new_i = (5 + i) % 7

    y.insert(new_i, y.pop(5))

    print(f'== {i} | {new_i} ==')
    print(y)
    print()


