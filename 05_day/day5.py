import re

def init_pile(data):
    start_pos = []
    instructions = []
    for elt in data:
        if 'move' in elt:
            instructions.append(elt)
        else:
            start_pos.append(elt)
    while '' in start_pos:
        start_pos.remove('')

    nb_piles = int(start_pos[-1].rstrip()[-1])
    start_pos.pop(-1)
    # Init piles
    piles = []
    for i in range(nb_piles):
        piles.append([])

    for line in range(len(start_pos)):
        start_pos[line] += " " # len = 36 now then divisible by 4
        for i in range(0,len(start_pos[line]),4):
            value = start_pos[line][i:i+4]
            value = re.findall(r'[\w]', value)
            
            if value != []:
                piles[i//4].append(value[0])
    
    for i in range(nb_piles):
        piles[i].reverse()
    return piles, instructions

# Part 1
def parse_instr(line):
    nb, start, end = re.findall(r'\d+', line)
    return int(nb), int(start) - 1, int(end) - 1

def push1(piles, start_pile, end_pile):
    piles[end_pile].append(piles[start_pile][-1])
    piles[start_pile].pop(-1)

def step1(piles, nb, start_pile, end_pile):
    for _ in range(nb):
        push1(piles, start_pile, end_pile)

# Part 2

def step2(piles, nb, start_pile, end_pile):
    sub = piles[start_pile][len(piles[start_pile]) - nb:]
    piles[end_pile].extend(sub)
    while nb > 0:
        piles[start_pile].pop(-1)
        nb -= 1

test = """
    [D]    \n
[N] [C]    \n
[Z] [M] [P]\n
 1   2   3 \n
\n
move 1 from 2 to 1\n
move 3 from 1 to 3\n
move 2 from 2 to 1\n
move 1 from 1 to 2
"""

with open("./day5.txt") as f:
    data = "".join(f.readlines()).split("\n")
    
    test = test.split("\n")
    piles, instructions = init_pile(data)

    for line in instructions:
        instr = parse_instr(line)
        # step1(piles, instr[0], instr[1], instr[2])
        step2(piles, instr[0], instr[1], instr[2])
    print(piles)

    # Print result
    for i in range(len(piles)):
        print(piles[i][-1], end='')
    print("\n")
