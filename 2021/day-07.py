"""day-07."""

crabs: list[int]
with open("inputs/day-07.txt") as f:
    crabs = [int(x) for x in f.read().split(",")]


crabs.sort()
# Part 1
median = crabs[len(crabs)//2]
fuel_cost = 0
for position in crabs:
    fuel_cost += abs(median - position)

print(f"Part 1: {fuel_cost}")
# 336120

# Part 2
least_fuel = -1
for possible_position in range(crabs[0], crabs[-1] + 1):
    fuel_cost = 0
    for crab_position in crabs:
        delta = abs(possible_position - crab_position)
        fuel_cost += delta * (delta + 1) // 2
    if least_fuel > fuel_cost or least_fuel == -1:
        least_fuel = fuel_cost


print(f"Part 2: {least_fuel}")
# 96864235
