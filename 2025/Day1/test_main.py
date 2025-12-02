import unittest
from Day1.main import *

class TestMain(unittest.TestCase):

    def test_right(self):
        self.assertEqual(shift(50, 'R3'), 53)

    def test_left(self):
        self.assertEqual(shift(50, 'L3'), 47)
    
    def test_loop_right(self):
        self.assertEqual(shift(90, 'R10'), 0)
        self.assertEqual(shift(90, 'R11'), 1)

    def test_loop_left(self):
        self.assertEqual(shift(10, 'L10'), 0)
        self.assertEqual(shift(10, 'L11'), 99)

    # Test to see if we can do all the rotations proper
    def test_rotations(self):
        test_data = readFile('D:/Documents/GitHub/advent-of-code/2025/Day1/testdata.txt')
        self.assertEqual(rotation(test_data), (32, 3))

    # Noticed test data had some R203 craziness
    def test_super_rotations(self):
        self.assertEqual(shift(80, 'R203'), 83)

if __name__ == '__main__':
    unittest.main()