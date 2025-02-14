class ProductOfNumbers
  def initialize
    @products = {}
  end

=begin
    :type num: Integer
    :rtype: Void
=end
  def add(num)
    if num == 0
      @products = {}
    else
      @products.each do |key, value|
        @products[key] = value * num
      end
      @products[@products.length] = num
    end
  end

=begin
    :type k: Integer
    :rtype: Integer
=end
  def get_product(k)
    return 0 if k > @products.length
    @products[@products.length - k]
  end
end

if __FILE__ == $0
  obj = ProductOfNumbers.new

  obj.add(3)
  obj.add(0)
  obj.add(2)
  obj.add(5)
  obj.add(4)

  if obj.get_product(2) != 20
    puts "1. Test failed, expected 20"
    exit
  end

  if obj.get_product(3) != 40
    puts "2. Test failed, expected 40"
    exit
  end

  if obj.get_product(4) != 0
    puts "3. Test failed, expected 0"
    exit
  end

  obj.add(8)

  if obj.get_product(2) != 32
    puts "4. Test failed, expected 32"
    exit
  end

  puts "Passed"
end

