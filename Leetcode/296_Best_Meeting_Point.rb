require 'bigdecimal'

Coordinate = Struct.new(:i, :j)

def distance(p1, p2)
  (p2.i - p1.i).abs + (p2.j - p1.j).abs
end

# @param {Integer[][]} grid
# @return {Integer}
def min_total_distance(grid)
  people_coordinates = []
  grid.each_index do |i|
    grid[i].each_index do |j|
     people_coordinates << Coordinate.new(i,j) if grid[i][j] == 1
    end
  end

  minimum_distance = BigDecimal::INFINITY
  # calculate minimum common distance
  grid.each_index do |i|
    grid[i].each_index do |j|
      c = Coordinate.new(i, j)
      d = 0
      for pc in people_coordinates
        d += distance(c, pc)
      end

      if d < minimum_distance
        minimum_distance = d
      end
    end
  end

  minimum_distance
end

grid = []
grid << [1, 0, 0, 0, 1]
grid << [0, 0, 0, 0, 0]
grid << [0, 0, 1, 0, 0]
puts min_total_distance(grid)
