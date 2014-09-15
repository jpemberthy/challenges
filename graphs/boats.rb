#!/usr/bin/env ruby

# http://uva.onlinejudge.org/external/3/383.html

# Sample Input

# 2                    An integer from 1 to 10 inclusive that represents the number of data sets in the input file. Each data set represents a new shipping configuration.
# 6  7  5              The first line of data in a data set will contain three integers, say M, N, and P: M is an integer from 1 to 30 inclusive representing the number of warehouses in the data set; N is an integer from 0 to M*(M-1)/2 inclusive that represents the number of legs between warehouses in the data set; P is an integer from 0 to 10 inclusive that represents the number of shipping requests for which cost information is required.
# AA CC QR FF DD AB    M two-letter code names for the M warehouses of the data sets. Only capital letters are used. A single blank separates code names.
# AA CC                N lines follow the line of code names, containing shipping leg information in the format: ``XX YY", with XX and YY being the codes for two distinct warehouses in the set that have a direct link (a shipping leg) between them. There will be a single blank between the warehouse codes.
# CC QR
# DD CC
# AA DD
# AA AB
# DD QR
# AB DD
# 5  AA AB             P lines of shipping requests, one request per line. Each shipping request will begin with an integer between 1 and 20 inclusive that represents the size of the shipment. The shipment size will be followed by a pair of code names in the format ``AA BB", with AA and BB being the code for two distinct warehouses in the set which represent the source and destination of the requested shipment.
# 14 DD CC
# 1  CC DD
# 2  AA FF
# 13 AB QR
# 3 0 1
# AA BB CC
# 5  AA CC

# Sample Output
#
# SHIPPING ROUTES OUTPUT
#
# DATA SET  1
#
# $500
# $1400
# $100
# NO SHIPMENT POSSIBLE
# $2600
#
# DATA SET  2
#
# NO SHIPMENT POSSIBLE
#
# END OF OUTPUT

# TEST: cat boats.txt | ./boats.rb

require 'set'

class Node
  attr_reader :code
  attr_accessor :connections

  def initialize(code)
    @code = code
    @connections = Set.new
  end

  def add_connection(connection)
    @connections << connection
  end
end


class Boat
  OO = 1_000_000
  attr_reader :graph, :requests, :data_set

  # Requests Array of strings with the expected requests format: [ "5 AA AB" ]
  def initialize(graph, requests, data_set=1)
    @graph = graph
    @requests = requests
    @data_set = data_set
  end


  def floyd_warshall
    nodes_count = graph.size
    dist = Array.new(nodes_count) { Array.new(nodes_count, OO) }

    graph.each do |node|
      v = graph.index { |n| n.code == node.code }
      dist[v][v] = 0
      node.connections.each do |connection|
        u = graph.index { |n| n.code == connection.code }
        dist[u][v] = 1
      end
    end

    nodes_count.times do |k|
      nodes_count.times do |i|
        nodes_count.times do |j|
          if dist[i][j] > dist[i][k] + dist[k][j]
            dist[i][j] = dist[i][k] + dist[k][j]
          end
        end
      end
    end

    dist
  end

  def output
    dist = floyd_warshall
    puts "SHIPPING ROUTES OUTPUT\n\n"
    puts "DATA SET #{data_set}\n\n" # Pass as parameter.
    requests.each do |request|
      size, origin_code, destination_code = request.split(/\s+/)
      i = graph.index { |n| n.code == origin_code }
      j = graph.index { |n| n.code == destination_code }


      if dist[i][j] == OO
        puts "NO SHIPMENT POSSIBLE"
      else
        cost = dist[i][j] * size.to_i * 100
        puts "$#{cost}"
      end
    end
  end
end

class InputReader
  attr_reader :vertices, :shipping_legs, :requests

  def initialize(vertices, shipping_legs, requests)
    @vertices = vertices
    @shipping_legs = shipping_legs
    @requests = requests
  end

  # Returns the connected graph from the input source.
  def graph
    nodes = vertices.each_with_object([]) do |code, ary|
      ary << Node.new(code)
    end

    shipping_legs.each do |leg|
      origin_code, destination_code = leg.split(/\s/)
      origin = nodes.detect { |n| n.code == origin_code }
      destination = nodes.detect { |n| n.code == destination_code }

      # All legs are bidirectional.
      origin.add_connection destination
      destination.add_connection origin
    end

    nodes
  end
end

input = []
$stdin.each_line { |line| input << line.chop }
i = 1
input.first.to_i.times do |data_set|
  nodes_count, shipping_legs_count, requests_count = input[i].split(/\s+/).map(&:to_i)
  vertices = input[i+1].split(/\s+/)
  shipping_legs = input[i+2, shipping_legs_count]
  requests = input[i+2+shipping_legs_count, requests_count]

  i = i + 2 + shipping_legs_count + requests_count

  r = InputReader.new(vertices, shipping_legs, requests)
  b = Boat.new(r.graph, r.requests, data_set+1)
  data_set += 1
  b.output
end
