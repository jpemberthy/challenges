# Summation of primes
# https://projecteuler.net/problem=10
# Solved by generating all the primes below n with the sieve of Eratosthenes. http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
# takes less than 3 secs in Ruby.

n = 1_999_999

ary = Array.new(n) { true }

(2..Math.sqrt(n)).each do |i|
  if ary[i]
    j = i ** 2
    while j < n
      ary[j] = false
      j += i
    end
  end
end

primes = ary.each_with_index.inject(0) do |sum, (value, index)|
  sum += (value ? index : 0)
end

p primes - 1
