# Mediator

*Mediator* is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.

## Problems

Say you have a dialog for creating and editing customer profiles. It consists of various form controls such as text fields, checkboxes, buttons, etc.

Some of the form elements may interact with others. For instance, selecting the “I have a dog” checkbox may reveal a hidden text field for entering the dog’s name. Another example is the submit button that has to validate values of all fields before saving the data.

By having this logic implemented directly inside the code of the form elements you make these elements’ classes much harder to reuse in other forms of the app. For example, you won’t be able to use that checkbox class inside another form, because it’s coupled to the dog’s text field. You can use either all the classes involved in rendering the profile form, or none at all.

## When to use

- Use the Mediator pattern when it’s hard to change some of the classes because they are tightly coupled to a bunch of other classes.

- Use the pattern when you can’t reuse a component in a different program because it’s too dependent on other components.

- Use the Mediator when you find yourself creating tons of component subclasses just to reuse some basic behavior in various contexts.

## Pros

- Single Responsibility Principle. You can extract the communications between various components into a single place, making it easier to comprehend and maintain.

- Open/Closed Principle. You can introduce new mediators without having to change the actual components.

- You can reduce coupling between various components of a program.

- You can reuse individual components more easily.

## Cons

- Over time a mediator can evolve into a God Object.
