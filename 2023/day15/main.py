def parse_input():
    return open('input.txt').read().split(',')

def do_steps(str):
    curr_value = 0
    for s in str:
        ascii_val = ord(s)
        curr_value += ascii_val
        curr_value *= 17
        remainder = curr_value % 256
        curr_value = remainder
    return curr_value


def part1():
    input = parse_input()
    return sum(do_steps(str) for str in input)
print(part1())