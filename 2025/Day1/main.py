def shift(input_number: int, movement_string: str):
    if movement_string.startswith('R'):
        output = input_number + int(movement_string[1:])
        if output >= 100:
            output %= 100
    elif movement_string.startswith('L'):
        output = input_number - int(movement_string[1:])
        if output < 0:
            output %= 100
    else:
        raise ValueError
    return output

def readFile(input_file: str):
    raw_list = []
    with open(input_file) as f:
        raw_list = f.read().splitlines()
    return raw_list

def rotation(input_list):
    dial_number = 50
    times_hit_0 = 0
    for i in input_list:
        dial_number = shift(dial_number, i)
        if dial_number == 0:
            times_hit_0 += 1
    return dial_number, times_hit_0

if __name__ == '__main__':
    print(readFile('testdata.txt'))
    print(rotation(readFile('testdata.txt')))
    # print(readFile('data.txt'))
    print(rotation(readFile('data.txt')))