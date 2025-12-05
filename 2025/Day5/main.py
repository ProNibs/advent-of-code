
def readFile(input_file: str):
    raw_list = []
    with open(input_file) as f:
        raw_list = f.read().splitlines()
    return raw_list

def fresh_ingredients(input: list[str]):
    ingredient_list = []
    for one_list in input:
        # Stop at the break spot
        if one_list == '':
            break
        one_list_ranges = one_list.split('-')
        ingredient_list.append((int(one_list_ranges[0]),int(one_list_ranges[1])))
    return ingredient_list

def isFresh(fresh_list: list[(int,int)], input: int):
    fresh = False
    for item in fresh_list:
        if item[0] <= input <= item[1]:
            fresh = True
            break
    return fresh
    

def solution_one(input: list[str]):
    fresh_items = fresh_ingredients(input)
    unknown_ingredients_start = input.index('')
    solution = 0
    # The index is on the break spot specifically
    for item in input[unknown_ingredients_start+1:]:
        if isFresh(fresh_items, int(item)):
            solution += 1
    return solution

if __name__ == '__main__':
    test_file = readFile('testdata.txt')
    real_file = readFile('data.txt')
    print(solution_one(test_file))
    print(solution_one(real_file))

