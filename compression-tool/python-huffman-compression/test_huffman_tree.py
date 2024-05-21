import unittest
from huffman_tree import HuffmanNode, build_huffman_tree

def compare_huffman_trees(tree1, tree2):
    if tree1 is None and tree2 is None:
        return True
    if tree1 is None or tree2 is None:
        print(f"One of the trees is None: tree1={tree1}, tree2={tree2}")
        return False
    if tree1.char != tree2.char or tree1.freq != tree2.freq:
        print(f"Mismatch in node attributes: tree1.char={tree1.char}, tree2.char={tree2.char}, tree1.freq={tree1.freq}, tree2.freq={tree2.freq}")
        return False
    print(f"Comparing left subtrees: tree1.left={tree1.left}, tree2.left={tree2.left}")
    left_equal = compare_huffman_trees(tree1.left, tree2.left)
    print(f"Comparing right subtrees: tree1.right={tree1.right}, tree2.right={tree2.right}")
    right_equal = compare_huffman_trees(tree1.right, tree2.right)
    return left_equal and right_equal

class TestHuffmanTree(unittest.TestCase):
    def test_build_huffman_tree(self):
        frequencies = {'a': 5, 'b': 2, 'c': 1, 'd': 1}
        expected_tree = HuffmanNode(None, 9,
                                    HuffmanNode(None, 4,
                                                HuffmanNode('b', 2),
                                                HuffmanNode(None, 2,
                                                            HuffmanNode('c', 1),
                                                            HuffmanNode('d', 1))),
                                    HuffmanNode('a', 5))
        actual_tree = build_huffman_tree(frequencies)
        self.assertEqual(actual_tree.freq, expected_tree.freq)
        self.assertEqual(actual_tree.char, expected_tree.char)
        self.assertTrue(compare_huffman_trees(actual_tree.left, expected_tree.left))
        self.assertTrue(compare_huffman_trees(actual_tree.right, expected_tree.right))

    def test_build_huffman_tree_single_character(self):
        frequencies = {'a': 1}
        expected_tree = HuffmanNode('a', 1)
        actual_tree = build_huffman_tree(frequencies)
        self.assertEqual(actual_tree.freq, expected_tree.freq)
        self.assertEqual(actual_tree.char, expected_tree.char)
        self.assertIsNone(actual_tree.left)
        self.assertIsNone(actual_tree.right)

if __name__ == '__main__':
    unittest.main()
