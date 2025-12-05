
def readFile(input_file: str):
    raw_list = []
    with open(input_file) as f:
        raw_list = f.read().split(',')
    return raw_list

def check_for_repeats_part_one(input: str):
    output = []
    split_them = input.split('-')
    for i in range(int(split_them[0]), int(split_them[1])+1):
        # Let's assume a 50/50 down the middle is all that is needed
        string_version = str(i)
        midpoint = len(string_version)//2
        # print(string_version[:midpoint], string_version[midpoint:])
        if string_version[:midpoint] == string_version[midpoint:]:
            output.append(i)
    return output


def check_for_repeats_part_two(input: str):
    output = []
    split_them = input.split('-')
    for i in range(int(split_them[0]), int(split_them[1])+1):
        string_version = str(i)
        if len(string_version) % 2 == 0:
            chunked_array = chunking(string_version, 2)
            if check_chunked_arrays(chunked_array):
                output.append(i)
                continue
        # From here on out, prime numbers matter? We only get so long though
        if len(string_version) % 3 == 0:
            # print(string_version)
            chunked_array = chunking(string_version, 3)
            if check_chunked_arrays(chunked_array):
                output.append(i)
                continue
        if len(string_version) % 5 == 0:
            chunked_array = chunking(string_version, 5)
            # print('Chunk array', chunked_array)
            if check_chunked_arrays(chunked_array):
                output.append(i)
                continue
        if len(string_version) % 7 == 0:
            chunked_array = chunking(string_version, 7)
            if check_chunked_arrays(chunked_array):
                output.append(i)
                continue
        if len(string_version) % 11 == 0:
            chunked_array = chunking(string_version, 11)
            if check_chunked_arrays(chunked_array):
                output.append(i)
                continue
        if len(string_version) % 13 == 0:
            chunked_array = chunking(string_version, 13)
            if check_chunked_arrays(chunked_array):
                output.append(i)
                continue
        # # More than 17 digits is...extreme
        if len(string_version) % 17 == 0:
            chunked_array = chunking(string_version, 17)
            if check_chunked_arrays(chunked_array):
                output.append(i)
                continue
    return output

def chunking(input: list[str], chunk_number: int):
    chunk_size = len(input) // chunk_number
    chunked_list = []
    # print('cHUNK', chunk_size, input)
    for i in range(0,len(input), chunk_size):
        chunked_list.append(int(input[i: i+chunk_size]))
    return chunked_list

def check_chunked_arrays(input: list[int]):
    # Loop through all the arrays and return true if they all match
    for i in input:
        if input[0] != i:
            return False
    return True

def solution_one(input):
    solution = 0
    for i in input:
        repeat_numbers = check_for_repeats_part_one(i)
        for j in repeat_numbers: 
            solution += j
    return solution

def solution_two(input):
    solution = 0
    for i in input:
        repeat_numbers = check_for_repeats_part_two(i)
        for j in repeat_numbers: 
            solution += j
    return solution


if __name__ == '__main__':
    check_for_repeats_part_one('11-22')
    # print(type(readFile('testdata.txt')))
    # print(solution_one(readFile('testdata.txt')))
    # print(solution_one(readFile('data.txt')))
    print(solution_two(readFile('data.txt')))
    # print(check_for_repeats_part_two('2121212118-2121212124'))