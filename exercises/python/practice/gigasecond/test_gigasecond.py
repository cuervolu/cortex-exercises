import unittest
from datetime import datetime
from gigasecond import add_gigasecond


def create_datetime(year, month, day, hour=0, minute=0, second=0):
    return datetime(year, month, day, hour, minute, second)


class GigasecondTest(unittest.TestCase):
    def test_date_only_specification_of_time(self):
        self.assertEqual(
            add_gigasecond(create_datetime(2011, 4, 25)),
            create_datetime(2043, 1, 1, 1, 46, 40)
        )

    def test_another_date_only_specification_of_time(self):
        self.assertEqual(
            add_gigasecond(create_datetime(1977, 6, 13)),
            create_datetime(2009, 2, 19, 1, 46, 40)
        )

    def test_third_date_only_specification_of_time(self):
        self.assertEqual(
            add_gigasecond(create_datetime(1959, 7, 19)),
            create_datetime(1991, 3, 27, 1, 46, 40)
        )

    def test_full_time_specified(self):
        self.assertEqual(
            add_gigasecond(create_datetime(2015, 1, 24, 22, 0, 0)),
            create_datetime(2046, 10, 2, 23, 46, 40)
        )

    def test_full_time_with_day_roll_over(self):
        self.assertEqual(
            add_gigasecond(create_datetime(2015, 1, 24, 23, 59, 59)),
            create_datetime(2046, 10, 3, 1, 46, 39)
        )


if __name__ == '__main__':
    unittest.main()
