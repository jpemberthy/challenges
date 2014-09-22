# Summation of primes
# https://projecteuler.net/problem=10
# Solved by generating all the primes below n with the sieve of Eratosthenes. http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
# takes ~ 1 sec in Ruby.

n = 1_999_999

ary = Array.new(n) { true }

sum = 0
j = 0
sqrt_n = Math.sqrt(n)

(0..n).each do |i|
  sum += i if ary[i]

  if j > 1 && j <= sqrt_n
    if ary[j]
      k = j ** 2
      while k < n
        ary[k] = false
        k += j
      end
    end
  end

  j += 1
end

p sum - 1
