require 'thwait'
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

# busy_work.call(1)

threads = 4.times.map do |id|
 Thread.new do
  busy_work.call(id)
 end
end

ThWait.all_waits(threads)
