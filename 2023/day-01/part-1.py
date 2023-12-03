import re

input = []
with open("input.txt") as f:
    for line in f:
        input.append(line)
    f.close()

# part 1
pattern = r"[A-Za-z\n]"
sum = 0
for line in input:
    l = list(re.sub(pattern, "", line))
    sum += (10 * int(l[0])) + int(l[-1])

print(f"part one: {sum}")


# part 2

# the replacement str has the word on both sides
# so if two digits share a char it will still work
# ex: eightwo
patterns = [
    (r"one", "one1one"),
    (r"two", "two2two"),
    (r"three", "three3three"),
    (r"four", "four4four"),
    (r"five", "five5five"),
    (r"six", "six6six"),
    (r"seven", "seven7seven"),
    (r"eight", "eight8eight"),
    (r"nine", "nine9nine"),
]
def str_to_int(in_str) -> str:
    for p in patterns:
        in_str = re.sub(p[0], p[1], in_str)

    return in_str

pattern = r"[A-Za-z\n]"
sum = 0
for line in input:  
    line = str_to_int(line)
    l = list(re.sub(pattern, "", line))
    sum += (10 * int(l[0])) + int(l[-1])

print(f"part two: {sum}")
