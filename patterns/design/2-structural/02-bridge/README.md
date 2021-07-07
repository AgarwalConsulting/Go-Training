# Bridge Pattern

Bridge is a structural design pattern that divides business logic or huge class into separate class hierarchies that can be developed independently.

Lets you split a large class or a set of closely related classes into two separate hierarchies—abstraction and implementation—which can be developed independently of each other.

## Problem

Say you have a geometric Shape class with a pair of subclasses: Circle and Square. You want to extend this class hierarchy to incorporate colors, so you plan to create Red and Blue shape subclasses. However, since you already have two subclasses, you’ll need to create four class combinations such as BlueCircle and RedSquare.

Adding new shape types and colors to the hierarchy will grow it exponentially. For example, to add a triangle shape you’d need to introduce two subclasses, one for each color. And after that, adding a new color would require creating three subclasses, one for each shape type. The further we go, the worse it becomes.

## When to use

- Use the Bridge pattern when you want to divide and organize a monolithic class that has several variants of some functionality (for example, if the class can work with various database servers).

- Use the pattern when you need to extend a class in several orthogonal (independent) dimensions.

- Use the Bridge if you need to be able to switch implementations at runtime.

## Pros

- You can create platform-independent classes and apps.

- The client code works with high-level abstractions. It isn’t exposed to the platform details.

- Open/Closed Principle. You can introduce new abstractions and implementations independently from each other.

- Single Responsibility Principle. You can focus on high-level logic in the abstraction and on platform details in the implementation.

## Cons

- You might make the code more complicated by applying the pattern to a highly cohesive class.
