# Facade

*Facade* is a structural design pattern that provides a simplified interface to a library, a framework, or any other complex set of classes.

A facade is a class that provides a simple interface to a complex subsystem which contains lots of moving parts. A facade might provide limited functionality in comparison to working with the subsystem directly. However, it includes only those features that clients really care about.

## Problem

Imagine that you must make your code work with a broad set of objects that belong to a sophisticated library or framework. Ordinarily, you’d need to initialize all of those objects, keep track of dependencies, execute methods in the correct order, and so on.

As a result, the business logic of your classes would become tightly coupled to the implementation details of 3rd-party classes, making it hard to comprehend and maintain.

## When to use

- Use the Facade pattern when you need to have a limited but straightforward interface to a complex subsystem.

- Use the Facade when you want to structure a subsystem into layers.

## Pros

- You can isolate your code from the complexity of a subsystem.

## Cons

- A facade can become a [god object](https://en.wikipedia.org/wiki/God_object) coupled to all classes of an app.
