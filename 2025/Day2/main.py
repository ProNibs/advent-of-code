
def readFile(input_file: str):
    raw_list = []
    with open(input_file) as f:
        raw_list = f.read().split(',')
    return raw_list

def check_for_repeats(input: str):
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

def solution_one(input):
    solution = 0
    for i in input:
        repeat_numbers = check_for_repeats(i)
        for j in repeat_numbers: 
            solution += j
    return solution

if __name__ == '__main__':
    check_for_repeats('11-22')
    # print(type(readFile('testdata.txt')))
    # print(solution_one(readFile('testdata.txt')))
    print(solution_one(readFile('data.txt')))
