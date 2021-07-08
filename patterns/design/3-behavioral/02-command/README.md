# Command

*Command* is a behavioral design pattern that turns a request into a stand-alone object that contains all information about the request. This transformation lets you pass requests as a method arguments, delay or queue a request’s execution, and support undoable operations.

## Problem

Imagine that you’re working on a new text-editor app. Your current task is to create a toolbar with a bunch of buttons for various operations of the editor. You created a very neat Button class that can be used for buttons on the toolbar, as well as for generic buttons in various dialogs.

While all of these buttons look similar, they’re all supposed to do different things. Where would you put the code for the various click handlers of these buttons? The simplest solution is to create tons of subclasses for each place where the button is used. These subclasses would contain the code that would have to be executed on a button click.

Before long, you realize that this approach is deeply flawed. First, you have an enormous number of subclasses, and that would be okay if you weren’t risking breaking the code in these subclasses each time you modify the base Button class. Put simply, your GUI code has become awkwardly dependent on the volatile code of the business logic.

## When to use

- Use the Command pattern when you want to parametrize objects with operations.

- Use the Command pattern when you want to queue operations, schedule their execution, or execute them remotely.

- Use the Command pattern when you want to implement reversible operations.

## Pros

- Single Responsibility Principle. You can decouple classes that invoke operations from classes that perform these operations.

- Open/Closed Principle. You can introduce new commands into the app without breaking existing client code.

- You can implement undo/redo.

- You can implement deferred execution of operations.

- You can assemble a set of simple commands into a complex one.

## Cons

- The code may become more complicated since you’re introducing a whole new layer between senders and receivers.
