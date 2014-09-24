# Design an algorithm and write code to remove the duplicate characters in a string without using any additional buffer.
# NOTE: One or two additional variables are fine. An extra copy of the array is not.

def remove_duplicates(str)
  i = 0
  while char = str[i]
    j = i + 1
    while str[j]
      str[j] = "" if str[j] == char
      j += 1
    end

    i += 1
  end

  str
end

p remove_duplicates("foobarbas")
