#!/usr/bin/env ruby

# https://www.hackerrank.com/challenges/runningtime
# TEST: cat running_time.txt | ./running_time.rb

# The size of the array is not really needed.
STDIN.gets
ary = STDIN.gets.split(/\s+/).map(&:to_i)

@shifts = 0

def shift!(ary, i, j)
  @shifts += 1
  x = ary[j]
  ary[j] = ary[i]
  ary[i] = x
  ary
end

i = 0
while i < ary.size do
  if ary[i+1] && ary[i] > ary[i+1]
    shift!(ary, i, i+1)
    j = i + 1
    i = 0
    while ary[j+1] && ary[j] > ary[j+1]
      shift!(ary, j, j+1)
      j += 1
    end
  else
    i += 1
  end
end

STDOUT.puts @shifts
