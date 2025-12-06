
def readFile(input_file: str):
    raw_list = []
    with open(input_file) as f:
        raw_list = f.read().splitlines()
    return raw_list

def determine_joltage(input: list[int]):
    # -1 is because we cannot pick the very last number
    largest_number = max(input[:-1])
    # Python searches left-to-right naturally, yay
    largest_index = input.index(largest_number)
    second_largest_number = max(input[largest_index+1:])
    # print(largest_number, largest_index, second_largest_number)
    return largest_number*10 + second_largest_number

def determine_joltage_part_two(input: list[int]):
    ordered_numbers = []
    # Seed initial number
    # -1 is because we cannot pick the very last number
    ordered_numbers.append(max(input[:-11]))
    # Python searches left-to-right naturally, yay
    previous_index = input.index(ordered_numbers[0])
    for i in range(11):
        print('Previous index', previous_index, 'Current iteration', i)
        # print(max(input[previous_index+1:-1*(11-i)]))
        if i == 10:
            # On last iteration, the 10-i turns into 0 which is not good
            ordered_numbers.append(max(input[previous_index+1:]))
        else:
            ordered_numbers.append(max(input[previous_index+1:-1*(10-i)]))
        previous_index = input.index(ordered_numbers[i+1], previous_index+1)
    print(ordered_numbers)
    return int("".join(str(num) for num in ordered_numbers))


def solution_one(input):
    solution = 0
    for battery in input:
        solution += determine_joltage([int(i) for i in battery])
    return solution

def solution_two(input):
    solution = 0
    for battery in input:
        solution += determine_joltage_part_two([int(i) for i in battery])
    return solution


if __name__ == '__main__':
    # determine_joltage([3,4,1,4])
    # print(solution_one(readFile('testdata.txt')))
    # print(solution_one(readFile('data.txt')))
    print(solution_two(readFile('data.txt')))
    # determine_joltage_part_two([9,8,7,6,5,4,3,2,1,9,1,1,1,1,9])
