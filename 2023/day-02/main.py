import re

input = []
with open("input.txt") as f:
    for line in f:
        rounds = []
        round = line.split(";")
        for s in round:
            r = re.search(r"\d* red", s)
            g = re.search(r"\d* green", s)
            b = re.search(r"\d* blue", s)

            rounds.append(
                    (
                        int(r.group().split(" ")[0]) if r else 0,
                        int(g.group().split(" ")[0]) if g else 0,
                        int(b.group().split(" ")[0]) if b else 0,
                    )
                )

        input.append(rounds)

    f.close()


# part 1
MAX_RED = 12
MAX_GREEN = 13
MAX_BLUE = 14

sum = 0
for idx, game in enumerate(input):
    impossible = False
    for round in game:
        r, g, b = round
        if r > MAX_RED or g > MAX_GREEN or b > MAX_BLUE:
            impossible = True
            break


    if not impossible:
        sum += idx + 1

print(f"part 1: {sum}")
# 2207

# part 2
sum = 0
for idx, game in enumerate(input):
    max_red, max_green, max_blue = 1, 1, 1

    for round in game:
        r, g, b = round

        if max_red < r > 0:
            max_red = r

        if max_green < g > 0:
            max_green = g

        if max_blue < b > 0:
            max_blue = b

    sum += max_red * max_green * max_blue

print(f"part 2: {sum}")
# 62241
