import itertools

inputs = []

# read inputs
with open("example", "r") as f:
  for line in f:
    split_line = line.strip().split(" ")
    val = split_line[0].replace(":","")
    inputs.append({'val':val, 'numbers':split_line[1:]})

# search possibilities
for input in inputs:
  numbs = input['numbers']
  # print(numbs)
  possible_ops = list(itertools.product(['+', '*'], repeat=len(numbs)-1))
  for possibility in possible_ops:
    possibilities = [m+n for m,n in zip(numbs,possibility)]
    possibilities.append(numbs[-1])
    print("".join(possibilities))