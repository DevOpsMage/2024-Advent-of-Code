import sys

with open(sys.argv[1], 'r') as file:
    data = file.read().strip()

left = []
right = []

for line in data.split('\n'):
    nl, nr = line.split('   ')
    left.append(int(nl))
    right.append(int(nr))

sorted_left = sorted(left)
sorted_right = sorted(right)

p1 = [abs(nl - sorted_right[i]) for i, nl in enumerate(sorted_left)]
print(f"Part 1: {sum(p1)}")

c = {}
for n in right:
    c[n] = c.get(n, 0) + 1

p2 = [n * c.get(n, 0) for n in left]
print(f"Part 2: {sum(p2)}")