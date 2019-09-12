var promise =  new Promise(function(resolve, reject) {
  console.log("Hello from promise!");

  resolve(42);
});

promise.then((val) => {
  console.log("Processed: ", val);
});

console.log("Hello, world!");
