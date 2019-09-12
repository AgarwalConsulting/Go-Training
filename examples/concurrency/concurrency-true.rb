require 'benchmark'

busy_work = Proc.new do |id|
  time = Benchmark.measure do
    1000.times do |i|
    1000.times do |j|
      i * j
      end
    end
  end
  puts "#{id} Took: #{time}"
end

4.times do |id|
  pid = Process.fork do
    busy_work.call(id)
  end
  puts pid
end

Process.wait
