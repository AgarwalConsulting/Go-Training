# Basic

## 0. FizzBuzz challenge

Print numbers 1 to 100:

- for every number divisible by 3 print "Fizz" instead
- for every number divisible by 5 print "Buzz" instead
- for numbers divisible by 3 & 5 print "FizzBuzz"
- for other numbers print them as is

### Extra Challenge

- Create a function `fizzBuzz` which takes a `int` and returns a `string`
- Ensure that you define the logic in a separate package.

## 1. Input from user

Have the user enter a string, then loop through the string to generate a new string in which every character is duplicated, e.g., "hello" => "hheelllloo".

Test with "世界" as input.

## 2. Write a function to generate 'n' numbers from fibonacci series

- Print the first `n` numbers from the fibonacci series

## 3. Write a function to find factorials for a given n

- Print the factorial of number `n`

## 4. Prime Numbers

- Loop through the numbers from 2 to 25 and print out which numbers are prime, and for those numbers which are not prime numbers, you should print them as a product of two factors
- Remember that prime = no divisors other than 1 and itself
- Don't worry about efficiency, but if you're interested, check out math.Sqrt()

```
2 is prime
3 is prime
4 is product of 2 * 2
...
```

### Extra Challenge

- Create a function which can return the products given a number
  - Think about the return value in case the given number is prime?
- DRY-up your logic

## 5. Exercism - [Hamming](https://github.com/AgarwalConsulting/Go-Training/tree/master/exercises/exercism/hamming)

## 6. DS

Using maps and slices, create a text interface to allow a user to add employee names to a department in a company. For example, "Add Sally to Engineering" or "Add Amir to Sales." Then let the user retrieve a list of all people in a department or all people in the company by department, sorted alphabetically.
