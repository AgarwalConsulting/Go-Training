# State

*State* is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.

The State pattern is closely related to the concept of a Finite-State Machine.

## Problems

Imagine that we have a Document class. A document can be in one of three states: Draft, Moderation and Published. The publish method of the document works a little bit differently in each state:

- In `Draft`, it moves the document to moderation.

- In `Moderation`, it makes the document public, but only if the current user is an administrator.

- In `Published`, it doesn’t do anything at all.

## When to use

- Use the State pattern when you have an object that behaves differently depending on its current state, the number of states is enormous, and the state-specific code changes frequently.

- Use the pattern when you have a class polluted with massive conditionals that alter how the class behaves according to the current values of the class’s fields.

- Use State when you have a lot of duplicate code across similar states and transitions of a condition-based state machine.

## Pros

- Single Responsibility Principle. Organize the code related to particular states into separate classes.

- Open/Closed Principle. Introduce new states without changing existing state classes or the context.

- Simplify the code of the context by eliminating bulky state machine conditionals.

## Cons

- Applying the pattern can be overkill if a state machine has only a few states or rarely changes.
