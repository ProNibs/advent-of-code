import unittest
from Day2.main import *

class TestMain(unittest.TestCase):

    def test_single_digit_repeats(self):
        self.assertEqual(check_for_repeats('54-55'), [55])

    def test_double_digit_repeats(self):
        self.assertEqual(check_for_repeats('6463-6466'), [6464])
    def test_triple_digit_repeats(self):
        self.assertEqual(check_for_repeats('123122-123125'), [123123])
    
    def test_based_on_example(self):
        self.assertEqual(check_for_repeats('11-22'), [11,22])
        self.assertEqual(check_for_repeats('95-115'), [99])
        self.assertEqual(check_for_repeats('998-1012'), [1010])
        self.assertEqual(check_for_repeats('1188511880-1188511890'), [1188511885])
        self.assertEqual(check_for_repeats('222220-222224'), [222222])
        self.assertEqual(check_for_repeats('1698522-1698528'), [])
        self.assertEqual(check_for_repeats('446443-446449'), [446446])
        self.assertEqual(check_for_repeats('38593856-38593862'), [38593859])
        self.assertEqual(check_for_repeats('565653-565659'), [])
        self.assertEqual(check_for_repeats('824824821-824824827'), [])
        self.assertEqual(check_for_repeats('2121212118-2121212124'), [])
        
    def test_data_answer(self):
        test_data = readFile('D:/Documents/GitHub/advent-of-code/2025/Day2/testdata.txt')
        self.assertEqual(solution_one(test_data), 1227775554)
        
if __name__ == '__main__':
    unittest.main()