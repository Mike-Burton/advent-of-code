"""day-04."""
import numpy as np

NUM_BOARDS = 100

numbers: list[int] = []
boards_np = np.zeros((NUM_BOARDS, 5, 5), dtype=int)
with open("inputs/day-04.txt") as f:
    i = -1
    for line in f:
        if line == "\n":
            continue
        clean_line = line.strip()
        if i == -1:
            number_line = clean_line.split(",")
            numbers = [int(x) for x in list(number_line)]
            i += 1
            continue
        board_line = list(filter(lambda x: x != "", clean_line.split(" ")))
        board_index = int(i / 5)
        row_index = i % 5
        for col_index, x in enumerate(board_line):
            boards_np[board_index][row_index][col_index] = int(x)
        i += 1
    f.close()

# Part 1
picked_numbers = set()


def check_board(board):
    for i in range(len(board)):
        # check row
        row = board[i, :]
        if len(picked_numbers.intersection(row)) == 5:
            return True

        # check col
        col = board[:, i]
        if len(picked_numbers.intersection(col)) == 5:
            return True
    return False


def unmarked_sum(board):
    board_set = set(board.flatten())
    diff = picked_numbers.intersection(board_set)
    left = board_set.difference(diff)
    return sum(left)


score = 0
for i in range(0, len(numbers)):
    picked_numbers.add(numbers[i])
    for board in boards_np:
        bingo = check_board(board)

        if bingo:
            score = numbers[i] * unmarked_sum(board)
            break
    if score != 0:
        break

print(f"Part 1: {score}")
# 8442

# part 2
winning_boards = [False for _ in range(NUM_BOARDS)]
win_count = 0
score = 0
for i in range(0, len(numbers)):
    picked_numbers.add(numbers[i])
    for j, board in enumerate(boards_np):
        if winning_boards[j]:
            continue
        bingo = check_board(board)

        if bingo:
            winning_boards[j] = bingo
            win_count += 1

        if win_count == NUM_BOARDS:
            score = numbers[i] * unmarked_sum(board)
            break

    if score != 0:
        break

print(f"Part 2: {score}")
# 4590
