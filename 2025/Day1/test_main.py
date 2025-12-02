import unittest
from Day1 import main

class TestMain(unittest.TestCase):

    def test_right():
        self.assertEqual(right_shift(50, 'R3'), 53)
    
if __name__ == '__main__':
    unittest.main()