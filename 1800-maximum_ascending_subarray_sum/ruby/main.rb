# @param {Integer[]} nums
# @return {Integer}
def max_ascending_sum(nums)
  max_sum = 0
  current_sum = nums[0]

  (1...nums.length).each do |i|
    if nums[i - 1] < nums[i]
      current_sum += nums[i]
    else
      max_sum = [max_sum, current_sum].max
      current_sum = nums[i]
    end
  end

  [max_sum, current_sum].max
end

if __FILE__ == $0
  nums = [10, 20, 30, 5, 10, 50]
  result = max_ascending_sum(nums)
  puts "Result: #{result}"

  if result == 65
    puts "Passed"
  else
    puts "Failed"
  end
end
