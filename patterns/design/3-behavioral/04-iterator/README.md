# Iterator

*Iterator* is a behavioral design pattern that lets you traverse elements of a collection without exposing its underlying representation (list, stack, tree, etc.).

## Problems

Most collections store their elements in simple lists. However, some of them are based on stacks, trees, graphs and other complex data structures.

But no matter how a collection is structured, it must provide some way of accessing its elements so that other code can use these elements. There should be a way to go through each element of the collection without accessing the same elements over and over.

This may sound like an easy job if you have a collection based on a list. You just loop over all of the elements. But how do you sequentially traverse elements of a complex data structure, such as a tree? For example, one day you might be just fine with depth-first traversal of a tree. Yet the next day you might require breadth-first traversal. And the next week, you might need something else, like random access to the tree elements.

## When to use

- Use the Iterator pattern when your collection has a complex data structure under the hood, but you want to hide its complexity from clients (either for convenience or security reasons).

- Use the pattern to reduce duplication of the traversal code across your app.

- Use the Iterator when you want your code to be able to traverse different data structures or when types of these structures are unknown beforehand.

## Pros

- Single Responsibility Principle. You can clean up the client code and the collections by extracting bulky traversal algorithms into separate classes.

- Open/Closed Principle. You can implement new types of collections and iterators and pass them to existing code without breaking anything.

- You can iterate over the same collection in parallel because each iterator object contains its own iteration state.

- For the same reason, you can delay an iteration and continue it when needed.

## Cons

- Applying the pattern can be an overkill if your app only works with simple collections.

- Using an iterator may be less efficient than going through elements of some specialized collections directly.
