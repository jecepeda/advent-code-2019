from copy import deepcopy
from part_one import execute_opcodes


def get_noun_verb_for(x, instructions):
    for i in range(0, 100):
        for j in range(0, 100):
            items = deepcopy(instructions)
            items[1] = i
            items[2] = j
            items = execute_opcodes(items)
            if items[0] == 19690720:
                return i, j
    return 0, 0


if __name__ == "__main__":
    instructions = None
    with open("input.txt") as f:
        line = f.readline()
        instructions = [int(x) for x in line.split(",")]
    noun, verb = get_noun_verb_for(19690720, instructions)
    print(noun, verb)
    print(100 * noun + verb)
