
def readFile(input_file: str):
    raw_list = []
    with open(input_file) as f:
        raw_list = f.read().splitlines()
    return raw_list

def determine_joltage(input: list[int]):
    # -2 is because we cannot pick the very last number
    largest_number = max(input[:-1])
    # Python searches left-to-right naturally, yay
    largest_index = input.index(largest_number)
    second_largest_number = max(input[largest_index+1:])
    # print(largest_number, largest_index, second_largest_number)
    return largest_number*10 + second_largest_number

def solution_one(input):
    solution = 0
    for battery in input:
        solution += determine_joltage([int(i) for i in battery])
    return solution

def solution_two(input):
    solution = 0
    return solution


if __name__ == '__main__':
    # determine_joltage([3,4,1,4])
    # print(solution_one(readFile('testdata.txt')))
    print(solution_one(readFile('data.txt')))
    # print(solution_two(readFile('data.txt')))
