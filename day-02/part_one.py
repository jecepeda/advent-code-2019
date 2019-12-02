from copy import deepcopy


def execute_opcodes(items):
    for i in range(0, len(items), 4):
        op = items[i]
        i1, i2 = items[i + 1], items[i + 2]
        result = items[i + 3]
        if op == 1:
            items[result] = items[i1] + items[i2]
        elif op == 2:
            items[result] = items[i1] * items[i2]
        elif op == 99:
            break
    return items


if __name__ == "__main__":
    original_instructions = None
    with open("input.txt") as f:
        line = f.readline()
        original_instructions = [int(x) for x in line.split(",")]
    original_instructions[1] = 12
    original_instructions[2] = 2
    items = execute_opcodes(deepcopy(original_instructions))
    print(items[0])

