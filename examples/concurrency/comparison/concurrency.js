function busyWork(id) {
  var iterCount = 30000;
  for(var i = 0; i < iterCount; i++) {
    for(var j = 0; j < iterCount; j++) {
      mul = i * j;
    }
  }
  console.log("Completed work: ", id);
}

// busyWork(1);
// busyWork(2);
// busyWork(3);

var promise = new Promise(function(resolve, reject) {
  busyWork(1);
  console.log("1. Hello from promise!");

  resolve(42);
});

promise.then((val) => {
  busyWork(2);
  console.log("2. Processed: ", val);
});

busyWork(3);
console.log("3. Hello, world!");
