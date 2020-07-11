run-c:
	gcc comparison/concurrency.c
	./a.out

run-js:
	node comparison/concurrency.js

run-py:
	python comparison/concurrency.py

run-ruby:
	ruby comparison/concurrency.rb

run-go:
	go run concurrency.go

run-go-preemption:
	> /tmp/mul.txt
	go run goroutines/sched_preemption_example.go &
	tail -f /tmp/mul.txt
