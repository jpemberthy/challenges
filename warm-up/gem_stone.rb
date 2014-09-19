#!/usr/bin/env ruby

# https://www.hackerrank.com/challenges/gem-stones
# TEST: cat gem_stone.txt | ./gem_stone.rb

require 'set'

rocks_count = STDIN.gets.to_i
rocks = STDIN.gets.chop.split(//).uniq.to_set

(rocks_count - 1).times do |gem_stone|
  rocks &= STDIN.gets.chop.split(//).to_set
end

STDOUT.puts rocks.size
