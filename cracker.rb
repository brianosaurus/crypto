#!/usr/bin/env ruby

def convert_string_to_hex(str)
  str.scan(/./).map(&:ord).map { |x| x.to_s(16) }.join
end

def convert_to_hex_array(str)
  str.scan(/../).map(&:hex)
end

def xor_hex_arrs(arr1, arr2)
  if arr1.size < arr2.size
    arr1.each_with_index.map do |a, index|
      "%02x" % (a ^ arr2[index])
    end.join
  else
    arr2.each_with_index.map do |a, index|
      "%02x" % (a ^ arr1[index])
    end.join
  end
end

string_to_conv = ARGV[0]
hex_string = ARGV[1]

puts "xor'ed hex arrays are #{xor_hex_arrs(convert_to_hex_array(convert_string_to_hex(string_to_conv)), convert_to_hex_array(hex_string))}"
