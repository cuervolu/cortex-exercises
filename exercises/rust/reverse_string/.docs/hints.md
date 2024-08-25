# Hints

1. You can convert a `String` to a vector of characters using `chars()` method.
2. Rust's `String` type stores data as a UTF-8 encoded byte string, but for this exercise, we only need to handle ASCII
   characters.
3. Remember that `String` in Rust is growable, while `&str` is an immutable reference to a string slice.
4. The `collect()` method can be used to create a `String` from an iterator of characters.
5. Consider using the `rev()` method on the iterator before collecting it back into a `String`.