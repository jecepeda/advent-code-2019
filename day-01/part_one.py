def calc_fuel(x):
    return (x // 3) - 2


if __name__ == "__main__":
    with open("input.txt") as f:
        lines = (int(l) for l in f.readlines())
        result = sum(calc_fuel(x) for x in lines)
        print(result)
