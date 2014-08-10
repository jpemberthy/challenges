#!/usr/bin/ruby

# Converts a decimal to binary with the repeated division by 2 method.
# Can be refactored to convert from decimal to any base.
def binary(decimal)
  result = decimal / 2
  binary = (decimal % 2).to_s

  while result > 0
    binary << (result % 2).to_s
    result /= 2
  end

  binary.reverse
end

decimals = []
50.times { |x| decimals << rand(1000) }
decimals.each do |d|
  b = binary(d)
  puts "#{d}(10) -> #{b}(2)"
  return "ERROR!" if d.to_s(2) != b
end
