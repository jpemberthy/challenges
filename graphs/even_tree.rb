#!/usr/bin/env ruby

# https://www.hackerrank.com/challenges/even-tree
# TEST: cat even_tree.txt | ./even_tree.rb

require 'set'

input = []
STDIN.each_line { |line| input << line }
nodes_count, edges_count = input.first.split(/\s+/)
nodes = input[1, input.size-1].map { |n| n.split(/\s+/) }

@graph = Hash.new { |hsh, key| hsh[key] = Set.new }
nodes.each { |n| @graph[n.last] << n.first }

def count_nodes(node)
  result = 1
  @graph[node].each { |n| result += count_nodes(n) }
  result
end

cuts = 0
@graph.keys.each do |node|
  cuts += 1 if count_nodes(node).even?
end

STDOUT.puts cuts - 1
