# Write a program to sort a stack in ascending order.
# You should not make any assumptions about how the stack is implemented.
# The following are the only functions that should be used to write this program: push | pop | peek | isEmpty

require 'delegate'

class Stack < SimpleDelegator
  def initialize
    @array = Array.new
    super(@array)
  end


  # Not using sort to avoid "overriding" Array#sort.
  def ssort
    temp = self.class.new

    while self.any?
      element = self.pop
      self.push(temp.pop) while temp.any? && temp.peek > element
      temp.push element
    end

    temp
  end

  def peek
    self.last
  end
end


s = Stack.new
10.times { s.push rand(100)  }
p s
p s.ssort
