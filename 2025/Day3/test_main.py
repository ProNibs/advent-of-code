import unittest
from Day3.main import *

test_data = readFile('D:/Documents/GitHub/advent-of-code/2025/Day3/testdata.txt')
class TestMain(unittest.TestCase):

    def test_easy_joltage_number(self):
        self.assertEqual(determine_joltage([3,2,1]), 32)

    def test_ignore_last_digi_joltage_number(self):
        self.assertEqual(determine_joltage([3,2,1,4]), 34)
        self.assertEqual(determine_joltage([3,4,1,4]), 34)

    def test_test_data_joltage_number(self):
        self.assertEqual(determine_joltage([9,8,7,6,5,4,3,2,1,1,1,1,1,1,1]), 98)
        self.assertEqual(determine_joltage([8,1,1,1,1,1,1,1,1,1,1,1,1,1,9]), 89)
        self.assertEqual(determine_joltage([2,3,4,2,3,4,2,3,4,2,3,4,2,7,8]), 78)
        self.assertEqual(determine_joltage([8,1,8,1,8,1,9,1,1,1,1,2,1,1,1]), 92)
        
    def test_data_answer(self):
        self.assertEqual(solution_one(test_data), 357)
        # self.assertEqual(solution_two(test_data), 4174379265)
        
if __name__ == '__main__':
    unittest.main()