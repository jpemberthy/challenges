# Given a binary search tree, design an algorithm which creates a linked list of all the nodes at each depth
# (i e , if you have a tree with depth D, you’ll have D linked lists)


@tree = Hash.new { |hsh, k| hsh[k] = Array.new }
# Works for any binary tree. Not only BST.
@tree[0] = [1] # root
@tree[1] = [2, 3]
@tree[3] = [4, 5]
@tree[4] = [10, 15]
@tree[5] = [20, 30]

@lists = Array.new

def build_lists(node, li=0)
  @lists[li] ||= Array.new
  @tree[node].each do |x|
    @lists[li] << x
    build_lists(x, li+1)
  end
end

build_lists 0

p @lists
