# Time traveling hash.
# O(n) solution ...

class TTH
  attr_reader :_hash

  def initialize()
    @_hash = Hash.new
  end

  def put(time, key, value)
    time_hash = _hash[key] || Hash.new
    time_hash[time] = value
    _hash[key] = time_hash
  end

  def get(time, key)
    if value = _hash[key][time]
      return value
    end

    previous_time = 0
    _hash[key].each do |t|
      if t.first < time && t.first > previous_time
        previous_time = t.first
      end
    end

    if previous_time != 0
      return _hash[key][previous_time]
    end
  end
end

t = TTH.new
t.put(1, "foo", "bar")
t.put(1.5, "foo", "baz")
t.put(1.3, "foo", "holi")

t.put(5.0, "omg", "really")

puts t.get(1.1, "foo") == "bar"
puts t.get(1.2, "foo") == "bar"
puts t.get(1.4, "foo") == "holi"
puts t.get(1.5, "foo") == "baz"
puts t.get(2.0, "foo") == "baz"

puts t.get(1.0, "omg") == nil
puts t.get(5.0, "omg") == "really"
