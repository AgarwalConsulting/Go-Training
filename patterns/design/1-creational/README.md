# Creational Patterns

Creational design patterns abstract the instantiation process. They help make a system independent of how it's objects are created, composed and represented.

- Builder
- Factories
- Prototype
- Singleton

## Relationship between Patterns

- Many designs start by using `Factory Method` (less complicated and more customizable via subclasses) and evolve toward `Abstract Factory`, `Prototype`, or `Builder` (more flexible, but more complicated).

- `Builder` focuses on constructing complex objects step by step. `Abstract Factory` specializes in creating families of related objects. `Abstract Factory` returns the product immediately, whereas `Builder` lets you run some additional construction steps before fetching the product.

- You can use `Builder` when creating complex Composite trees because you can program its construction steps to work recursively.

- You can combine `Builder` with `Bridge`: the director class plays the role of the abstraction, while different builders act as implementations.

- `Abstract Factories`, `Builders` and `Prototypes` can all be implemented as `Singletons`.

- You can use `Factory Method` along with `Iterator` to let collection subclasses return different types of iterators that are compatible with the collections.

- `Prototype` isn’t based on inheritance, so it doesn’t have its drawbacks. On the other hand, `Prototype` requires a complicated initialization of the cloned object. `Factory Method` is based on inheritance but doesn’t require an initialization step.
