def valid(line, groups):
    n = len(line)
    runs = []

    i = 0
    while i < n:
        while i < n and not line[i]:
            i += 1
        if i == n:
            break
        j = i
        c = 0
        while j < n and line[j]:
            j += 1
            c += 1
        runs.append(c)
        i = j
            
    return groups == runs

def permute(pattern, groups):
    line = []
    idxs = []

    for i, s in enumerate(pattern):
        if s == '.':
            line.append(0)
        if s == '?':
            line.append(-1)
            idxs.append(i)
        if s == '#':
            line.append(1)
    
    count = 0
    for mask in range(1 << len(idxs)):
        linecpy = line.copy()
        for i in range(len(idxs)):
            if mask & (1 << i):
                linecpy[idxs[i]] = 0
            else:
                linecpy[idxs[i]] = 1
        if valid(linecpy, groups):
            count += 1
    return count

def toInt(str):
    return int(str)

ans = 0
lines = open('input.txt').read().split('\n')
patterns = []
runs = []
for line in lines:
    split = line.split(' ')
    pattern = split[0]
    groups = split[1].split(',')
    patterns.append(pattern)
    nums = map(toInt, groups)
    runs.append(list(nums))
    
for i in range(len(patterns)):
    ans += permute(patterns[i], runs[i])
print(ans)