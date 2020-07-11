require 'thwait'
require 'benchmark'

busy_work = Proc.new do |id|
  iter_count = 300 + Random.rand(100)

  time = Benchmark.measure do
    iter_count.times do |i|
      iter_count.times do |j|
        mul = i * j

        File.open("/dev/null", "w") { |file|
          file.puts mul
        }
      end
    end
  end

  puts "Hello, world! from #{id} busy_work; Iter Count: #{iter_count}\n Took: #{time}"
end

# busy_work.call(1)

# -------- Threads ----------

threads = 4.times.map do |id|
 Thread.new do
  busy_work.call(id)
 end
end

ThWait.all_waits(threads)

# -------- Processes --------

# 4.times do |id|
#   pid = Process.fork do
#     busy_work.call(id)
#   end
#   puts pid
# end

# Process.wait
