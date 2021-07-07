# Builder pattern

Lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.

The potential problem with the multistage building process is that a partially built and unstable product may be exposed to the client. The Builder pattern keeps the product private until it’s fully built.

## Problem

Imagine a complex object that requires laborious, step-by-step initialization of many fields and nested objects. Such initialization code is usually buried inside a monstrous constructor with lots of parameters. Or even worse: scattered all over the client code.

## When to use

- Use the Builder pattern to get rid of a “telescopic constructor”.

- Use Builder pattern when the object constructed is big and requires multiple steps. It helps in less size of the constructor.  The construction of the house becomes simple and it does not require a large constructor

- Use the Builder pattern when you want your code to be able to create different representations of some product (for example, stone and wooden houses).

- Use the Builder to construct Composite trees or other complex objects.

## Pros

- You can construct objects step-by-step, defer construction steps or run steps recursively.

- You can reuse the same construction code when building various representations of products.

- Single Responsibility Principle. You can isolate complex construction code from the business logic of the product.

## Cons

- The overall complexity of the code increases since the pattern requires creating multiple new classes.
