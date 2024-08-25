use hello_world::hello;

#[test]
fn test_hello_world() {
    assert_eq!("Hello, World!", hello());
}