# Write a method to decide if two strings are anagrams or not

def anagram?(str_1, str_2)
  str_1.each_char do |char|
    i = 0
    while i < str_2.size
      if str_2[i].downcase == char.downcase
        str_2[i] = ""
        i = str_2.size
      else
        i += 1
      end
    end
  end

  str_2.size == 0
end

p anagram?("foo", "oFox")
