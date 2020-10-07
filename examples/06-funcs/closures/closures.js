function adder() {
  var sum
  sum = 0

  return (x) => {
    sum += x
    return sum
  }
}

function exec() {
  fn1 = adder()
  fn2 = adder()

  console.log(fn1(10))
  console.log(fn1(4))

  console.log(fn2(5))
}

exec()
