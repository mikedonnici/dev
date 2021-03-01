pattern_unit = ".|."
pattern_filler = "-"
text = "WELCOME"
lines, width = input().split(" ")
lines = int(lines)
width = int(width)
centre_line_index = lines // 2
half_width = width // 2
welcome_mat = []
pattern_multiplier = 1
for idx in range(0, centre_line_index + 1):
    pattern = pattern_unit * pattern_multiplier
    if idx == centre_line_index:
        welcome_mat.append(f"{text:{pattern_filler}^{width}}")
        break
    welcome_mat.append(f"{pattern:{pattern_filler}^{width}}")
    pattern_multiplier += 2
welcome_mat = welcome_mat + welcome_mat[::-1][1: centre_line_index + 1]

for row in welcome_mat:
    print(row)
