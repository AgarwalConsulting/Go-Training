run-c:
	gcc concurrency.c
	./a.out

run-js:
	node concurrency.js

run-ruby:
	ruby concurrency.rb

run-go:
	go run concurrency.go

run-go-preemption:
	> /tmp/mul.txt
	go run sched_preemption_example.go &
	tail -f /tmp/mul.txt
