from collections import defaultdict

n = int(input())
grid = defaultdict(int, {(0, 0):1})
walk_len = 1
x, y = 1, 0

def sumNeighbours(x, y):
    total = grid[(x+1, y)]
    total += grid[(x+1, y+1)]
    total += grid[(x, y+1)]
    total += grid[(x-1, y+1)]
    total += grid[(x-1, y)]
    total += grid[(x-1, y-1)]
    total += grid[(x, y-1)]
    total += grid[(x+1, y-1)]
    print('{}, {}: {}'.format(x, y, total))
    if total > n:
        print('Value: ', total)
        raise SystemExit        
    return total

while True:
    for _ in range(walk_len):
        grid[(x, y)] = sumNeighbours(x, y)
        y += 1
    walk_len += 1
    for _ in range(walk_len):
        grid[(x, y)] = sumNeighbours(x, y)
        x -= 1
    for _ in range(walk_len):
        grid[(x, y)] = sumNeighbours(x, y)
        y -= 1
    walk_len += 1
    for _ in range(walk_len):
        grid[(x, y)] = sumNeighbours(x, y)
        x += 1
