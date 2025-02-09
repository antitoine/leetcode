# @param {Integer[]} nums
# @return {Integer}
def count_bad_pairs(nums)
  difference_counter = Hash.new(0)
  bad_pairs_count = 0
  nums.each_with_index do |number, index|
    diff = index - number
    bad_pairs_count += index - difference_counter[diff]
    difference_counter[diff] += 1
  end
  bad_pairs_count
end

if __FILE__ == $0
  nums = [4,1,3,3]
  result = count_bad_pairs(nums)
  puts "Result: #{result}"

  if result == 5
    puts "Passed"
  else
    puts "Failed"
  end
end

