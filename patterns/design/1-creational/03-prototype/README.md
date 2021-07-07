# Prototype pattern

Prototype is a creational design pattern that lets you copy existing objects without making your code dependent on their classes.

- Complicated objects aren't designed from scratch

- An existing design (partial or otherwise) is a Prototype

- We make a copy of the prototype and customize it

- To perform a copy of a prototype, we need to be able to perform a deep copy

The Prototype pattern delegates the cloning process to the actual objects that are being cloned. The pattern declares a common interface for all objects that support cloning. This interface lets you clone an object without coupling your code to the class of that object. Usually, such an interface contains just a single clone method.

## Problem

Say you have an object, and you want to create an exact copy of it. How would you do it? First, you have to create a new object of the same class. Then you have to go through all the fields of the original object and copy their values over to the new object.

## When to use

- Use the Prototype pattern when your code shouldn’t depend on the concrete classes of objects that you need to copy.

- Use the pattern when you want to reduce the number of subclasses that only differ in the way they initialize their respective objects. Somebody could have created these subclasses to be able to create objects with a specific configuration.

## Pros

- You can clone objects without coupling to their concrete classes.

- You can get rid of repeated initialization code in favor of cloning pre-built prototypes.

- You can produce complex objects more conveniently.

- You get an alternative to inheritance when dealing with configuration presets for complex objects.

## Cons

- Cloning complex objects that have circular references might be very tricky.
