package main

import (
	"fmt"
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
    frequencies := map[rune]int{
        'A': 5,
        'B': 2,
        'C': 1,
        'D': 1,
    }

    expectedTree := InternalNode{
        Left: InternalNode{
            Right: InternalNode{
                Left: LeafNode{Char: 'C', Freq: 1},
                Right: LeafNode{Char: 'D', Freq: 1},
                Freq: 2,
            },
            Left: LeafNode{Char: 'B', Freq: 2},
            Freq: 4,
        },
        Right: LeafNode{Char: 'A', Freq: 5},
        Freq: 9,
    }

    actualTree := BuildHuffmanTree(frequencies)
    if !treeEqual(actualTree, expectedTree) {
        t.Errorf("Actual tree does not match expected tree")
    }
}

func treeEqual(t1, t2 HuffmanNode) bool {
    if t1.IsLeaf() != t2.IsLeaf() {
        fmt.Printf("Mismatch: IsLeaf - t1: %v, t2: %v\n", t1.IsLeaf(), t2.IsLeaf())
        return false
    }

    if t1.Weight() != t2.Weight() {
        fmt.Printf("Mismatch: Weight - t1: %d, t2: %d\n", t1.Weight(), t2.Weight())
        return false
    }

    if t1.IsLeaf() {
        l1 := t1.(LeafNode)
        l2 := t2.(LeafNode)
        if l1.Char != l2.Char {
            fmt.Printf("Mismatch: Char - l1: %c, l2: %c\n", l1.Char, l2.Char)
            return false
        }
    } else {
        i1 := t1.(InternalNode)
        i2 := t2.(InternalNode)
        fmt.Printf("Comparing left subtree: t1=%+v, t2=%+v\n", i1.Left, i2.Left)
        if !treeEqual(i1.Left, i2.Left) {
            fmt.Println("Mismatch: Left subtree")
            return false
        }
        fmt.Printf("Comparing right subtree: t1=%+v, t2=%+v\n", i1.Right, i2.Right)
        if !treeEqual(i1.Right, i2.Right) {
            fmt.Println("Mismatch: Right subtree")
            return false
        }
    }

    return true
}

