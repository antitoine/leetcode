# @param {Integer} limit
# @param {Integer[][]} queries
# @return {Integer[]}
def query_results(limit, queries)
  colors = Hash.new(0)
  balls = {}
  result = []

  queries.each_with_index do |query, i|
    query_ball = query[0]
    query_color = query[1]

    old_color = balls[query_ball]
    if old_color
      colors[old_color] -= 1
      colors.delete(old_color) if colors[old_color].zero?
    end

    colors[query_color] += 1
    balls[query_ball] = query_color

    result << colors.size
  end

  result
end

if __FILE__ == $0
  limit = 4
  queries = [[1, 4], [2, 5], [1, 3], [3, 4]]
  result = query_results(limit, queries)
  puts "Result: #{result}"

  if result == [1, 2, 2, 3]
    puts "Passed"
  else
    puts "Failed"
  end
end
