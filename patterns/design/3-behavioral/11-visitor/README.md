# Visitor

*Visitor* is a behavioral design pattern that lets you separate algorithms from the objects on which they operate.

## Problems

Imagine that your team develops an app which works with geographic information structured as one colossal graph. Each node of the graph may represent a complex entity such as a city, but also more granular things like industries, sightseeing areas, etc. The nodes are connected with others if there’s a road between the real objects that they represent. Under the hood, each node type is represented by its own class, while each specific node is an object.

At some point, you got a task to implement exporting the graph into XML format. At first, the job seemed pretty straightforward. You planned to add an export method to each node class and then leverage recursion to go over each node of the graph, executing the export method. The solution was simple and elegant: thanks to polymorphism, you weren’t coupling the code which called the export method to concrete classes of nodes.

Unfortunately, the system architect refused to allow you to alter existing node classes. He said that the code was already in production and he didn’t want to risk breaking it because of a potential bug in your changes.

Besides, he questioned whether it makes sense to have the XML export code within the node classes. The primary job of these classes was to work with geodata. The XML export behavior would look alien there.

## When to use

- Use the Visitor when you need to perform an operation on all elements of a complex object structure (for example, an object tree).

- Use the Visitor to clean up the business logic of auxiliary behaviors.

- Use the pattern when a behavior makes sense only in some classes of a class hierarchy, but not in others.

## Pros

- Open/Closed Principle. You can introduce a new behavior that can work with objects of different classes without changing these classes.

- Single Responsibility Principle. You can move multiple versions of the same behavior into the same class.

- A visitor object can accumulate some useful information while working with various objects. This might be handy when you want to traverse some complex object structure, such as an object tree, and apply the visitor to each object of this structure.

## Cons

- You need to update all visitors each time a class gets added to or removed from the element hierarchy.

- Visitors might lack the necessary access to the private fields and methods of the elements that they’re supposed to work with.
