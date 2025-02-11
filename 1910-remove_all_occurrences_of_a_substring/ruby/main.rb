# @param {String} s
# @param {String} part
# @return {String}
def remove_occurrences(s, part)
  while s.include?(part)
    index = s.index(part)
    s = s[0, index] + s[index + part.length, s.length]
  end
  return s
end

if __FILE__ == $0
  s = "daabcbaabcbc"
  part = "abc"
  result = remove_occurrences(s, part)
  puts "Result: #{result}"

  if result == "dab"
    puts "Passed"
  else
    puts "Failed"
  end
end

