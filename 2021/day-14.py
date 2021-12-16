"""day-14."""

template: list[str] = []
polyize: dict[tuple[str, str], list[tuple[str, str]]] = {}
with open("inputs/day-14.txt") as f:
    i: int = 0
    for line in f:
        clean_line = line.strip()
        if i == 0:
            template = list(clean_line)
        elif i == 1:
            pass
        else:
            before, _, after = clean_line.split(" ")
            polyize[(before[0], before[1])] = [(before[0], after), (after, before[1])]
        i += 1
    f.close()

char_count: dict[str, int] = {}
pair_count: dict[tuple[str, str], int] = {}

# Parse template
for i in range(1, len(template)):
    pair = (template[i-1], template[i])
    if pair in pair_count:
        pair_count[pair] += 1
    else:
        pair_count[pair] = 1

for char in template:
    if char in char_count:
        char_count[char] += 1
    else:
        char_count[char] = 1

for i in range(1, 41):
    new_pair_count = dict()
    for pair, count in pair_count.items():
        pair_a, pair_b = polyize[pair]
        if pair_a in new_pair_count:
            new_pair_count[pair_a] += count
        else:
            new_pair_count[pair_a] = count

        if pair_b in new_pair_count:
            new_pair_count[pair_b] += count
        else:
            new_pair_count[pair_b] = count

        new_char = pair_a[1]
        if new_char in char_count:
            char_count[new_char] += count
        else:
            char_count[new_char] = count

    pair_count = new_pair_count.copy()

    if i % 10 == 0:
        min_char = min(char_count.keys(), key=(lambda key: char_count[key]))
        max_char = max(char_count.keys(), key=(lambda key: char_count[key]))
        print(f"Step {i}: {char_count[max_char] - char_count[min_char]}")

# Step 10: 2194
# Step 20: 2264701
# Step 30: 2305686490
# Step 40: 2360298895777
