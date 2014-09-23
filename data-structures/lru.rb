# A simple Ruby LRU cache with get and set operations in constant time O(1).

# Capacity in terms of elements to make things simpler.
# lru = LRU.new(capacity: 10)

# lru.get("foo") -> nil
# lru.set("foo", "foo bar baz") -> true
# lru.get("foo")  ->  "foo bar baz"

# lru = LRU.new(capacity: 2)
# lru.set("ruby", "Ruby") -> true
# lru.set("go", "Go") -> true
# lru.set("python", "Python") -> true
#
# lru.get("ruby") -> nil
# lru.get("go") -> "GO"
# lru.get("pyton") -> "Python"
#
# lru.set("ruby", "ZOMG NEW RUBY") -> true
# lru.get("go") -> nil
# lru.get("pyton") -> "Python"

class LRU
  attr_reader :capacity, :cache

  def initialize(capacity:)
    @capacity = capacity
    @cache = Hash.new
    @list = Array.new(capacity) { nil }
  end

  def get(key)
    promote(key) if cache[key]
    cache[key]
  end

  def set(key, value)
    promote(key)
    prune if @list.size >= capacity
    cache[key] = value
  end

  private

  def promote(key)
    if index = @list.index(key)
      @list.insert(0, @list.delete_at(index))
    else
      @list.unshift(key)
    end
  end

  def prune
    cache.delete(@list.pop) while @list.size > capacity
  end
end

l = LRU.new(capacity: 1)
p l.get("foo")
l.set("foo", "foo bar baz")
p l.get("foo")

lru = LRU.new(capacity: 2)
lru.set("ruby", "Ruby")
lru.set("go", "Go")
lru.set("python", "Python")

p lru.get("ruby")
p lru.get("go")
p lru.get("python")

lru.set("ruby", "ZOMG NEW RUBY")
p lru.get("ruby")
p lru.get("go")
p lru.get("python")
