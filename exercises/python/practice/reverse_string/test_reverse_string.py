import unittest
from reverse_string import reverse


class ReverseStringTest(unittest.TestCase):
    def test_empty_string(self):
        self.assertEqual(reverse(""), "")

    def test_word(self):
        self.assertEqual(reverse("doctor"), "rotcod")

    def test_uppercase_word(self):
        self.assertEqual(reverse("FUNCTION"), "NOITCNUF")

    def test_with_spaces(self):
        self.assertEqual(reverse("I like cookies"), "seikooc ekil I")

    def test_with_punctuation(self):
        self.assertEqual(reverse("Hello, World!"), "!dlroW ,olleH")


if __name__ == '__main__':
    unittest.main()
