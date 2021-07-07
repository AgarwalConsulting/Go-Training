# Structural Patterns

Structural patterns are concerned with how objects are composed of to form large structures.

Structural patterns explain how to assemble objects and classes into larger structures while keeping these structures flexible and efficient.

- Adapter
- Bridge
- Composite
- Decorator
- Facade
- Flyweight
- Proxy

## Relationship between patterns

- `Bridge` is usually designed up-front, letting you develop parts of an application independently of each other. On the other hand, `Adapter` is commonly used with an existing app to make some otherwise-incompatible classes work together nicely.

- `Adapter` changes the interface of an existing object, while `Decorator` enhances an object without changing its interface. In addition, `Decorator` supports recursive composition, which isn’t possible when you use `Adapter`.

- `Adapter` provides a different interface to the wrapped object, `Proxy` provides it with the same interface, and `Decorator` provides it with an enhanced interface.

- `Facade` defines a new interface for existing objects, whereas `Adapter` tries to make the existing interface usable. `Adapter` usually wraps just one object, while `Facade` works with an entire subsystem of objects.

- `Bridge`, `State`, `Strategy` (and to some degree `Adapter`) have very similar structures. Indeed, all of these patterns are based on composition, which is delegating work to other objects. However, they all solve different problems. A pattern isn’t just a recipe for structuring your code in a specific way. It can also communicate to other developers the problem the pattern solves.

- You can use `Abstract Factory` along with `Bridge`. This pairing is useful when some abstractions defined by `Bridge` can only work with specific implementations. In this case, `Abstract Factory` can encapsulate these relations and hide the complexity from the client code.

- You can combine `Builder` with `Bridge`: the director class plays the role of the abstraction, while different builders act as implementations.

- `Composite` and `Decorator` have similar structure diagrams since both rely on recursive composition to organize an open-ended number of objects.

- `Chain of Responsibility` is often used in conjunction with `Composite`. In this case, when a leaf component gets a request, it may pass it through the chain of all of the parent components down to the root of the object tree.

- You can use `Iterators` to traverse `Composite` trees.

- You can use `Visitor` to execute an operation over an entire `Composite` tree.

- You can implement shared leaf nodes of the `Composite` tree as `Flyweights` to save some RAM.

- Designs that make heavy use of `Composite` and `Decorator` can often benefit from using `Prototype`. Applying the pattern lets you clone complex structures instead of re-constructing them from scratch.

- `Decorator` and `Proxy` have similar structures, but very different intents. Both patterns are built on the composition principle, where one object is supposed to delegate some of the work to another. The difference is that a `Proxy` usually manages the life cycle of its service object on its own, whereas the composition of `Decorators` is always controlled by the client.

- `Decorator` lets you change the skin of an object, while `Strategy` lets you change the guts.

- A `Facade` class can often be transformed into a `Singleton` since a single facade object is sufficient in most cases.

- `Facade` and `Mediator` have similar jobs: they try to organize collaboration between lots of tightly coupled classes.

- `Facade` is similar to `Proxy` in that both buffer a complex entity and initialize it on its own. Unlike `Facade`, `Proxy` has the same interface as its service object, which makes them interchangeable.

- `Flyweight` would resemble `Singleton` if you somehow managed to reduce all shared states of the objects to just one flyweight object. But there are two fundamental differences between these patterns:

  - There should be only one Singleton instance, whereas a Flyweight class can have multiple instances with different intrinsic states.

  - The Singleton object can be mutable. Flyweight objects are immutable.

- `Flyweight` shows how to make lots of little objects, whereas `Facade` shows how to make a single object that represents an entire subsystem.
