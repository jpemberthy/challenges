require 'set'
require 'securerandom'
require 'thread'

# Trees provide an efficient insertion and searching
class BinaryTree
  attr_reader :root, :nodes

  def initialize(root_value:)
    @nodes = Hash.new
    @root = add_node(value: root_value)
  end

  # The height of a tree is a height of the root.
  def height
    root.height
  end

  # A full binary tree.is a binary tree in which each node has exactly zero or two children.
  def full?
  end

  # A complete binary tree is a binary tree, which is completely filled,
  # with the possible exception of the bottom level, which is filled from left to right.
  def complete?
  end

  # Find the next parent to make the tree complete. Uses BFS approach.
  def find_parent
    q = Queue.new
    discovered = Set.new
    discovered << root
    q.push root

    while !q.empty?
      t = q.pop
      return t unless t.children.size == 2
      t.each_child do |u|
        if !discovered.include?(u)
          discovered << u
          q.push u
        end
      end
    end
  end

  def insert(value)
    parent = find_parent
    child = add_node(value: value, parent: parent)
    parent.left.nil? ? parent.left = child : parent.right = child
    child
  end

  def add_node(args)
    n = Node.new(args.merge(tree: self))
    @nodes[n.reference] = n
    n
  end

  def render
    nodes.each do |_k, v|
      right = v.right ? v.right.value : ""
      left = v.left ? v.left.value : ""
      puts "#{v.value} -> (#{left} | #{right})"
    end
  end
end

class Node
  attr_reader :parent, :reference
  attr_accessor :left, :right, :value

  def initialize(value:, tree: nil, parent: nil, left: nil, right: nil)
    @value, @parent = value, parent
    @left, @right = left, right
    @reference = SecureRandom.hex(10)
  end

  # The depth of a node is the number of edges from the root to the node.
  def depth
  end

  # The height of a node is the number of edges from the node to the deepest leaf.
  def height
  end

  def children
    [ left, right ].compact
  end

  def each_child(&block)
    children.compact.each { |n| block.call(n) }
  end

end

b = BinaryTree.new(root_value: 1)
b.insert(5)
b.insert(7)
b.insert(8)
b.insert(9)
b.insert(10)
b.insert(11)

b.render
