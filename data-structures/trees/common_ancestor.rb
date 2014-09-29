# Design an algorithm and write code to find the first common ancestor of two nodes in a binary tree.
# Avoid storing additional nodes in a data structure NOTE: This is not necessarily a binary search tree

@tree = Hash.new { |hsh, k| hsh[k] = Array.new }
@tree[0] = [1]
@tree[1] = [2, 3]
@tree[3] = [4, 5]
@tree[4] = [10, 15]
@tree[5] = [20, 30]


def fca(node_a, node_b)
  fixed_parent = parent(node_a)
  dynamic_parent = parent(node_b)

  return fixed_parent if fixed_parent == dynamic_parent

  while fixed_parent != 0
    while dynamic_parent != 0
      return dynamic_parent if dynamic_parent == fixed_parent
      dynamic_parent = parent(dynamic_parent)
    end

    dynamic_parent = parent(node_b)
    fixed_parent = parent(fixed_parent)
  end

  fixed_parent
end

def parent(node)
  # Could build a struct to hold the :parent value.
  # For simplicity just finding it here on the fly. O(n)
  @tree.detect { |k, nodes| nodes.include?(node) }.first
end

puts fca(5, 30)
puts fca(10, 15)
puts fca(2, 30)
puts fca(10, 3)
puts fca(20, 15)
