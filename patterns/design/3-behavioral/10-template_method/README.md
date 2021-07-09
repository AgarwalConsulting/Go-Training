# Template Method

*Template Method* is a behavioral design pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.

## Problems

Imagine that you’re creating a data mining application that analyzes corporate documents. Users feed the app documents in various formats (PDF, DOC, CSV), and it tries to extract meaningful data from these docs in a uniform format.

The first version of the app could work only with DOC files. In the following version, it was able to support CSV files. A month later, you “taught” it to extract data from PDF files.

At some point, you noticed that all three classes have a lot of similar code. While the code for dealing with various data formats was entirely different in all classes, the code for data processing and analysis is almost identical. Wouldn’t it be great to get rid of the code duplication, leaving the algorithm structure intact?

## When to use

- Use the Template Method pattern when you want to let clients extend only particular steps of an algorithm, but not the whole algorithm or its structure.

- Use the pattern when you have several classes that contain almost identical algorithms with some minor differences. As a result, you might need to modify all classes when the algorithm changes.

## Pros

- You can let clients override only certain parts of a large algorithm, making them less affected by changes that happen to other parts of the algorithm.

- You can pull the duplicate code into a superclass.

## Cons

- Some clients may be limited by the provided skeleton of an algorithm.

- You might violate the Liskov Substitution Principle by suppressing a default step implementation via a subclass.

- Template methods tend to be harder to maintain the more steps they have.
