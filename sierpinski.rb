# _________1_________ # _________1_________
# ________111________ # ________111________
# _______11111_______ # _______1___1_______
# ______1111111______ # ______111_111______

# _____111111111_____ # _____1_______1_____
# ____11111111111____ # ____111_____111____
# ___1111111111111___ # ___1___1___1___1___
# __111111111111111__ # __111_111_111_111__

#!/usr/bin/ruby
# Ruby Algorithm to solve: https://www.hackerrank.com/challenges/functions-and-fractals-sierpinski-triangles
# Author: Juan Pemberthy. juan@pemberthy.com

class Sierpinski
  attr_accessor :n

  def initialize(n)
    @n = n
    # @m = Array.new(32, Array.new(63, "_"))
    @m = fill Array.new(32) { Array.new(63, "_") }
    puts "original"
    render(@m)
  end

  def sierpinskizy!
    base = Array.new(32) { Array.new(63, "_") }
    puts "reduced"
    reduced = reduce(@m)
    render reduce(@m)
    puts "transformed"
    transform(base, reduced)
    render base

  end

  def output
    sierpinskizy!
  end

  private

  def transform(base, reduced)
    triangle_size = reduced.last.length
    columns = reduced.length

    reduced.length.times do |level|
      onefy(base[level], level, base.first.length/2)
    end

    seeds = base[columns-1].each_index.select { |i| base[columns-1][i] == "1" }

    reduced.length.times do |level|
      onefy(base[columns + level], level, seeds.first - 1)
      onefy(base[columns + level], level, seeds.last + 1)
    end
  end

  def fill(m)
    mid = m.first.length / 2
    m.each_with_index { |row, level| onefy(row, level, mid) }
  end

  def onefy(row, level, seed_position)
    if level == 0
      row[seed_position] = "1"
    else
      row[seed_position-level] = row[seed_position+level] = "1"
      onefy(row, level - 1, seed_position)
    end
  end

  def reduce(m, n)
    return m if m.length == 1
    reduced = m.clone
    n.times do
      reduced = fill Array.new(reduced.length/2) { Array.new(reduced.first.length/2, "_") }
    end
    reduced
  end

  def render(m)
    m.each { |x| puts x.join("") }
  end
end

s = Sierpinski.new(1)
s.output
