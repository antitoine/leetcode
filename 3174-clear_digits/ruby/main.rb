# @param {String} s
# @return {String}
def clear_digits(s)
  result = ""
  to_remove = 0
  (s.length - 1).downto(0) do |i|
    if s[i] >= '0' && s[i] <= '9'
      to_remove += 1
    elsif to_remove == 0
      result = s[i] + result
    else
      to_remove -= 1
    end
  end
  result
end

if __FILE__ == $0
  s = "abc"
  result = clear_digits(s)
  puts "Result: #{result}"

  if result == "abc"
    puts "Passed"
  else
    puts "Failed"
  end
end

