#!/usr/bin/env ruby

# http://uva.onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=24&page=show_problem&problem=3873
# TEST: cat martians.txt | ./martians.rb

require 'set'
k = STDIN.gets.to_i

def dfs(graph, node, discovered = Set.new)
  discovered << node

  graph[node].each do |g|
    dfs(graph, g, discovered) unless discovered.include? g
  end

  discovered.size
end

k.times do |_case|
  graph = Hash.new { |hsh, key| hsh[key] = Set.new }
  weights = Hash.new { |hsh, key| hsh[key] = 0 }

  nodes_count = STDIN.gets.to_i

  nodes_count.times do
    node = STDIN.gets.split(/\s+/)
    graph[node.first] << node.last
  end

  graph.each do |node, _v|
    weights[node] = dfs(graph, node)
  end

  max_martians = weights.sort_by { |_k, v| -v }.first.last
  martian = weights.find_all { |k, v| v == max_martians }.sort_by { |x| x.first }.first.first
  STDOUT.puts "CASE #{_case+1}: #{martian}"
end
