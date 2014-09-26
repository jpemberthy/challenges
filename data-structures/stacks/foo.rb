# Imagine a (literal) stack of plates If the stack gets too high, it might topple There- fore,
# in real life, we would likely start a new stack when the previous stack exceeds some threshold.
# Implement a data structure SetOfStacks that mimics this SetOf- Stacks should be composed of several stacks,
# and should create a new stack once the previous one exceeds capacity.
# SetOfStacks push() and SetOfStacks pop() should behave identically to a single stack (that is, pop() should return the same values as it would if there were just a single stack)

class StackSet
  # threshold max number of elements/stack.
  # size the total number of elements (sum of all the elements in all the stacks)
  attr_reader :threshold, :size

  def initialize(threshold:)
    @threshold = threshold
    @stacks = Array.new
    @size = 0
  end

  def stack_index
    [ @size - 1, 0 ].max / threshold
  end

  def push(element)
    @size += 1
    @stacks[stack_index] ||= Array.new
    @stacks[stack_index].push(element)
  end

  def pop
    @stacks[stack_index].pop
    @size -= 1
  end
end

stack = StackSet.new(threshold: 2)
7.times { |i| stack.push(i) }
3.times { |i| stack.pop }

p stack
