# https://projecteuler.net/problem=1

multiples = (1..999).inject(0) do |sum, i|
  sum += ((i % 3 == 0) || (i % 5 == 0)  ? i : 0)
end

p multiples
