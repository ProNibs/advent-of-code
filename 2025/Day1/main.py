import math

def shift(input_number: int, movement_string: str):
    we_passed_zero = 0
    if movement_string.startswith('R'):
        output = input_number + int(movement_string[1:])
        if output >= 100:
            if output != 100:
                we_passed_zero += math.floor(output / 100)
            output %= 100
            # Dumb, I know
            if output == 0 and int(movement_string[1:]) > 100:
                we_passed_zero -= 1

    elif movement_string.startswith('L'):
        output = input_number - int(movement_string[1:])
        if output < 0:
            if input_number != 0:
                we_passed_zero += abs(math.floor(output / 100))
            output %= 100
    else:
        raise ValueError
    return output, we_passed_zero

def readFile(input_file: str):
    raw_list = []
    with open(input_file) as f:
        raw_list = f.read().splitlines()
    return raw_list

def rotation(input_list):
    dial_number = 50
    times_hit_0 = 0
    for i in input_list:
        times_hit_0 += shift(dial_number, i)[1]
        dial_number = shift(dial_number, i)[0]
        if dial_number == 0:
            times_hit_0 += 1
        print('Dial:', dial_number, '#', times_hit_0)
    return dial_number, times_hit_0



if __name__ == '__main__':
    # print(readFile('testdata.txt'))
    # print(rotation(readFile('testdata.txt')))
    # print(readFile('data.txt'))
    print(rotation(['L50', 'R1', 'L101']))
    # print(rotation(readFile('data.txt')))