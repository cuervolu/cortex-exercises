# Hints

1. The `time` crate is already included as a dependency. You can use its types and methods to work with dates and times.

2. Look into the documentation of the `time` crate to find methods for adding duration to a `PrimitiveDateTime`.

3. A gigasecond is 10^9 seconds. You might find the `Duration` type useful for representing this amount of time.

4. Remember that Rust has type inference, but sometimes you might need to specify types explicitly, especially when working with generic functions or methods.

5. The `assert_eq!` macro in the tests compares the result of your `after` function with the expected date and time. Make sure your implementation returns a `PrimitiveDateTime` that matches the expected result exactly.