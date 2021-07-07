# Composite Pattern

Composite is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.

Using the Composite pattern makes sense only when the core model of your app can be represented as a tree.

## Problem

Imagine that you have two types of objects: Products and Boxes. A Box can contain several Products as well as a number of smaller Boxes. These little Boxes can also hold some Products or even smaller Boxes, and so on.

Say you decide to create an ordering system that uses these classes. Orders could contain simple products without any wrapping, as well as boxes stuffed with products...and other boxes. How would you determine the total price of such an order?

## When to use

- Use the Composite pattern when you have to implement a tree-like object structure.

- Use the pattern when you want the client code to treat both simple and complex elements uniformly.

## Pros

- You can work with complex tree structures more conveniently: use polymorphism and recursion to your advantage.

- Open/Closed Principle. You can introduce new element types into the app without breaking the existing code, which now works with the object tree.

## Cons

- It might be difficult to provide a common interface for classes whose functionality differs too much. In certain scenarios, you’d need to overgeneralize the component interface, making it harder to comprehend.
