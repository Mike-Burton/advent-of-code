
input = []
with open("input.txt") as f:
    for line in f:
        input.append(list(line)[:-1])

# part 1
DIGESTS = {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

sum = 0
part_num = 0
valid_part = False
for idx, line in enumerate(input):
    if valid_part:
        sum += part_num

    part_num = 0
    valid_part = False

    for jdx, c in enumerate(line):
        if c not in DIGESTS:
            if valid_part:
                sum += part_num

            part_num = 0
            valid_part = False


        if c in DIGESTS:
            for di in [-1, 0, 1]:
                for dj in [-1, 0, 1]:
                    if di == 0 and dj == 0:
                        continue

                    i = idx + di
                    j = jdx + dj

                    if i < 0 or i >= len(input):
                        continue

                    if j < 0 or j >= len(line):
                        continue

                    if input[i][j] not in DIGESTS and input[i][j] != ".":
                        valid_part = True

            part_num = (part_num * 10) + int(input[idx][jdx])

if valid_part:
    sum += part_num

print(f"part 1: {sum}")
# 553079


# part 2
DIGESTS = {"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

sum = 0
for idx, line in enumerate(input):
    for jdx, c in enumerate(line):
        if c != "*":
            continue

        touches = []
        for di in [-1, 0, 1]:
            for dj in [-1, 0, 1]:
                if di == 0 and dj == 0:
                    continue

                i = idx + di
                j = jdx + dj


                if i < 0 or i >= len(input):
                    continue

                if j < 0 or j >= len(line):
                    continue

                if input[i][j] not in DIGESTS:
                    continue

                touches.append((i, j))

        seen = set()
        for t in touches:
            if len(seen) > 2:
                continue

            i, j = t
            start = j
            end = j
            while start-1 >= 0 and input[i][start-1] in DIGESTS:
                start -= 1

            while end+1 < len(input[i]) and input[i][end+1] in DIGESTS:
                end += 1

            if (i, start, end) in seen:
                continue

            seen.add((i, start, end))

        if len(seen) != 2:
            continue

        ratio = 1
        for s in seen:
            part_num = 0
            i, start, end = s
            for j in range(start, end +1):
                part_num = (part_num * 10) + int(input[i][j])

            ratio *= part_num

        sum += ratio


print(f"part 2: {sum}")
# 84363105
