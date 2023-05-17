# Factory Pattern

The factory pattern is a creational design pattern that can decide which object needs to be created following some default values or business logic. Its approach can both hide the object itself and prevent the client from changing it, or it can expose it, allowing the client to change its default values.

## When to Use the Builder Pattern

Consider using the Factory pattern when:

- Your object has logic associated with it.
- Your object is relatively simple.
- You have a fixed set of properties.

## When NOT to Use the Factory Pattern

Consider not using the factory pattern if your object has many properties or has a complex building process. For those cases, consider using the [builder pattern](../builder/README.md)

## Pros

- Allows you to build objects with default values.
- Can expose the object so the client can change it.
- Implementation flexibility, both with functional and structural approaches.
- Promotes loose coupling between classes by separating the responsibility of object creation from the client code.

## Cons

- Difficult to build complex objects.
- If the business logic changes, it also needs to change the builder.

## Best Practices

Here are some tips for using the Factory pattern effectively:

- Before creating the concrete classes, define an interface that all the classes should implement. This will make it easier to change or add new implementations later on.
- Create a method in the factory that accepts some input and returns an instance of the appropriate concrete class.
- Implement the concrete classes that will be created by the factory method. These classes should all implement the interface defined earlier.
- If you want to provide default values for the objects being created, set them in the factory method.
- Instead of creating objects directly, use the factory method to create them. This will make it easier to change the implementation later on, and will make your code more modular.
- If you're using a dependency injection framework, you can configure it to use the factory method to create objects, instead of creating them directly. This will make it easier to switch out implementations in the future.
- Test the factory to make sure it creates the correct objects with the correct default values. You can also test the concrete classes to make sure they implement the interface correctly.

## Summary

- The constructor is a helper function for creating the objects.
- Factory is any entity (funcs, structs, etc) that handles the object creation.
