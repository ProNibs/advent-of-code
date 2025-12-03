import unittest
from Day1.main import *

class TestMain(unittest.TestCase):

    def test_right(self):
        self.assertEqual(shift(50, 'R3')[0], 53)

    def test_left(self):
        self.assertEqual(shift(50, 'L3')[0], 47)
    
    def test_loop_right(self):
        self.assertEqual(shift(90, 'R10')[0], 0)
        self.assertEqual(shift(90, 'R11')[0], 1)

    def test_loop_left(self):
        self.assertEqual(shift(10, 'L10')[0], 0)
        self.assertEqual(shift(10, 'L11')[0], 99)

    # Test to see if we can do all the rotations proper
    def test_rotations(self):
        test_data = readFile('D:/Documents/GitHub/advent-of-code/2025/Day1/testdata.txt')
        # Altered for part 2
        self.assertEqual(rotation(test_data), (32, 6))

    # Noticed test data had some R203 craziness
    def test_super_rotations(self):
        self.assertEqual(shift(80, 'R203')[0], 83)

    # Tests for Part 2
    def test_we_passed_zero_right(self):
        self.assertEqual(shift(90, 'R10')[1], 0)
        self.assertEqual(shift(90, 'R11')[1], 1)

    def test_we_passed_zero_left(self):
        self.assertEqual(shift(10, 'L10')[1], 0)
        self.assertEqual(shift(10, 'L11')[1], 1)
    
    def test_we_start_zero(self):
        self.assertEqual(shift(0, 'L10')[1], 0)
        self.assertEqual(shift(0, 'R10')[1], 0)

    def test_super_rotations(self):
        self.assertEqual(shift(50, 'R1000')[1], 10)
        self.assertEqual(shift(50, 'L1000')[1], 10)
        self.assertEqual(shift(80, 'R203')[1], 2)
        self.assertEqual(shift(80, 'L203')[1], 2)

    def test_rotate_to_zero(self):
        self.assertEqual(rotation(['R50']), (0,1))
        self.assertEqual(rotation(['R150']), (0,2))
        self.assertEqual(rotation(['R250']), (0,3))
        self.assertEqual(rotation(['L50']), (0,1))
        self.assertEqual(rotation(['L150']), (0,2))
        self.assertEqual(rotation(['R250']), (0,3))

    # Answer is wrong, so I must be missing some edge case
    def test_edge_cases(self):
        # Go to zero, then move off of it
        self.assertEqual(rotation(['L50', 'R2'])[1], 1)
        self.assertEqual(rotation(['R50', 'L2'])[1], 1)
        self.assertEqual(rotation(['L50', 'R2', 'L102'])[1], 3)
        self.assertEqual(rotation(['R50', 'L2', 'R102'])[1], 3)
        self.assertEqual(rotation(['L150', 'R2', 'L202'])[1], 5)
        self.assertEqual(rotation(['R150', 'L2', 'R202'])[1], 5)
        
        # Double-checking dumb stuff I wrote
        self.assertEqual(rotation(['L50', 'R1', 'L99'])[1], 2)
        self.assertEqual(rotation(['L50', 'R1', 'L100'])[1], 2)
        self.assertEqual(rotation(['L50', 'R1', 'L101'])[1], 3)
        self.assertEqual(rotation(['L50', 'R1', 'L102'])[1], 3)

        self.assertEqual(rotation(['R50', 'L1', 'R99'])[1], 2)
        self.assertEqual(rotation(['R50', 'L1', 'R100'])[1], 2)
        self.assertEqual(rotation(['R50', 'L1', 'R101'])[1], 3)
        self.assertEqual(rotation(['R50', 'L1', 'R102'])[1], 3)
        
        # Double-checking the crazy numbers math
        self.assertEqual(rotation(['L863', 'L595', 'R67'])[1], 9+6+1)
        self.assertEqual(rotation(['R863', 'R595', 'L67'])[1], 9+6+1)
        # This one is from the beginning of the data.txt answer file
        self.assertEqual(rotation(['R49', 'R27', 'R22', 'R5', 'R6', 'R10', 'L13'])[1], 1)
        self.assertEqual(rotation(['R50', 'L50', 'R50', 'L100', 'R200'])[1], 5)
        
if __name__ == '__main__':
    unittest.main()