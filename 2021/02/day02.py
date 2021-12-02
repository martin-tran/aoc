pos1, depth1, pos2, depth2, aim = 0, 0, 0, 0, 0
while True:
    try:
        direction, magnitude = input().split()
        magnitude = int(magnitude)
        if direction[0] == 'f':
            pos1 += magnitude
            pos2 += magnitude
            depth2 += aim * magnitude
        elif direction[0] == 'd':
            depth1 += magnitude
            aim += magnitude
        else:
            depth1 -= magnitude
            aim -= magnitude
    except EOFError:
        break

print('Part 1:', pos1 * depth1)
print('Part 2:', pos2 * depth2)
