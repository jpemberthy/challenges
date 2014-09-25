# Write code to remove duplicates from an unsorted linked list FOLLOW UP
# How would you solve this problem if a temporary buffer is not allowed?

require_relative '../linked_list'


l = LinkedList.new
values = [ 4, 1, 3, 5, 1, 3, 9, 4, 5 ]

values.reverse.each do |value|
  l.insert_at_head(Node.new(data: value))
end

def remove_duplicates(l)
  current_node = l.head_node

  while current_node
    previous_node = current_node
    next_node = l.nodes[current_node.next]

    while next_node
      if next_node.data == current_node.data
        l.remove_after(previous_node)
        next_node = l.nodes[previous_node.next]
      else
        previous_node = next_node
        next_node = l.nodes[next_node.next]
      end
    end

    current_node = l.nodes[current_node.next]
  end

end

remove_duplicates(l)

l.each { |n| p n.data }
