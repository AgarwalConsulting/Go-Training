# Factories

Factory Method is a creational design pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.

The Factory Method pattern suggests that you replace direct object construction calls with calls to a special factory method.

## Problem

Imagine that you’re creating a logistics management application. The first version of your app can only handle transportation by trucks, so the bulk of your code lives inside the Truck class.

At present, most of your code is coupled to the Truck class. Adding Ships into the app would require making changes to the entire codebase. Moreover, if later you decide to add another type of transportation to the app, you will probably need to make all of these changes again.

## When to use

- Use the Factory Method when you don’t know beforehand the exact types and dependencies of the objects your code should work with.

- The Factory Method separates product construction code from the code that actually uses the product. Therefore it’s easier to extend the product construction code independently from the rest of the code.

- Use the Factory Method when you want to provide users of your library or framework with a way to extend its internal components.

- Use the Factory Method when you want to save system resources by reusing existing objects instead of rebuilding them each time.

## Pros

- You avoid tight coupling between the creator and the concrete products.

- Single Responsibility Principle. You can move the product creation code into one place in the program, making the code easier to support.

- Open/Closed Principle. You can introduce new types of products into the program without breaking existing client code.

## Cons

- The code may become more complicated since you need to introduce a lot of new subclasses to implement the pattern. The best case scenario is when you’re introducing the pattern into an existing hierarchy of creator classes.
