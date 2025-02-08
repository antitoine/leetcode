class NumberContainers
  def initialize
    @by_idx = {}
    @by_val = Hash.new { |hash, key| hash[key] = [] }
  end

  def change(index, number)
    if @by_idx.key?(index)
      old_number = @by_idx[index]
      return if old_number == number

      @by_val[old_number].delete(index)
      @by_val.delete(old_number) if @by_val[old_number].empty?
    end

    @by_idx[index] = number

    idx = @by_val[number].bsearch_index { |x| x >= index } || @by_val[number].size
    @by_val[number].insert(idx, index)
  end

  def find(number)
    @by_val[number].first || -1
  end
end

if __FILE__ == $0
  obj = NumberContainers.new

  if obj.find(10) != -1
    puts "1. Test failed, expected -1"
    exit
  end

  obj.change(2, 10)
  obj.change(1, 10)
  obj.change(3, 10)
  obj.change(5, 10)

  if obj.find(10) != 1
    puts "2. Test failed, expected 1"
    exit
  end

  obj.change(1, 20)

  if obj.find(10) != 2
    puts "3. Test failed, expected 2"
    exit
  end

  puts "Passed"
end

