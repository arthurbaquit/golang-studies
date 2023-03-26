# Builder Pattern

The Builder pattern is a creational design pattern that separates the construction of an object from its representation. It allows for the creation of complex objects in a step-by-step manner, while providing flexibility to use different representations of the same object construction process. Also, with this pattern, the client only have access to the builder, instead of the object itself.

## When to Use the Builder Pattern

Consider using the Builder pattern when:

- You want to create complex objects in a step-by-step manner
- You want to separate the construction of an object from its representation
- You want to provide flexibility to use different representations of the same object construction process

## When NOT to Use the Builder Pattern

One should avoid this pattern particularly if the object being constructed is simple or has a fixed set of properties. In these cases, a simpler factory pattern may be more appropriate.

## Pros

- Allows you to vary a product's internal representation.
- Encapsulates the object construction process, making it easier to manage and modify.
- Provides control over steps of construction process.
- Separates the construction of an object from its representation, improving maintainability and testability.

## Cons

- A distinct ConcreteBuilder must be created for each type of product.
- Builder classes must be mutable.
- Requires the creation of additional classes or functions to implement the pattern, which can increase the codebase size.

## Best Practices

Here are some tips for using the Builder pattern effectively:

- Use the pattern when creating complex objects that can benefit from being constructed in a step-by-step manner.
- Consider using a factory pattern instead of the Builder pattern if the object being constructed is simple or has a fixed set of properties.
- Keep the number of builders to a minimum to avoid adding unnecessary complexity.
- Ensure that the builders are easy to reason about by keeping their functionality limited to a small set of related properties or components.
- Use immutable builders where possible to avoid side effects and improve safety.
- Consider using interfaces to define the builders to make them easier to swap out or mock during testing.

## Summary

- A builder is a separate component used for building a object.
- To make build fluent, return the receiver, which allows chaining.
- Different facets of an object be built with different builders working in tandem via a common struct.
