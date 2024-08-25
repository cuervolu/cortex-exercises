use anagram::anagrams;

#[test]
fn test_no_matches() {
    let word = "diaper";
    let inputs = &["hello", "world", "zombies", "pants"];
    let outputs: Vec<&str> = vec![];
    assert_eq!(anagrams(word, inputs), outputs);
}

#[test]
fn test_detect_simple_anagram() {
    let word = "ant";
    let inputs = &["tan", "stand", "at"];
    let outputs: Vec<&str> = vec!["tan"];
    assert_eq!(anagrams(word, inputs), outputs);
}

#[test]
fn test_does_not_confuse_different_duplicates() {
    let word = "galea";
    let inputs = &["eagle"];
    let outputs: Vec<&str> = vec![];
    assert_eq!(anagrams(word, inputs), outputs);
}

#[test]
fn test_eliminate_anagram_subsets() {
    let word = "good";
    let inputs = &["dog", "goody"];
    let outputs: Vec<&str> = vec![];
    assert_eq!(anagrams(word, inputs), outputs);
}

#[test]
fn test_detect_anagram() {
    let word = "listen";
    let inputs = &["enlists", "google", "inlets", "banana"];
    let outputs: Vec<&str> = vec!["inlets"];
    assert_eq!(anagrams(word, inputs), outputs);
}

#[test]
fn test_multiple_anagrams() {
    let word = "allergy";
    let inputs = &["gallery", "ballerina", "regally", "clergy", "largely", "leading"];
    let mut outputs: Vec<&str> = vec!["gallery", "regally", "largely"];
    outputs.sort_unstable();
    let mut result = anagrams(word, inputs);
    result.sort_unstable();
    assert_eq!(result, outputs);
}

#[test]
fn test_case_insensitive_anagrams() {
    let word = "Orchestra";
    let inputs = &["cashregister", "Carthorse", "radishes"];
    let outputs: Vec<&str> = vec!["Carthorse"];
    assert_eq!(anagrams(word, inputs), outputs);
}

#[test]
fn test_unicode_anagrams() {
    let word = "ΑΒΓ";
    let inputs = &["ΒΓΑ", "ΒΓΔ", "γβα"];
    let outputs: Vec<&str> = vec!["ΒΓΑ", "γβα"];
    assert_eq!(anagrams(word, inputs), outputs);
}

#[test]
fn test_misleading_unicode_anagrams() {
    // Despite what a human might think these words different:
    let word = "ΑΒΓ";
    let inputs = &["ABΓ"];
    let outputs: Vec<&str> = vec![];
    assert_eq!(anagrams(word, inputs), outputs);
}