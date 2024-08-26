import unittest
from space_age import SpaceAge


def assert_in_delta(expected, actual):
    delta = 0.01
    if abs(expected - actual) > delta:
        raise AssertionError(f"Your result of {actual} should be within {delta} of the expected result {expected}")


class SpaceAgeTest(unittest.TestCase):
    def test_earth_age(self):
        age = SpaceAge(1_000_000_000)
        assert_in_delta(31.69, age.on_earth())

    def test_mercury_age(self):
        age = SpaceAge(2_134_835_688)
        assert_in_delta(280.88, age.on_mercury())

    def test_venus_age(self):
        age = SpaceAge(189_839_836)
        assert_in_delta(9.78, age.on_venus())

    def test_mars_age(self):
        age = SpaceAge(2_329_871_239)
        assert_in_delta(39.25, age.on_mars())

    def test_jupiter_age(self):
        age = SpaceAge(901_876_382)
        assert_in_delta(2.41, age.on_jupiter())

    def test_saturn_age(self):
        age = SpaceAge(3_000_000_000)
        assert_in_delta(3.23, age.on_saturn())

    def test_uranus_age(self):
        age = SpaceAge(3_210_123_456)
        assert_in_delta(1.21, age.on_uranus())

    def test_neptune_age(self):
        age = SpaceAge(8_210_123_456)
        assert_in_delta(1.58, age.on_neptune())


if __name__ == '__main__':
    unittest.main()
