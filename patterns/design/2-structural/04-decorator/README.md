# Decorator

*Decorator* is a structural design pattern that lets you attach new behaviors to objects by placing these objects inside special wrapper objects that contain the behaviors.

## Problem

Imagine that you’re working on a notification library which lets other programs notify their users about important events.

The initial version of the library was based on the Notifier class that had only a few fields, a constructor and a single send method.

At some point, you realize that users of the library expect more than just email notifications. Many of them would like to receive an SMS about critical issues. Others would like to be notified on Facebook and, of course, the corporate users would love to get Slack notifications.

How hard can that be? You extended the Notifier class and put the additional notification methods into new subclasses. Now the client was supposed to instantiate the desired notification class and use it for all further notifications.

But then someone reasonably asked you, "Why can’t you use several notification types at once? If your house is on fire, you’d probably want to be informed through every channel."

## When to use

- Use the Decorator pattern when you need to be able to assign extra behaviors to objects at runtime without breaking the code that uses these objects.

- Use the pattern when it’s awkward or not possible to extend an object’s behavior using inheritance.

## Pros

- You can extend an object’s behavior without making a new subclass.

- You can add or remove responsibilities from an object at runtime.

- You can combine several behaviors by wrapping an object into multiple decorators.

- Single Responsibility Principle. You can divide a monolithic class that implements many possible variants of behavior into several smaller classes.

## Cons

- It’s hard to remove a specific wrapper from the wrappers stack.

- It’s hard to implement a decorator in such a way that its behavior doesn’t depend on the order in the decorators stack.

- The initial configuration code of layers might look pretty ugly.
