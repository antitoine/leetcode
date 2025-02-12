def maximum_sum(nums)
  max_sum = -1
  digit_sum_max_num = Hash.new(0)
  nums.each do |num|
    digit_sum = 0
    temp_num = num
    while temp_num > 0
      digit_sum += temp_num % 10
      temp_num /= 10
    end
    if digit_sum_max_num.key?(digit_sum)
      max_sum = [max_sum, digit_sum_max_num[digit_sum] + num].max
    end
    digit_sum_max_num[digit_sum] = [digit_sum_max_num[digit_sum], num].max
  end
  max_sum
end

if __FILE__ == $0
  nums = [18,43,36,13,7]
  result = maximum_sum(nums)
  puts "Result: #{result}"

  if result == 54
    puts "Passed"
  else
    puts "Failed"
  end
end

