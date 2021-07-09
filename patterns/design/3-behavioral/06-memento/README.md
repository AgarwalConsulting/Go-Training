# Memento

*Memento* is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation.

## Problems

Imagine that you’re creating a text editor app. In addition to simple text editing, your editor can format text, insert inline images, etc.

At some point, you decided to let users undo any operations carried out on the text. This feature has become so common over the years that nowadays people expect every app to have it. For the implementation, you chose to take the direct approach. Before performing any operation, the app records the state of all objects and saves it in some storage. Later, when a user decides to revert an action, the app fetches the latest snapshot from the history and uses it to restore the state of all objects.

Let’s think about those state snapshots. How exactly would you produce one? You’d probably need to go over all the fields in an object and copy their values into storage. However, this would only work if the object had quite relaxed access restrictions to its contents. Unfortunately, most real objects won’t let others peek inside them that easily, hiding all significant data in private fields.

Ignore that problem for now and let’s assume that our objects behave like hippies: preferring open relations and keeping their state public. While this approach would solve the immediate problem and let you produce snapshots of objects’ states at will, it still has some serious issues. In the future, you might decide to refactor some of the editor classes, or add or remove some of the fields. Sounds easy, but this would also require changing the classes responsible for copying the state of the affected objects.

## When to use

- Use the Memento pattern when you want to produce snapshots of the object’s state to be able to restore a previous state of the object.

- Use the pattern when direct access to the object’s fields/getters/setters violates its encapsulation.

## Pros

- You can produce snapshots of the object’s state without violating its encapsulation.

- You can simplify the originator’s code by letting the caretaker maintain the history of the originator’s state.

## Cons

- The app might consume lots of RAM if clients create mementos too often.

- Caretakers should track the originator’s lifecycle to be able to destroy obsolete mementos.

- Most dynamic programming languages, such as PHP, Python and JavaScript, can’t guarantee that the state within the memento stays untouched.
