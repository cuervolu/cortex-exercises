import unittest
from anagram import anagrams


class AnagramTest(unittest.TestCase):
    def test_no_matches(self):
        word = "diaper"
        inputs = ["hello", "world", "zombies", "pants"]
        self.assertEqual(anagrams(word, inputs), [])

    def test_detect_simple_anagram(self):
        word = "ant"
        inputs = ["tan", "stand", "at"]
        self.assertEqual(anagrams(word, inputs), ["tan"])

    def test_does_not_confuse_different_duplicates(self):
        word = "galea"
        inputs = ["eagle"]
        self.assertEqual(anagrams(word, inputs), [])

    def test_eliminate_anagram_subsets(self):
        word = "good"
        inputs = ["dog", "goody"]
        self.assertEqual(anagrams(word, inputs), [])

    def test_detect_anagram(self):
        word = "listen"
        inputs = ["enlists", "google", "inlets", "banana"]
        self.assertEqual(anagrams(word, inputs), ["inlets"])

    def test_multiple_anagrams(self):
        word = "allergy"
        inputs = ["gallery", "ballerina", "regally", "clergy", "largely", "leading"]
        self.assertEqual(sorted(anagrams(word, inputs)), sorted(["gallery", "regally", "largely"]))

    def test_case_insensitive_anagrams(self):
        word = "Orchestra"
        inputs = ["cashregister", "Carthorse", "radishes"]
        self.assertEqual(anagrams(word, inputs), ["Carthorse"])

    def test_unicode_anagrams(self):
        word = "ΑΒΓ"
        inputs = ["ΒΓΑ", "ΒΓΔ", "γβα"]
        self.assertEqual(sorted(anagrams(word, inputs)), sorted(["ΒΓΑ", "γβα"]))

    def test_misleading_unicode_anagrams(self):
        word = "ΑΒΓ"
        inputs = ["ABΓ"]
        self.assertEqual(anagrams(word, inputs), [])


if __name__ == '__main__':
    unittest.main()
