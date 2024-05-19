use std::{cmp::Ordering, collections::{BinaryHeap, HashMap}};

#[derive(Debug)]
pub struct Node {
    pub char: Option<char>,
    pub freq: usize, 
    pub left: Option<Box<Node>>,
    pub right: Option<Box<Node>>,
}

impl Node {
    pub fn new_leaf(char: char, freq: usize) -> Self {
        Node {
            char: Some(char),
            freq, 
            left: None,
            right: None,
        }
    }

    pub fn new_internal(left: Box<Node>, right: Box<Node>) -> Self {
        let freq = left.freq + right.freq;
        Node {
            char: None,
            freq,
            left: Some(left),
            right: Some(right),
        }
    }

    pub fn build_huffman_tree(freq_map: &HashMap<char, usize>) ->  Self {
        let mut heap = BinaryHeap::new();

        for (&char, &freq) in freq_map {
            let node = Node::new_leaf(char, freq);
            heap.push(Box::new(node));
        }

        while heap.len() > 1 {
            let right = heap.pop().unwrap();
            let left = heap.pop().unwrap();
            let internal_node = Node::new_internal(left, right);
            heap.push(Box::new(internal_node));
        }

        *heap.pop().unwrap()
    }
}

impl Ord for Node {
    fn cmp(&self, other: &Self) -> Ordering {
        other.freq.cmp(&self.freq)
    }
}

impl PartialOrd for Node {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl PartialEq for Node {
    fn eq(&self, other: &Self) -> bool {
        self.freq == other.freq
    }
}

impl Eq for Node {}

pub fn traverse_tree(node: &Node, depth: usize) {
    if let Some(c) = node.char {
        println!("{:indent$}{}: {}", "", c, node.freq, indent = depth * 2);
    } else {
        println!("{:indent$}Internal Node: {}", "", node.freq, indent = depth * 2);
        if let Some(left) = &node.left {
            traverse_tree(left, depth + 1);
        }
        if let Some(right) = &node.right {
            traverse_tree(right, depth + 1);
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_create_leaf_node() {
        let node = Node::new_leaf('A', 5 );
        assert_eq!(node.char, Some('A'));
        assert_eq!(node.freq, 5);
        assert!(node.left.is_none());
        assert!(node.right.is_none());
    }

    #[test]
    fn test_create_internal_node() {
        let left_child = Box::new(Node::new_leaf('A', 5));
        let right_child = Box::new(Node::new_leaf('B', 3));
        let node = Node::new_internal(left_child, right_child);
        assert_eq!(node.char, None);
        assert_eq!(node.freq, 8);
        assert!(node.left.is_some());
        assert!(node.right.is_some());
    }

    // #[test] 
    // fn test_build_huffman_tree() {
    //     let freq_map = HashMap::from([('A', 5), ('B', 3), ('C', 2), ('D', 1)]);
    //     let root = Node::build_huffman_tree(&freq_map);
    //     assert_eq!(root.freq, 11);
    //     assert_eq!(root.left.as_ref().unwrap().freq, 5);
    //     assert_eq!(root.right.as_ref().unwrap().freq, 6);
    //     assert_eq!(root.right.as_ref().unwrap().left.as_ref().unwrap().char, Some('B'));
    //     assert_eq!(root.right.as_ref().unwrap().left.as_ref().unwrap().char, Some('C'));
    //     assert_eq!(root.right.as_ref().unwrap().right.as_ref().unwrap().right.as_ref().unwrap().char, Some('D'));
    // }
    #[test]
    fn test_build_huffman_tree() {
        let freq_map = HashMap::from([('A', 5), ('B', 3), ('C', 2), ('D', 1)]);
        let root = Node::build_huffman_tree(&freq_map);
        assert_eq!(root.freq, 11);

        let mut char_freqs = HashMap::new();
        traverse_tree(&root, &mut char_freqs);

        assert_eq!(char_freqs[&'A'], 5);
        assert_eq!(char_freqs[&'B'], 3);
        assert_eq!(char_freqs[&'C'], 2);
        assert_eq!(char_freqs[&'D'], 1);
    }

    fn traverse_tree(node: &Node, char_freqs: &mut HashMap<char, usize>) {
        if let Some(c) = node.char {
            char_freqs.insert(c, node.freq);
        } else {
            if let Some(left) = &node.left {
                traverse_tree(left, char_freqs);
            }
            if let Some(right) = &node.right {
                traverse_tree(right, char_freqs);
            }
        }
    }
}
