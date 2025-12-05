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
        
    def test_unique_fresh_list_overlap(self):
        self.assertEqual(unique_fresh([(3,5), (4,6)]), [(3,6)])
        self.assertEqual(unique_fresh([(3,5), (2,4)]), [(2,5)])
        self.assertEqual(unique_fresh([(3, 5), (10, 18), (12, 20)]), [(3,5),(10,20)])
 

    def test_unique_fresh_list_shared_edge(self):
        self.assertEqual(unique_fresh([(3,5), (5,7)]), [(3,7)])
        self.assertEqual(unique_fresh([(3,5), (1,3)]), [(1,5)])
    
    def test_unique_fresh_list_adjacent(self):
        self.assertEqual(unique_fresh([(3,5), (6,7)]), [(3,7)])
        self.assertEqual(unique_fresh([(3,5), (1,2)]), [(1,5)])
    
    def test_unique_fresh_list_no_overlap(self):
        self.assertEqual(unique_fresh([(3,5), (7,8)]), [(3,5),(7,8)])
    
    def test_unique_fresh_list_all_overlap(self):
        self.assertEqual(unique_fresh([(3,6), (4,5)]), [(3,6)])

    def test_sort_by_lower_end(self):
        self.assertEqual(sorted([(3,5),(10,14),(16,20),(12,18)]), [(3,5), (10,14), (12,18),(16,20)])

    # def test_combine_after_first_pass(self):
    #     self.assertEqual(unique_fresh([(3, 5), (10, 18), (16, 20)]), [(3,5), (10,20)]) 
    
    # def test_unique_fresh_against_test_data(self):
    #     fresh_items = fresh_ingredients(test_data)
    #     self.assertEqual(unique_fresh(fresh_items), [(3,5),(10,20)])

    def test_data_answer(self):
        self.assertEqual(solution_one(test_data), 3)
        self.assertEqual(solution_two(test_data), ([(3,5),(10,20)],14))
        
if __name__ == '__main__':
    unittest.main()