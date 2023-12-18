def is_reflection(grid, disance_to_match):
    for idx in range(len(grid)):
        if idx == 0:
            continue
        if sum(find_distance(l, r) for l, r in zip(reversed(grid[:idx]), grid[idx:])) == disance_to_match:
            return idx
    return 0

def score(lines, disance_to_match):
    grid = lines.split('\n')
    if row := is_reflection(grid, disance_to_match):
        return 100 * row
    if col := is_reflection(rotate(grid), disance_to_match):
        return col

def rotate(grid):
    return list(zip(*grid))

def find_distance(left, right):
    return sum(a != b for a, b in zip(left, right))


def part1():
    lines = open('input.txt').read().split('\n\n')
    return sum(score(grid, 0) for grid in lines)
def part2():
    lines = open('input.txt').read().split('\n\n')
    return sum(score(grid, 1) for grid in lines)
print(part1())
print(part2())