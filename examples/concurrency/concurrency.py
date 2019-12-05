from threading import Thread
from random import randint
import time

class Say(Thread):
    def __init__(self, _id, message):
        super().__init__()
        self._id = _id
        self.message = message

    def busy_work(self):
        iter_count = 300 + randint(1, 100)

        for i in range(iter_count):
            for j in range(iter_count):
                mul = i * j
                with open("/dev/null", "w") as f:
                    f.write(str(mul))

    def run(self):
        self.busy_work()

        print(f"Hello, {self.message}! from {self._id} thread")


threads = [Say(i, "World!") for i in range(8)]

start = time.time()

for thread in threads:
    thread.start() # not run()

# wait for all 10 threads to complete
for thread in threads:
    thread.join()

print(f"Got {len(threads)} temps in {time.time() - start} seconds")
