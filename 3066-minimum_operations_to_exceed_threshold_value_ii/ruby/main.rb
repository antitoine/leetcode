require 'algorithms'

# @param {Integer[]} nums
# @param {Integer} k
# @return {Integer}
def min_operations(nums, k)
  heap = Containers::MinHeap.new
  nums.each { |num| heap.push(num) }
  operations_count = 0
  while heap.size > 1 && heap.min < k
    first_smallest = heap.pop
    second_smallest = heap.pop
    new_element = first_smallest * 2 + second_smallest
    heap.push(new_element)
    operations_count += 1
  end
  operations_count
end

if __FILE__ == $0
  nums = [2, 11, 10, 1, 3]
  k = 10
  result = min_operations(nums, k)
  puts "Result: #{result}"
  if result == 2
    puts "Passed"
  else
    puts "Failed"
  end
end

