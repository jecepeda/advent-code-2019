from typing import Dict, List, Tuple


class StepOperation:
    def __init__(self, op, steps):
        self.op = op
        self.steps = steps

    def __repr__(self):
        return "op: {} steps: {}".format(self.op, self.steps)


def get_op(elem):
    return StepOperation(op=elem[0], steps=int(elem[1:]))


def traverse_path(path: List[StepOperation]):
    result: Dict[Tuple[int, int], bool] = {}
    result_steps: Dict[Tuple[int, int], int] = {}
    x, y, steps = 0, 0, 0
    for step in path:
        if step.op == "R":
            partial = {(x_inc, y): True for x_inc in range(x, x + step.steps + 1)}
            partial_steps = {
                (x_inc, y): steps + (x_inc - x)
                for x_inc in range(x, x + step.steps + 1)
            }
            result_steps.update(partial_steps)
            result.update(partial)
            x += step.steps
        elif step.op == "L":
            partial = {(x_inc, y): True for x_inc in range(x, x - step.steps - 1, -1)}
            partial_steps = {
                (x_inc, y): steps + (x - x_inc)
                for x_inc in range(x, x - step.steps - 1, -1)
            }
            result.update(partial)
            result_steps.update(partial_steps)
            x -= step.steps
        elif step.op == "U":
            partial = {(x, y_inc): True for y_inc in range(y, y + step.steps + 1)}
            partial_steps = {
                (x, y_inc): steps + (y_inc - y)
                for y_inc in range(y, y + step.steps + 1)
            }
            result.update(partial)
            result_steps.update(partial_steps)
            y += step.steps
        elif step.op == "D":
            partial = {(x, y_inc): True for y_inc in range(y, y - step.steps - 1, -1)}
            partial_steps = {
                (x, y_inc): steps + (y - y_inc)
                for y_inc in range(y, y - step.steps - 1, -1)
            }
            result.update(partial)
            result_steps.update(partial_steps)
            y -= step.steps
        steps += step.steps
    return result, result_steps


def get_manhattan_distances(elems):
    return [abs(elem[0]) + abs(elem[1]) for elem in elems]


def get_intersection_steps(steps_1, steps_2, matches):
    return [steps_1[elem] + steps_2[elem] for elem in matches]


if __name__ == "__main__":
    first_path = None
    second_path = None

    with open("input.txt") as f:
        first_path = [get_op(elem) for elem in f.readline().split(",")]
        second_path = [get_op(elem) for elem in f.readline().split(",")]
    positions_1, steps_1 = traverse_path(first_path)
    positions_2, steps_2 = traverse_path(second_path)
    matches = set(positions_1.keys()) & set(positions_2.keys())
    manhattan_distances = get_manhattan_distances(matches)
    lowest_steps = get_intersection_steps(steps_1, steps_2, matches)
    print(sorted(manhattan_distances)[1])
    print(sorted(lowest_steps)[1])

