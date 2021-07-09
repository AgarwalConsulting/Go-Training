# Behavioral Patterns

Behavioral patterns are concerned with algorithms and the assignment of responsibilities between objects.

Behavioral patterns describe not just patterns of objects but also the patterns of communication between them.

- Chain of Responsibility
- Command
- Interpreter
- Iterator
- Mediator
- Memento
- Observer
- State
- Strategy
- Template Method
- Visitor

## Relationship between patterns

- `Chain of Responsibility`, `Command`, `Mediator` and `Observer` address various ways of connecting senders and receivers of requests:

  - `Chain of Responsibility` passes a request sequentially along a dynamic chain of potential receivers until one of them handles it.

  - `Command` establishes unidirectional connections between senders and receivers.

  - `Mediator` eliminates direct connections between senders and receivers, forcing them to communicate indirectly via a `mediator` object.

  - `Observer` lets receivers dynamically subscribe to and unsubscribe from receiving requests.

- `Chain of Responsibility` is often used in conjunction with `Composite`. In this case, when a leaf component gets a request, it may pass it through the chain of all of the parent components down to the root of the object tree.

- Handlers in `Chain of Responsibility` can be implemented as `Command`s. In this case, you can execute a lot of different operations over the same context object, represented by a request.

  - However, there’s another approach, where the request itself is a `Command` object. In this case, you can execute the same operation in a series of different contexts linked into a chain.

- `Chain of Responsibility` and `Decorator` have very similar class structures. Both patterns rely on recursive composition to pass the execution through a series of objects. However, there are several crucial differences.

- The `CoR` handlers can execute arbitrary operations independently of each other. They can also stop passing the request further at any point. On the other hand, various Decorators can extend the object’s behavior while keeping it consistent with the base interface. In addition, decorators aren’t allowed to break the flow of the request.

- `Command` and `Strategy` may look similar because you can use both to parameterize an object with some action. However, they have very different intents.

  - You can use `Command` to convert any operation into an object. The operation’s parameters become fields of that object. The conversion lets you defer execution of the operation, queue it, store the history of commands, send commands to remote services, etc.

  - On the other hand, `Strategy` usually describes different ways of doing the same thing, letting you swap these algorithms within a single context class.

- `Prototype` can help when you need to save copies of `Commands` into history.

- You can treat `Visitor` as a powerful version of the `Command` pattern. Its objects can execute operations over various objects of different classes.

- `Facade` and `Mediator` have similar jobs: they try to organize collaboration between lots of tightly coupled classes.

  - `Facade` defines a simplified interface to a subsystem of objects, but it doesn’t introduce any new functionality. The subsystem itself is unaware of the `facade`. Objects within the subsystem can communicate directly.

  - `Mediator` centralizes communication between components of the system. The components only know about the `mediator` object and don’t communicate directly.

- The difference between `Mediator` and `Observer` is often elusive. In most cases, you can implement either of these patterns; but sometimes you can apply both simultaneously.

- You can use `Memento` along with `Iterator` to capture the current iteration state and roll it back if necessary.
