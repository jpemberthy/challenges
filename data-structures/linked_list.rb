# Ruby implementation of a singly linked list.
require 'securerandom'

class LinkedList
  attr_accessor :nodes, :head

  def initialize
    @nodes = { }
  end

  def insert_after(node, new_node)
    new_node.next = node.next
    node.next = new_node.key
    nodes[new_node.key] = new_node
  end

  def insert_at_head(node)
    node.next = self.head
    self.head = node.key
    nodes[node.key] = node
  end

  def remove_after(node)
    if to_remove = nodes[node.next]
      node.next = to_remove.next
      nodes.delete(to_remove.key)
    end
  end

  def remove_head
    if old_head_node = nodes[self.head]
      self.head = old_head_node.next
    end

    nodes.delete(old_head_node.key)
  end

  def append_to_tail(node)
    each do |n|
      if n.next == nil
        n.next = node.key
        nodes[node.key] = node
        return # explicit but this n should be the last item always.
      end
    end
  end

  def each(&block)
    current_node = nodes[head]

    while current_node
      block.call(current_node)
      current_node = nodes[current_node.next]
    end
  end
end

class Node
  attr_accessor :next, :data
  attr_reader :key

  def initialize(data:)
    @data = data
    @key = SecureRandom.hex(10)
  end
end


l = LinkedList.new
world = Node.new(data: "world")
l.insert_at_head(world)
hello = Node.new(data: "Hello!")
l.insert_at_head(hello)
crazy = Node.new(data: "crazy")
l.insert_after(hello, crazy)
l.each { |n| puts n.data }

puts "\nremoving 'crazy'"
l.remove_after(hello)
l.each { |n| puts n.data }

puts
l.insert_at_head(Node.new(data: "OMG!"))
l.remove_head
l.each { |n| puts n.data }


puts
l.append_to_tail(Node.new(data: "and bye!"))
l.each { |n| puts n.data }
