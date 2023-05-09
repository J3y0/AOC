import string

test = "vJrwpWtwJgWrhcsFMMfFFhFp\n\
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL\n\
PmmdzqPrVvPwwTWBwg\n\
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn\n\
ttgJtRGJQctTZtZT\n\
CrZsJsPPZsGzwwsLwLmpwMDw"

# Part 1
def find_common_item(ruckstacks):
    stacks = ruckstacks.split("\n")
    commons = []
    for s in stacks:
        d = {}
        common = ""
        m = len(s)//2
        s1 = s[:m]
        s2 = s[m:]

        for c in s1:
            if c in d:
                d[c] += 1
            else:
                d[c] = 1
        
        for c in s2:
            if c in d and d[c] > 0:
                common += c
                d[c] -= 1
        common = set(common)
        for c in common:
            common = c
        commons.append(common)

    return commons

def find_index(letter):
    index = 0
    for i in range(len(string.ascii_letters)):
        if letter == string.ascii_letters[i]:
            index = i
    return index

def compute_priority(items):
    sum = 0
    for letter in items:
        priority = find_index(letter) + 1
        sum += priority

    return sum

# Part 2
def find_badge(ruckstacks):
    stacks = ruckstacks.split("\n")
    commons = []
    for i in range(0, len(stacks), 3):
        s1 = stacks[i]
        s2 = stacks[i+1]
        s3 = stacks[i+2]

        d = {}
        common12 = ""
        for c in s1:
            if c in d:
                d[c] += 1
            else:
                d[c] = 1
        
        for c in s2:
            if c in d and d[c] > 0:
                common12 += c
                d[c] -= 1
        
        common = ""
        for c in set(common12):
            common += c
        
        common12 = common
        common = ""
        
        for c in s3:
            if c in common12:
                common += c
        common = set(common)
        for c in common:
            common = c
        commons.append(common)

    return commons

with open("./day3.txt") as f:
    data = "".join(f.readlines())
    items = find_common_item(data)
    print("Sum priority = ", compute_priority(items))

    badges = find_badge(data)
    print(compute_priority(badges))