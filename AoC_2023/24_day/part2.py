import sympy
import time

t_start = time.time()
hailstones = [tuple(map(int, line.replace("@", ",").split(","))) for line in open("./data/day24.txt")]

xr, yr, zr, vxr, vyr, vzr = sympy.symbols("xr, yr, zr, vxr, vyr, vzr", integer=True)

equations = []

# Only 6 equations are needed as we have 6 unknowns
for xh, yh, zh, vxh, vyh, vzh in hailstones[:3]:
    equations.append((xh - xr)*(vyr - vyh) - (vxr - vxh)*(yh - yr))
    equations.append((xh - xr)*(vzr - vzh) - (vxr - vxh)*(zh - zr))

answer = sympy.solve(equations)
assert len(answer) == 1
ans = answer[0]

print("Part 2:", ans[xr] + ans[yr] + ans[zr], time.time() - t_start)