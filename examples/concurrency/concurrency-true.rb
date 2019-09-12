require 'benchmark'

busy_work = Proc.new do |id|
  time = Benchmark.measure do
    300.times do |i|
      300.times do |j|
        i * j

        File.open("/dev/null", "w") { |file|
          file.puts mul
        }
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
