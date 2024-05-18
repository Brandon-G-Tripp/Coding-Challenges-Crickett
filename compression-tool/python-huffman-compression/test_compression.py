import unittest
from unittest.mock import patch
from io import StringIO
from compression import count_frequencies

class TestCompression(unittest.TestCase):
    def test_invalid_file(self):
        with self.assertRaises(FileNotFoundError):
            with open('invalid_file.txt', 'r') as file:
                count_frequencies(file)

    def test_frequency_count(self):
        test_data = 'XXXtttttt'
        expected_frequencies = {'X': 3, 't': 6}
        self.assertEqual(count_frequencies(StringIO(test_data)), expected_frequencies)
