#!/usr/bin/env ruby

# https://www.hackerrank.com/challenges/even-tree
# TEST: cat cavity_map.txt | ./cavity_map.rb
# Making it work with Ruby 1.9.3

require 'matrix'

n = STDIN.gets.to_i
rows = []
n.times do
  rows << STDIN.gets.gsub(/\s+/, "").split(//).map(&:to_i)
end

@matrix = Matrix.rows rows

def cavity?(i, j)
  return false if border?(i) || border?(j)
  depth = @matrix[i, j]
  (depth > @matrix[i, j+1] && depth > @matrix[i, j-1]) && (depth > @matrix[i+1, j] && depth > @matrix[i-1, j])
end

def border?(index)
  # hackerrank runs 1.9.3 wich defines `row_size` on Matrix instead of 2+ `row_count`.
  index == 0 || (index == @matrix.row_size - 1)
end

i = 0
while i < n
  j = 0
  row = ""
  while j < n
    row << (cavity?(i, j) ? "X" : @matrix[i,j].to_s)
    j += 1
  end
  STDOUT.puts row
  i += 1
end
