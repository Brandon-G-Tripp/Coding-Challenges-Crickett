import { readFileSync } from "node:fs";
import { countFrequencies } from "./compression";
import process from "node:process";
import { buildHuffmanTree, HuffmanNode, InternalNode, LeafNode } from "./huffmanTree";

const filePath = process.argv[2];

try {
    const text = readFileSync(filePath, 'utf-8');
    const frequencies = countFrequencies(text);
    console.log(`Character frequencies: ${frequencies}`);

    const huffmanTree = buildHuffmanTree(frequencies);
    console.log("Huffman Tree: ");
    printHuffmanTree(huffmanTree);
} catch (error) {
    console.error('Error reading file:', (error as Error).message);
    process.exit(1);
}

function printHuffmanTree(node: HuffmanNode, indent: string = ""): void {
    if (node.type === "leaf") {
        console.log(`${indent}Leaf: ${(node as LeafNode).char} (${node.freq})`);
    } else {
        const internalNode = node as InternalNode;
        console.log(`${indent}Internal Node (${internalNode.freq})`);
        printHuffmanTree(internalNode.left, indent + " ");
        printHuffmanTree(internalNode.right, indent + " ");
    }
}
