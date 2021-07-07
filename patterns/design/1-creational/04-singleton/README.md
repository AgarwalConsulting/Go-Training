# Singleton

Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.

> When discussing which patterns to drop, we found that we still love them all. (Not really -- I'm in favor of dropping Singleton. It's use is almost always a design smell.) - Elrich Gamma

The Singleton pattern solves two problems at the same time, violating the Single Responsibility Principle:

- Ensure that a class has just a single instance.
- Provide a global access point to that instance.

## Problem

- Ensure that a class has just a single instance.

- Provide a global access point to that instance.

## When to use

- Use the Singleton pattern when a class in your program should have just a single instance available to all clients; for example, a single database object shared by different parts of the program.

- Use the Singleton pattern when you need stricter control over global variables.

## Pros

- You can be sure that a class has only a single instance.

- You gain a global access point to that instance.

- The singleton object is initialized only when it’s requested for the first time.

## Cons

- Violates the Single Responsibility Principle. The pattern solves two problems at the time.

- The Singleton pattern can mask bad design, for instance, when the components of the program know too much about each other.

- The pattern requires special treatment in a multi-threaded environment so that multiple threads won’t create a singleton object several times.

- It may be difficult to unit test the client code of the Singleton because many test frameworks rely on inheritance when producing mock objects. Since the constructor of the singleton class is private and overriding static methods is impossible in most languages, you will need to think of a creative way to mock the singleton. Or just don’t write the tests. Or don’t use the Singleton pattern.
