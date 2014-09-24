# Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes,
# write a method to rotate the image by 90 degrees Can you do this in place?

n = 3
matrix = Array.new(n) { Array.new(n) }

[ "X", "0", "|" ].each_with_index do |pixel, i|
  3.times { |j| matrix[i][j] = pixel }
end

def rotate(matrix)
  n = matrix.size
  layer = 0
  while layer < n/2
    first = layer
    last = n - 1 - layer
    i = first
    while i < last
      offset = i - first
      top = matrix[first][i]
      # left -> top
      matrix[first][i] = matrix[last-offset][first]
      # bottom -> left
      matrix[last-offset][first] = matrix[last][last-offset]
      # right -> bottom
      matrix[last][last-offset] = matrix[i][last]
      # top -> right
      matrix[i][last] = top # right <- saved top

      i += 1
    end

    layer += 1
  end

  matrix
end

p matrix

p rotate(matrix)
