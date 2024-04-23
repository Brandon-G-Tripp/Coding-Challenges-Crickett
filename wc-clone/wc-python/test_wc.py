import unittest
import os 
from wc import count_bytes

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

    def test_file_not_found(self):
        # Call the count_bytes function with a non-existent file
        count = count_bytes("nonexistent.txt")

        # Assert that the count is None (indicating an error)
        self.assertIsNone(count)

if __name__ == "__main__":
    unittest.main()
