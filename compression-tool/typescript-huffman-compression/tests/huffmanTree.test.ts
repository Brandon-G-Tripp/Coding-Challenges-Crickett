import { buildHuffmanTree, HuffmanNode, InternalNode, LeafNode } from '../src/huffmanTree';

describe('buildHuffmanTree', () => {

    it('should build the correct Huffman tree for the given frequencies', () => {
        const frequencies: { [char: string]: number } = {
            a: 4,
            b: 2,
            c: 1,
            d: 1,
        };

        const expectedTree: InternalNode = {
            type: 'internal',
            freq: 8,
            left: {
                type: 'leaf',
                char: 'a',
                freq: 4,
            },
            right: {
                type: 'internal',
                freq: 4,
                left: {
                    type: 'leaf',
                    char: 'b',
                    freq: 2,
                },
                right: {
                    type: 'internal',
                    freq: 2,
                    left: {
                        type: 'leaf',
                        char: 'c',
                        freq: 1,
                    },
                    right: {
                        type: 'leaf',
                        char: 'd',
                        freq: 1,
                    },
                },
            },
        };


        expect(buildHuffmanTree(frequencies)).toEqual(expectedTree);
    });

    it('should handle frequencies with a single character', () => {
        const frequencies: { [char: string]: number} = {
            a: 5,
        };

        const expectedTree: LeafNode = {
            type: 'leaf',
            char: 'a',
            freq: 5,
        };

        expect(buildHuffmanTree(frequencies)).toEqual(expectedTree);
    });
})
