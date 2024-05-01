import unittest
import io
import sys
import os 
from unittest import mock
from wc import count_bytes, count_lines, count_words, count_chars, default_count

class TestCountBytes(unittest.TestCase):
    def test_count_bytes(self):
        # Create temp file with sample content
        with open("test.txt", "w") as file:
            file.write("Sample content")

        # Call the count_bytes function
        count = count_bytes("test.txt")

        # Assert the expected byte count
        self.assertEqual(count, 14)

        # Clean up the temp file
        os.remove("test.txt")

    def test_default_count(self):
        with open("test.txt", "w") as file:
            file.write("Line 1\nLine 2\nLine 3\n")

        line_count, word_count, char_count = default_count("test.txt")

        self.assertEqual(line_count, 3)
        self.assertEqual(word_count, 6)
        self.assertEqual(char_count, 21)

        os.remove("test.txt")


    def test_file_not_found(self):
        # Call the count_bytes function with a non-existent file
        count = count_bytes("nonexistent.txt")

        # Assert that the count is None (indicating an error)
        self.assertIsNone(count)

    def test_count_lines(self):
        with open("test.txt", "w") as file:
            file.write("Line 1\nLine 2\nLine 3\n")

        count = count_lines("test.txt")

        self.assertEqual(count, 3)

        os.remove("test.txt")

    def test_count_words(self):
        with open("test.txt", "w") as file:
            file.write("This is a sample file\nwith multiple words\non each line\n")

        count = count_words("test.txt")

        self.assertEqual(count, 11)

        os.remove("test.txt")

    def test_count_chars(self):
        with open("test.txt", "w", encoding="utf-8") as file:
            file.write("Sample content with 🚀 emoji")

        char_count = count_chars("test.txt")
        byte_count = count_bytes("test.txt")

        self.assertEqual(char_count, 27)
        self.assertNotEqual(char_count, byte_count)

        os.remove("test.txt")

    def test_count_lines_stdin(self):
        input_data = "Line 1\nLine 2\nLine 3\n"
        with mock.patch('sys.stdin', io.StringIO(input_data)):
            count = count_lines(sys.stdin)
            self.assertEqual(count, 3)

    def test_count_words_stdin(self):
        input_data = "This is a sample input\nwith multiple words\non each line\n"
        with mock.patch('sys.stdin', io.StringIO(input_data)):
            count = count_words(sys.stdin)
            self.assertEqual(count, 11)

    def test_count_chars_stdin(self):
        input_data = "Sample content with 🚀 emoji"
        with mock.patch('sys.stdin', io.StringIO(input_data)):
            char_count = count_chars(sys.stdin)
            self.assertEqual(char_count, 27)

    def test_count_bytes_stdin(self):
        input_data = "Sample content"
        with mock.patch('sys.stdin', io.StringIO(input_data)):
            count = count_bytes(sys.stdin)
            self.assertEqual(count, 14)


if __name__ == "__main__":
    unittest.main()
