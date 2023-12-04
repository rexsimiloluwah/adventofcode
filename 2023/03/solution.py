"""Started late with low confidence. So I struggled with day 3.

So this is not original work. I only coded these for learning purposes.

The original logic behind this solution was found on Reddit: https://www.reddit.com/r/adventofcode/comments/189m3qw/2023_day_3_solutions/
"""
import math
import argparse
from typing import List
from collections import defaultdict


def read_input_schematic(filename: str) -> List[List[str]]:
    """Read the input schematic"""
    schematic = []

    with open(filename, "r") as f:
        lines = f.readlines()

    schematic = [list(line.strip()) for line in lines]
    return schematic


def extract_symbols_with_coords(schematic):
    return {
        (x, y): cell
        for y, row in enumerate(schematic)
        for x, cell in enumerate(row)
        if not cell.isdigit() and cell != "."
    }


def extract_symbols_with_adjacent_numbers_coords(schematic: List[List[str]]):
    """Extract every symbol and its adjacent numbers from the schematic"""
    # Found a better way to extract the symbols with its coordinates (Credit: https://www.reddit.com/r/adventofcode/comments/189m3qw/2023_day_3_solutions/)
    symbol_with_coords = extract_symbols_with_coords(schematic)

    # The eight possible directions for adjacency in the grid
    directions = [(-1, -1), (0, -1), (1, -1), (-1, 0), (1, 0), (-1, 1), (0, 1), (1, 1)]

    symbols_with_adjacent_numbers_coords = defaultdict(list)

    for x, y in symbol_with_coords.keys():
        symbols_with_adjacent_numbers_coords[(x, y)].extend(
            [
                (x + dx, y + dy)
                for dx, dy in directions
                if schematic[y + dy][x + dx] and schematic[y + dy][x + dx].isdigit()
            ]
        )

    return symbols_with_adjacent_numbers_coords


def part1(schematic):
    part_numbers = []
    visited = set()

    symbols_with_adjacent_numbers_coords = extract_symbols_with_adjacent_numbers_coords(
        schematic
    )

    for adjacent_number_coords in symbols_with_adjacent_numbers_coords.values():
        for x, y in adjacent_number_coords:
            if (x, y) in visited:
                continue

            start_ix, end_ix = x, x

            while start_ix >= 0 and schematic[y][start_ix].isdigit():
                visited.add((start_ix, y))
                start_ix += -1

            while end_ix < len(schematic) and schematic[y][end_ix].isdigit():
                visited.add((end_ix, y))
                end_ix += 1

            num = int("".join(schematic[y][start_ix + 1 : end_ix]))
            part_numbers.append(num)

    return sum(part_numbers)


def part2(schematic):
    visited = set()
    symbol_adjacent_gears = defaultdict(list)

    symbols_with_adjacent_numbers_coords = extract_symbols_with_adjacent_numbers_coords(
        schematic
    )

    for (
        sym_x,
        sym_y,
    ), adjacent_number_coords in symbols_with_adjacent_numbers_coords.items():
        for x, y in adjacent_number_coords:
            if (x, y) in visited:
                continue

            start_ix, end_ix = x, x

            while start_ix >= 0 and schematic[y][start_ix].isdigit():
                visited.add((start_ix, y))
                start_ix += -1

            while end_ix < len(schematic) and schematic[y][end_ix].isdigit():
                visited.add((end_ix, y))
                end_ix += 1

            num = int("".join(schematic[y][start_ix + 1 : end_ix]))
            symbol_adjacent_gears[(sym_x, sym_y)].append(num)

    symbols_with_coords = extract_symbols_with_coords(schematic)
    star_coords = dict(filter(lambda c: c[1] == "*", symbols_with_coords.items()))

    gears = [
        [g for g in symbol_adjacent_gears[(x, y)]]
        for x, y in star_coords
        if len(symbol_adjacent_gears[(x, y)]) == 2
    ]

    power_sum = sum(list(map(math.prod, gears)))

    return power_sum


if __name__ == "__main__":
    parser = argparse.ArgumentParser()

    parser.add_argument("--part", help="Part number i.e 1, 2", type=int, default=1)
    parser.add_argument(
        "--input_file",
        help="Path to input file",
        type=str,
        default="./input.txt",
    )

    args = parser.parse_args()

    schematic = read_input_schematic(args.input_file)

    if args.part == 1:
        print(f"Answer: {part1(schematic)}")
    elif args.part == 2:
        print(f"Answer: {part2(schematic)}")
    else:
        print("Part number can either be 1 or 2")
