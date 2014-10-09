# https://projecteuler.net/problem=2

@cache = {}
@cache[0] = 0
@cache[1] = 1
@cache[2] = 2

def fib(n)
  return @cache[n] if @cache[n]
  @cache[n] = fib(n-1) + fib(n-2)
end

n = 4_000_000
sum = val = i = 0

while val < n
  val = fib(i)
  sum += val if val % 2 == 0 && val < n
  i += 1
end

p sum
