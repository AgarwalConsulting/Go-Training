# SOLID principles

In object-oriented computer programming, *SOLID* is a mnemonic acronym for five design principles intended to make software designs more understandable, flexible, and maintainable. Introduced by Robert C. Martin (c. 2000).

- **S**ingle-responsibility principle

- **O**pen-Closed principle

- **L**iskov substitution principle

- **I**nterface segregation principle

- **D**ependency inversion principle

## Single-responsibility principle

> A class should have one, and only one, reason to change. - Robert C Martin

aka. Separation of Concerns

Most common anti-pattern: *God* classes

Benefits:

- Testing – A class with one responsibility will have far fewer test cases

- Lower coupling – Less functionality in a single class will have fewer dependencies

- Organization – Smaller, well-organized classes are easier to search than monolithic ones

## Open-Closed principle

> Software entities should be open for extension, but closed for modification. - Bertrand Meyer

Go does not support function overloading or even overriding

## Liskov Substitution principle

Coined by Barbara Liskov, the Liskov substitution principle states, roughly, that two types are substitutable if they exhibit behaviour such that the caller is unable to tell the difference.

In a class based language, Liskov’s substitution principle is commonly interpreted as a specification for an abstract base class with various concrete subtypes. But Go does not have classes, or inheritance, so substitution cannot be implemented in terms of an abstract class hierarchy.

> Require no more, promise no less. – Jim Weirich

## Interface segregation principle

> Clients should not be forced to depend on methods they do not use. - Robert C Martin

> A great rule of thumb for Go is accept interfaces, return structs. - Jack Lindamood

## Dependency Inversion principle

> High-level modules should not depend on low-level modules. Both should depend on abstractions.
> Abstractions should not depend on details. Details should depend on abstractions.
> – Robert C. Martin

- If you’ve applied all the principles we’ve talked about up to this point then your code should already be factored into discrete packages, each with a single well defined responsibility or purpose.

- Your code should describe its dependencies in terms of interfaces, and those interfaces should be factored to describe only the behaviour those functions require.

- In other words, there shouldn’t be much left to do.

In Go, your import graph must be acyclic. A failure to respect this acyclic requirement is grounds for a compilation failure, but more gravely represents a serious error in design.

---

## Summary

The Single Responsibility Principle encourages you to structure the functions, types, and methods into packages that exhibit natural cohesion; the types belong together, the functions serve a single purpose.

The Open / Closed Principle encourages you to compose simple types into more complex ones using embedding.

The Liskov Substitution Principle encourages you to express the dependencies between your packages in terms of interfaces, not concrete types. By defining small interfaces, we can be more confident that implementations will faithfully satisfy their contract.

The Interface Substitution Principle takes that idea further and encourages you to define functions and methods that depend only on the behaviour that they need. If your function only requires a parameter of an interface type with a single method, then it is more likely that this function has only one responsibility.

The Dependency Inversion Principle encourages you move the knowledge of the things your package depends on from compile time–in Go we see this with a reduction in the number of import statements used by a particular package–to run time.

> interfaces let you apply the SOLID principles to Go programs. - Dave Cheney
