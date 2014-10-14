# https://projecteuler.net/problem=2

cache = {}
cache[0] = 0
cache[1] = 1
cache[2] = 2

def fib(n):
    if n in cache: return cache[n]
    cache[n] = fib(n-1) + fib(n-2)
    return cache[n]

n = 4000000
# there's a sum built in function.
_sum = val = i = 0

while val < n:
    val = fib(i)
    if val % 2 == 0 and val < n: _sum += val
    i += 1

print(_sum)
