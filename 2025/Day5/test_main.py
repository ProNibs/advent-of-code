import unittest
from Day5.main import *

test_data = readFile('D:/Documents/GitHub/advent-of-code/2025/Day5/testdata.txt')

class TestMain(unittest.TestCase):


    def test_fresh_ingredient_list(self):
        # self.assertEqual(fresh_ingredients(test_data), [3,4,5,10,11,12,13,14,15,16,17,18,19,20])
        self.assertEqual(fresh_ingredients(test_data), [(3,5), (10,14), (16,20),(12,18)])

    def test_is_ingredient_fresh(self):
        fresh_list = [(3,5)]
        self.assertEqual(isFresh(fresh_list, 2), False)
        self.assertEqual(isFresh(fresh_list, 3), True)
        self.assertEqual(isFresh(fresh_list, 4), True)
        self.assertEqual(isFresh(fresh_list, 5), True)
        self.assertEqual(isFresh(fresh_list, 6), False)
        
    def test_data_answer(self):
        self.assertEqual(solution_one(test_data), 3)
        
if __name__ == '__main__':
    unittest.main()