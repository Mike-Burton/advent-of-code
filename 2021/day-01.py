"""day-01."""

depths: list[int] = []
with open("inputs/day-01.txt") as f:
    for reading in f:
        depths.append(int(reading))
    f.close()

# Part 1
increased_count = 0
last_depth = depths[0]
for depth in depths[1:]:
    if last_depth < depth:
        increased_count += 1
    last_depth = depth

print(f"Part 1: {increased_count}")

# Part 2
increased_count = 0
rolling_sum = sum(depths[:3])
for i in range(3, len(depths)):
    next_sum = rolling_sum + depths[i] - depths[i - 3]
    if next_sum > rolling_sum:
        increased_count += 1
    rolling_sum = next_sum

print(f"Part 2: {increased_count}")
