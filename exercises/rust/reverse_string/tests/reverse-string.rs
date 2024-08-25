use reverse_string::reverse;

#[test]
fn test_empty_string() {
    assert_eq!("", reverse(""));
}

#[test]
fn test_word() {
    assert_eq!("rotcod", reverse("doctor"));
}

#[test]
fn test_uppercase_word() {
    assert_eq!("NOITCNUF", reverse("FUNCTION"));
}

#[test]
fn test_with_spaces() {
    assert_eq!("seikooc ekil I", reverse("I like cookies"));
}

#[test]
fn test_with_punctuation() {
    assert_eq!("!dlroW ,olleH", reverse("Hello, World!"));
}