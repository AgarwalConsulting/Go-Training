# Adapter Pattern

Provide an interface for creating families of related or dependent objects without specifying their concrete classes.

A construct which adapts an existing interface X to conform to existing interface Y.

## Problem

Imagine that you’re creating a stock market monitoring app. The app downloads the stock data from multiple sources in XML format and then displays nice-looking charts and diagrams for the user.

At some point, you decide to improve the app by integrating a smart 3rd-party analytics library. But there’s a catch: the analytics library only works with data in JSON format.

You could change the library to work with XML. However, this might break some existing code that relies on the library. And worse, you might not have access to the library’s source code in the first place, making this approach impossible.

## When to use

- Use the Adapter class when you want to use some existing class, but its interface isn’t compatible with the rest of your code.

- Use the pattern when you want to reuse several existing subclasses that lack some common functionality that can’t be added to the superclass.

## Pros

- Single Responsibility Principle. You can separate the interface or data conversion code from the primary business logic of the program.

- Open/Closed Principle. You can introduce new types of adapters into the program without breaking the existing client code, as long as they work with the adapters through the client interface.

## Cons

- The overall complexity of the code increases because you need to introduce a set of new interfaces and classes. Sometimes it’s simpler just to change the service class so that it matches the rest of your code.
