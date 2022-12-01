# Started out via copy+paste the data set

# Read the file into a list
raw_list = []
with open('data.txt') as f:
    raw_list = f.readlines()

# Convert raw list into elf amounts
calorie_list = []
calorie_count = 0
for idx, i in enumerate(raw_list):
    # Find a line break aka empty space
    if i.strip() == '':
        calorie_list.append(calorie_count)
        calorie_count = 0
    else:
        calorie_count += int(i.strip())
# Answer to part one
print(max(calorie_list))

# Top 3 calorie counts
top_three = sorted(calorie_list, reverse=True)[:3]
total = 0
for idx, i in enumerate(top_three):
    total += i
print(total)





    