"""day-02."""

DOWN = 0
UP = 1
FORWARD = 2
DIRECTION_MAP = {
    "down": DOWN,
    "up": UP,
    "forward": FORWARD,
}

instructions: list[tuple[str, int]] = []
with open("inputs/day-02.txt") as f:
    for line in f:
        direction, magnitude = line.split(" ")
        instructions.append((DIRECTION_MAP[direction], int(magnitude)))
    f.close()

# Part 1
depth = 0
horizontal = 0
for jump, how_high in instructions:
    if jump == DOWN:
        depth += how_high
    elif jump == UP:
        depth -= how_high
    else:
        horizontal += how_high


print(f"Part 1: {depth * horizontal}")
# 1728414

# Part 2
aim = 0
depth = 0
horizontal = 0
for jump, how_high in instructions:
    if jump == DOWN:
        aim += how_high
    elif jump == UP:
        aim -= how_high
    else:
        horizontal += how_high
        depth += aim * how_high

print(f"Part 2: {depth * horizontal}")
# 1765720035
