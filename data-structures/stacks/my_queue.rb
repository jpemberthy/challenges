# Implement a MyQueue class which implements a queue using two stacks

class MyQueue
  def initialize
    @stack_a = Array.new
    @stack_b = Array.new
  end

  def enqueue(element)
    @stack_a.push element
  end

  def dequeue
    origin = @stack_a
    destination = @stack_b

    destination.push(origin.pop) while !origin.empty?
    retval = destination.pop
    origin.push(destination.pop) while !destination.empty?
    retval
  end
end

queue = MyQueue.new
5.times { |i| queue.enqueue i }

p queue

queue.dequeue

p queue

queue.enqueue 5

p queue

queue.dequeue

p queue
