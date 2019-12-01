def calc_fuel(x):
    return (x // 3) - 2


if __name__ == "__main__":
    with open("input.txt") as f:
        lines = (int(l) for l in f.readlines())
        result = 0
        for l in lines:
            fuel = calc_fuel(l)
            while fuel > 0:
                result += fuel
                fuel = calc_fuel(fuel)
        print(result)
