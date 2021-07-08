# Chain of Responsibility

*Chain of Responsibility* is behavioral design pattern that allows passing request along the chain of potential handlers until one of them handles request.

Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.

The pattern allows multiple objects to handle the request without coupling sender class to the concrete classes of the receivers. The chain can be composed dynamically at runtime with any handler that follows a standard handler interface.

## Problem

A hospital could have multiple departments such as:

- Reception
- Doctor
- Medicine room
- Cashier

Whenever any patient arrives, they first get to Reception, then to Doctor, then to Medicine Room, and then to Cashier (and so on). The patient is being sent through a chain of departments, where each department sends the patient further down the chain once their function is completed.

The pattern is applicable when there are multiple candidates to process the same request. It is also useful when you don’t want the client to choose the receiver as there are multiple objects can handle the request. Another useful case is when you want to decouple the client from receivers—the client will only need to know the first element in the chain.

## When to use

- Use the Chain of Responsibility pattern when your program is expected to process different kinds of requests in various ways, but the exact types of requests and their sequences are unknown beforehand.

- Use the pattern when it’s essential to execute several handlers in a particular order.

- Use the CoR pattern when the set of handlers and their order are supposed to change at runtime.

## Pros

- You can control the order of request handling.

- Single Responsibility Principle. You can decouple classes that invoke operations from classes that perform operations.

- Open/Closed Principle. You can introduce new handlers into the app without breaking the existing client code.

## Cons

- Some requests may end up unhandled.
