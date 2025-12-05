
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
    
def unique_fresh(input: list[(int,int)]):
    output = []
    for next_item in input:
        # print(output)
        # Seed with first item in list
        if output == []:
            output.append(next_item)
            continue
        for index,previous_item in enumerate(output):
            # print(next_item, index, previous_item)
            # These account for overlap
            if (previous_item[0] <= next_item[0]) and (previous_item[1] >= next_item[1]):
                # print('Complete overlap')
                break    # Complete overlap, move on 
            elif (previous_item[0] <= next_item[0] <= previous_item[1]) and (next_item[1] > previous_item[1]):
                    # print('Left overlap')
                    output[index] = (previous_item[0], next_item[1])
                    break
            elif (previous_item[0] <= next_item[1] <= previous_item[1]) and (next_item[0] < previous_item[0]):
                    # print('Right overlap')
                    output[index] = (next_item[0], previous_item[1])
                    break
            # These account for adjacent (combine them)
            elif previous_item[0] == next_item[1]+1:
                # print('Adjacent left overlap')
                output[index] = (next_item[0], previous_item[1])
                break
            elif previous_item[1] == next_item[0]-1:
                # print('Adjacent right overlap')
                output[index] = (previous_item[0], next_item[1])
                break
        else: # Checked against everything already put in, so cannot be combined
            output.append(next_item)
    return output
    
def sorted(input: list[(int,int)]):
    input.sort()
    return input


def solution_one(input: list[str]):
    fresh_items = fresh_ingredients(input)
    unknown_ingredients_start = input.index('')
    solution = 0
    # The index is on the break spot specifically
    for item in input[unknown_ingredients_start+1:]:
        if isFresh(fresh_items, int(item)):
            solution += 1
    return solution

def solution_two(input: list[str]):
    fresh_items = fresh_ingredients(input)
    solution = 0
    # Already know I cannot create a jumbo array....
    unique_items = sorted(fresh_items)
    unique_items = unique_fresh(fresh_items)
    # print(unique_items)
    # Couple of passes is enough to combine some of them that got bridged later, right?
    unique_items = sorted(unique_items)
    unique_items = unique_fresh(unique_items)
    unique_items = sorted(unique_items)
    unique_items = unique_fresh(unique_items)
    unique_items = unique_fresh(unique_items)
    unique_items = unique_fresh(unique_items)
    unique_items = unique_fresh(unique_items)
    # print(unique_items)
    for one_item in unique_items:
        # Assuming right item is always higher than left; +1 is for inclusive numbers
        solution += one_item[1] - one_item[0] + 1
    return unique_items, solution

if __name__ == '__main__':
    test_file = readFile('testdata.txt')
    real_file = readFile('data.txt')
    # print(solution_one(test_file))
    # print(solution_one(real_file))
    # print(unique_fresh([(3, 5), (10, 18), (12, 20)]))
    print(solution_two(test_file))
    print(solution_two(real_file))

