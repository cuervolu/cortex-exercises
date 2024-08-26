import unittest
from hello_world import hello


class HelloWorldTest(unittest.TestCase):
    def test_hello_world(self):
        self.assertEqual("Hello, World!", hello())


if __name__ == '__main__':
    unittest.main()
