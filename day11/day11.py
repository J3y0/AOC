import math

class Monkey(object):
    def __init__(self, operation = lambda x: x, nb = 0, starting_items = [], div = 0) -> None:
        self.number = nb
        self.items = starting_items
        self.operation = operation
        self.div = div
        self.to_who = {"true": 0, "false": 0}
        self.inspected = 0

    def inspect1(self) -> int:
        worry_level = self.operation(self.items[0])
        self.inspected += 1
        return math.floor(worry_level/3)
    
    def inspect2(self) -> int:
        self.inspected += 1
        return self.operation(self.items[0])

    def throw(self, value, other_monkey) -> None:
        other_monkey.items.append(value)
        self.items.pop(0)
    
    def __str__(self) -> str: # Flemme de parse operation ?
        return f"Monkey {self.number}:\n" +\
        f"  Starting items: {self.items}\n" +\
        f"  Operation: new = \n" +\
        f"  Test: divisible by {self.div}\n" +\
        "    If true: throw to monkey {}\n".format(self.to_who["true"]) +\
        "    If false: throw to monkey {}".format(self.to_who["false"])


def parse_comma_values(line: str):
    return list(map(lambda x: int(x), line.split(", ")))

def parse_monkey(desc: str) -> Monkey:
    lines = desc.split("\n")
    m = Monkey()
    m.number = int(lines[0].split(":")[0])
    m.items = parse_comma_values(lines[1].split(": ")[1])
    eq = lines[2].split("= ")[1]
    value = eq.split(" ")
    symbol = value[1]
    if value[2] != "old":
        number = int(value[2])
        if symbol == "+":
            m.operation = lambda x: x + number
        elif symbol == "-":
            m.operation = lambda x: x - number
        else:
            m.operation = lambda x: x * number
    else:
        if symbol == "+":
            m.operation = lambda x: x + x
        elif symbol == "-":
            m.operation = lambda x: x - x
        else:
            m.operation = lambda x: x * x

    m.div = int(lines[3].split("by ")[1])
    m.to_who["true"] = int(lines[4].split("monkey ")[1])
    m.to_who["false"] = int(lines[5].split("monkey ")[1])

    return m

def compute(inspected):
    result = sorted(inspected)
    return result[-1] * result[-2]

if __name__ == "__main__":
    with open("./day11.txt") as f:
        content = f.read();
        data = content.split("Monkey ")
        data.pop(0) # Removing the first item as it is an empty string
        
        # Init monkeys
        # We are forced to modulo the worry values else it's too big for part 2
        # We use a modulo = all_div_value multiplied together
        # Like that, we are not affecting the divisors of the value
        mod = 1
        monkeys = []
        for m in data:
            monkeys.append(parse_monkey(m))
            mod *= monkeys[-1].div

        # Part 1
        # round = 20

        # Part 2
        round = 10000
        

        for i in range(round):
            for m in monkeys:
                while len(m.items) > 0:
                    worry_level = m.inspect2() % mod
                    if worry_level % m.div == 0:
                        m.throw(worry_level, monkeys[m.to_who["true"]])
                    else:
                        m.throw(worry_level, monkeys[m.to_who["false"]])
        
        inspected = []
        # Check
        for m in monkeys:
            inspected.append(m.inspected)
        
        print(inspected)
        print(compute(inspected))