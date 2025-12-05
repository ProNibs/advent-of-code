import unittest
from Day2.main import *

class TestMain(unittest.TestCase):

    def test_single_digit_repeats(self):
        self.assertEqual(check_for_repeats_part_one('54-55'), [55])

    def test_double_digit_repeats(self):
        self.assertEqual(check_for_repeats_part_one('6463-6466'), [6464])
    def test_triple_digit_repeats(self):
        self.assertEqual(check_for_repeats_part_one('123122-123125'), [123123])
    
    def test_based_on_example(self):
        self.assertEqual(check_for_repeats_part_one('11-22'), [11,22])
        self.assertEqual(check_for_repeats_part_one('95-115'), [99])
        self.assertEqual(check_for_repeats_part_one('998-1012'), [1010])
        self.assertEqual(check_for_repeats_part_one('1188511880-1188511890'), [1188511885])
        self.assertEqual(check_for_repeats_part_one('222220-222224'), [222222])
        self.assertEqual(check_for_repeats_part_one('1698522-1698528'), [])
        self.assertEqual(check_for_repeats_part_one('446443-446449'), [446446])
        self.assertEqual(check_for_repeats_part_one('38593856-38593862'), [38593859])
        self.assertEqual(check_for_repeats_part_one('565653-565659'), [])
        self.assertEqual(check_for_repeats_part_one('824824821-824824827'), [])
        self.assertEqual(check_for_repeats_part_one('2121212118-2121212124'), [])
        
    # Part 2 is...different.
    def test_single_digit_repeats_part_two(self):
        self.assertEqual(check_for_repeats_part_two('54-55'), [55])
    def test_double_digit_repeats_part_two(self):
        self.assertEqual(check_for_repeats_part_two('6463-6466'), [6464])
    def test_triple_digit_repeats_part_two(self):
        self.assertEqual(check_for_repeats_part_two('123122-123125'), [123123])
    
    def test_chunking(self):
        self.assertEqual(chunking('11',2), [1,1])
        self.assertEqual(chunking('111',3), [1,1,1])
        self.assertEqual(chunking('1111',2), [11,11])
        self.assertEqual(chunking('1111',4), [1,1,1,1])
        self.assertEqual(chunking('1234',2), [12,34])
        self.assertEqual(chunking('1234',4), [1,2,3,4])
        self.assertEqual(chunking('12345',5), [1,2,3,4,5])
        # self.assertEqual(chunking('1111',3), [1,1,1])
        self.assertEqual(chunking('192427',3), [19,24,27])
        

    def test_triple_repeating(self):
        self.assertEqual(check_for_repeats_part_two('110-112'), [111])
        self.assertEqual(check_for_repeats_part_two('121211-121213'), [121212])
    
    def test_sixtuple_repeating(self):
        self.assertEqual(check_for_repeats_part_two('222221-222223'), [222222])
        self.assertEqual(check_for_repeats_part_two('212120-212122'), [212121])
    
    def test_five_equal_repeats(self):
        self.assertEqual(check_for_repeats_part_two('2121212118-2121212124'), [2121212121])

    def test_based_on_example_part_two(self):
        self.assertEqual(check_for_repeats_part_two('11-22'), [11,22])
        self.assertEqual(check_for_repeats_part_two('95-115'), [99, 111])
        self.assertEqual(check_for_repeats_part_two('998-1012'), [999, 1010])
        self.assertEqual(check_for_repeats_part_two('1188511880-1188511890'), [1188511885])
        self.assertEqual(check_for_repeats_part_two('222220-222224'), [222222])
        self.assertEqual(check_for_repeats_part_two('1698522-1698528'), [])
        self.assertEqual(check_for_repeats_part_two('446443-446449'), [446446])
        self.assertEqual(check_for_repeats_part_two('38593856-38593862'), [38593859])
        self.assertEqual(check_for_repeats_part_two('565653-565659'), [565656])
        self.assertEqual(check_for_repeats_part_two('824824821-824824827'), [824824824])
        self.assertEqual(check_for_repeats_part_two('2121212118-2121212124'), [2121212121])
    
    def test_data_answer(self):
        test_data = readFile('D:/Documents/GitHub/advent-of-code/2025/Day2/testdata.txt')
        self.assertEqual(solution_one(test_data), 1227775554)
        self.assertEqual(solution_two(test_data), 4174379265)
        
if __name__ == '__main__':
    unittest.main()