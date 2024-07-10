import unittest
from io import StringIO
from unittest.mock import patch
from sort import sort_file, main

class TestSort(unittest.TestCase):
    def test_sort_file(self):
        input_content = """ZEBRA
APPLE
BANANA
CHERRY
APPLE
"""
        expected_output = """APPLE
APPLE
BANANA
CHERRY
ZEBRA"""

        with patch('builtins.open', return_value=StringIO(input_content)):
            result = sort_file('dummy_file.txt')
        self.assertEqual(result.strip(), expected_output.strip())


    def test_main_function(self):
        input_content = """ZEBRA
APPLE
BANANA
CHERRY
APPLE
"""

        expected_output = """APPLE
APPLE
BANANA
CHERRY
ZEBRA"""

        with patch('sys.argv', ['sort.py', 'dummy_file.txt']), \
            patch('builtins.open', return_value=StringIO(input_content)), \
            patch('sys.stdout', new=StringIO()) as fake_out:
            main()
            self.assertEqual(fake_out.getvalue().strip(), expected_output.strip())

    def test_sort_file_unique(self):
        input_content = """ZEBRA
APPLE
BANANA
CHERRY
APPLE
"""
        expected_output = """APPLE
BANANA
CHERRY
ZEBRA"""

        with patch('builtins.open', return_value=StringIO(input_content)):
            result = sort_file('dummy_file.txt', unique=True)
        self.assertEqual(result.strip(), expected_output.strip())

    def test_main_function_unique(self):
        input_content = """ZEBRA
APPLE
BANANA
CHERRY
APPLE
"""

        expected_output = """APPLE
BANANA
CHERRY
ZEBRA
"""

        with patch('sys.argv', ['sort.py', '-u', 'dummy_file.txt']), \
                patch('builtins.open', return_value=StringIO(input_content)), \
                patch('sys.stdout', new=StringIO()) as fake_out:
                main()
                self.assertEqual(fake_out.getvalue().strip(), expected_output.strip())



if __name__ == '__main__':
    unittest.main()

