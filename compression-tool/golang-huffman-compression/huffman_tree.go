package main

import (
	"sort"
)

type HuffmanNode interface {
    IsLeaf() bool
    Weight() int
}

type LeafNode struct {
    Char rune
    Freq int
}

func (n LeafNode) IsLeaf() bool {
    return true
}

func (n LeafNode) Weight() int {
    return n.Freq
}

type InternalNode struct {
    Left HuffmanNode
    Right HuffmanNode
    Freq int
}

func (n InternalNode) IsLeaf() bool {
    return false
}

func (n InternalNode) Weight() int {
    return n.Freq
}


func BuildHuffmanTree(frequencies map[rune]int) HuffmanNode {
    var nodes []HuffmanNode
    for char, freq := range frequencies {
        nodes = append(nodes, LeafNode{Char: char, Freq: freq})
    }

    sort.Slice(nodes, func(i, j int) bool {
        if nodes[i].Weight() == nodes[j].Weight() {
            if nodes[i].IsLeaf() && nodes[j].IsLeaf() {
                return nodes[i].(LeafNode).Char < nodes[j].(LeafNode).Char
            }
            return nodes[i].IsLeaf()
        }
        return nodes[i].Weight() < nodes[j].Weight()
    })

    for len(nodes) > 1 {
        left, right := nodes[0], nodes[1]
        nodes = nodes[2:]

        internalNode := InternalNode{
            Left: left, 
            Right: right, 
            Freq: left.Weight() + right.Weight(),
        }
        nodes = append(nodes, internalNode)

        sort.Slice(nodes, func(i, j int) bool {
            if nodes[i].Weight() == nodes[j].Weight() {
                if nodes[i].IsLeaf() && nodes[j].IsLeaf() {
                    return nodes[i].(LeafNode).Char < nodes[j].(LeafNode).Char
                }
                return nodes[i].IsLeaf()
            }
            return nodes[i].Weight() < nodes[j].Weight()
        })
    }

    return nodes[0]
}

type HuffmanNodeHeap []HuffmanNode

func (h HuffmanNodeHeap) Len() int {
    return len(h)
}

func (h HuffmanNodeHeap) Less(i, j int) bool {
    if h[i].Weight() == h[j].Weight() {
        if h[i].IsLeaf() && h[j].IsLeaf() {
            return h[i].(LeafNode).Char < h[j].(LeafNode).Char
        }
        return h[i].IsLeaf()
    }
    return h[i].Weight() < h[j].Weight()
}

func (h HuffmanNodeHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *HuffmanNodeHeap) Push(x interface{}) {
    *h = append(*h, x.(HuffmanNode))
}

func (h *HuffmanNodeHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

