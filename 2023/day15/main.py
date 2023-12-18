def parse_input():
    return open('input.txt').read().split(',')

def hash_algo(lens: str) -> int:
    curr_value = 0
    for s in lens:
        ascii_val = ord(s)
        curr_value += ascii_val
        curr_value *= 17
        remainder = curr_value % 256
        curr_value = remainder
    return curr_value

def hashmapping(lens: str, mapp: dict) -> dict:
    if lens.find('-') != -1:
        symbol = '-'
    elif lens.find('=') != -1:
        symbol = '='

    split = lens.split(symbol)
    key = split[0]
    val = split[1]

    # Run the hash algo on the key
    key_hash = hash_algo(key)
    if symbol == '-':
        lenses = mapp.get(key_hash)
        if lenses is not None and len(lenses) > 0:
            # Remove the lens
            lenses.pop(key, None)
            # Readd the lenses to the mapp
            mapp[key_hash] = lenses
    elif symbol == '=':
        lenses = mapp.get(key_hash)
        if lenses is not None and len(lenses) > 0:
            # Overwrite the value
            lenses[key] = val
            # Readd the lenses to the mapp
            mapp[key_hash] = lenses
        else:
            # Create a new lens
            lens = {}
            lens[key] = val
            # Add the lens to the mapp
            mapp[key_hash] = lens

    return mapp

def part2():
    input = parse_input()
    mapp = {}
    sum = 0
    for s in input:
        mapp = hashmapping(s, mapp)
    for box_no, lenses in mapp.items():
            for key, focal_len in lenses.items():
                sum += (box_no + 1) * (list(lenses.keys()).index(key) + 1) * int(focal_len)

    return sum

def part1():
    input = parse_input()
    return sum(hash_algo(s) for s in input)
print('Part 1: ' + str(part1()))
print('Part 2: ' + str(part2()))